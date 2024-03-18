package controllers

import (
	"time"

	"github.com/gin-gonic/gin"

	"Easy-Gin/config"
	"Easy-Gin/internal/service"
	res "Easy-Gin/pkg/response"
)

const CtxUserID = "UserID"
const jti = "jti"

/*  用户操作  */

type UserController struct{}

// Router 注册路由
func (c *UserController) Router(r *gin.RouterGroup) {
	// 登出
	r.POST("logout", c.Logout)

	// 获取用户信息
	r.GET("info", c.Info)
}

// Logout 登出
func (c *UserController) Logout(ctx *gin.Context) {
	// 写入Redis, 设置为黑名单内容
	config.GinRedis.Set(ctx.GetString(jti), true, time.Duration(config.GinConfig.JWT.ExpiresTime))
	res.ResSuccess(ctx, "退出成功")
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

/*  管理员操作  */

type UserAdminController struct{}

// Router 注册路由
func (c *UserAdminController) Router(r *gin.RouterGroup) {

}
