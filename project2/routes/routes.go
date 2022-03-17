package routes

import (
	"project/project2/controller"
	"project/project2/middleware"

	"github.com/gin-gonic/gin"
)

// 这里专门来弄路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	// 注册
	r.POST("/api/auth/register", controller.Register)
	// 登录
	r.POST("/api/auth/login", controller.Login)
	// 用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
