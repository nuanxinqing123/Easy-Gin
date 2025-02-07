package controllers

import (
	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/service"
	res "Easy-Gin/pkg/response"
)

const CtxUserID = "UserID"

/*  用户操作  */

type UserController struct{}

// Router 注册路由
func (c *UserController) Router(r *gin.RouterGroup) {
	// 获取用户信息
	r.GET("info", c.Info)
}

// Info 获取用户信息
func (c *UserController) Info(ctx *gin.Context) {
	// 获取 UserID
	userId := ctx.GetString(CtxUserID)

	// 业务处理
	resCode, msg := service.Info(userId)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // 成功
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // 失败
	}
}
