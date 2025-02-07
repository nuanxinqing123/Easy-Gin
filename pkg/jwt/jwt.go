package jwt

import (
	"errors"
	"time"

	"Easy-Gin/config"

	"github.com/golang-jwt/jwt/v4"
)

// TokenType 定义token类型
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type CustomClaims struct {
	BaseClaims
	TokenType  TokenType // token类型：access 或 refresh
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	JTI    string // 全局唯一ID
	UserId string // 用户ID
	Role   string // 角色
}

type JWT struct {
	SigningKey []byte
}

// TokenInfo 返回给客户端的token信息
type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var (
	ErrTokenExpired     = errors.New("token is expired")           // token已过期
	ErrTokenNotValidYet = errors.New("token not active yet")       // token未激活
	ErrTokenMalformed   = errors.New("that's not even a token")    // token不是一个token
	ErrTokenInvalid     = errors.New("couldn't handle this token") // token无法处理
	ErrTokenTypeError   = errors.New("token type is invalid")      // token类型不正确
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.GinConfig.JWT.SigningKey),
	}
}

// createClaims 创建Claims
func (j *JWT) createClaims(baseClaims BaseClaims, tokenType TokenType) CustomClaims {
	var expiresTime time.Duration

	// 根据token类型设置过期时间
	if tokenType == AccessToken {
		expiresTime = time.Hour * 24 * time.Duration(config.GinConfig.JWT.ExpiresTime)
	} else {
		expiresTime = time.Hour * 24 * time.Duration(config.GinConfig.JWT.RefreshExpiresTime)
	}

	claims := CustomClaims{
		BaseClaims: baseClaims,
		TokenType:  tokenType,
		BufferTime: config.GinConfig.JWT.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Millisecond)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)),
			Issuer:    config.GinConfig.JWT.Issuer,
		},
	}
	return claims
}

// CreateTokens 创建Access Token和Refresh Token
func (j *JWT) CreateTokens(baseClaims BaseClaims) (*TokenInfo, error) {
	// 创建Access Token
	accessClaims := j.createClaims(baseClaims, AccessToken)
	accessToken, err := j.createToken(accessClaims)
	if err != nil {
		return nil, err
	}

	// 创建Refresh Token
	refreshClaims := j.createClaims(baseClaims, RefreshToken)
	refreshToken, err := j.createToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// createToken 创建一个token
func (j *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// RefreshTokens 使用Refresh Token刷新Access Token和Refresh Token
func (j *JWT) RefreshTokens(refreshToken string) (*TokenInfo, error) {
	// 解析refresh token
	claims, err := j.ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// 验证token类型
	if claims.TokenType != RefreshToken {
		return nil, ErrTokenTypeError
	}

	// 创建新的token对
	return j.CreateTokens(claims.BaseClaims)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	} else {
		return nil, ErrTokenInvalid
	}
}
