package initialize

import (
	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/middleware"
	"Easy-Gin/internal/router"
	res "Easy-Gin/pkg/response"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.Logger())
	Router.Use(middleware.Recovery())

	// 允许跨域
	Router.Use(middleware.Cors())

	// (可选项)
	// PID 限流基于实例的 CPU 使用率，通过拒绝一定比例的流量, 将实例的 CPU 使用率稳定在设定的阈值上。
	// 地址: https://github.com/bytedance/pid_limits
	// Router.Use(adaptive.PlatoMiddlewareGinDefault(0.8))

	PingGroup := Router.Group("")
	{
		// 存活检测
		PingGroup.GET("/ping", func(c *gin.Context) {
			res.ResSuccess(c, "pong")
		})
	}

	ApiGroupOpen := Router.Group("/")
	router.InitRouterOpen(ApiGroupOpen)

	ApiGroupAuth := Router.Group("/api")
	ApiGroupAuth.Use(middleware.Auth())
	router.InitRouterUser(ApiGroupAuth)

	ApiGroupAdmin := Router.Group("/api/admin")
	ApiGroupAdmin.Use(middleware.AdminAuth())
	router.InitRouterAdmin(ApiGroupAdmin)

	return Router
}
