package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/model"
	_jwt "Easy-Gin/pkg/jwt"
	res "Easy-Gin/pkg/response"
)

const CtxUserID = "UserID"

var jwt _jwt.JWT

// 定义白名单
var whiteList []string

// Auth 用户基于JWT的认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// 检查请求路径是否在白名单中
			path := c.Request.URL.Path
			for _, v := range whiteList {
				if path == v {
					c.Next()
					return
				}
			}
			// 非白名单请求且无Token
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResErrorWithMsg(c, res.CodeNeedLogin, "用户状态已失效")
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResErrorWithMsg(c, res.CodeNeedLogin, "用户状态已失效")
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(CtxUserID, mc.UserId)
		c.Next()
	}
}

// AdminAuth 管理员JWT的认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResErrorWithMsg(c, res.CodeNeedLogin, "用户状态已失效")
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResErrorWithMsg(c, res.CodeNeedLogin, "用户状态已失效")
			c.Abort()
			return
		}

		// 检查是否属于管理员
		if mc.Role != model.AdminRole {
			c.Abort()
			res.ResErrorWithMsg(c, res.CodeNeedLogin, "无访问权限或认证已过期")
			return
		} else {
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set(CtxUserID, mc.UserId)
			c.Next()
		}
	}
}
