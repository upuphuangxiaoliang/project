package controller

import (
	"net/http"
	"project/project2/dto"
	"project/project2/model"

	"github.com/gin-gonic/gin"
)

// 用户信息
func Info(c *gin.Context) {
	// 获取中间件验证过后传过来的user
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
