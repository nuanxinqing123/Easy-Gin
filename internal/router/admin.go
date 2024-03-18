package router

import (
	"github.com/gin-gonic/gin"

	"Easy-Gin/internal/controllers"
)

// InitRouterAdmin API
func InitRouterAdmin(r *gin.RouterGroup) {
	user := controllers.UserAdminController{}
	user.Router(r)
}
