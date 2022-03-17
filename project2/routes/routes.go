package routes

import (
	"project/project2/controller"

	"github.com/gin-gonic/gin"
)

// 这里专门来弄路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	// 注册
	r.POST("/api/auth/register", controller.Register)
	// 登录
	r.POST("/api/auth/login", controller.Login)
	return r
}
