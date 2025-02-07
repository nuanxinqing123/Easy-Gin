package controllers

import (
	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/model"
	"Easy-Gin/internal/service"
	res "Easy-Gin/pkg/response"
)

type OpenController struct{}

// Router 注册路由
func (c *OpenController) Router(r *gin.RouterGroup) {
	// 登录
	r.POST("login", c.Login)
	// 注册
	r.POST("register", c.Register)
	// 刷新Token
	r.POST("refresh", c.RefreshToken)
}

// Login 用户登录
func (c *OpenController) Login(ctx *gin.Context) {
	// 获取参数
	p := new(model.Login)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, err)
		return
	}

	// 业务处理
	resCode, msg := service.Login(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// Register 用户注册
func (c *OpenController) Register(ctx *gin.Context) {
	// 获取参数
	p := new(model.Register)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, err)
		return
	}

	// 业务处理
	resCode, msg := service.Register(p)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}

// RefreshToken 刷新Token
func (c *OpenController) RefreshToken(ctx *gin.Context) {
	// 获取参数
	type RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"` // refresh token
	}
	p := new(RefreshTokenRequest)
	if err := ctx.ShouldBindJSON(&p); err != nil {
		res.ResErrorWithMsg(ctx, res.CodeInvalidParam, err)
		return
	}

	// 业务处理
	resCode, msg := service.RefreshToken(p.RefreshToken)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}
