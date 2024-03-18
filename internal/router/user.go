package router

import (
	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/controllers"
)

// InitRouterUser Api
func InitRouterUser(r *gin.RouterGroup) {
	user := controllers.UserController{}
	user.Router(r)
}
