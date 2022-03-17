package controller

import (
	"log"
	"net/http"
	"project/project2/common"
	datebase "project/project2/database"
	"project/project2/model"
	"project/project2/response"

	"github.com/gin-gonic/gin"
)

// 利用手机号登录
func Login(c *gin.Context) {
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 验证数据
	if len(telephone) != 11 {
		// r如果不是标准的 那么返回json数据  gin.H 是一个字典树组 value为任意值  key为字符串
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")

		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码规格不对")

		return
	}

	// 后端打印
	log.Println(telephone, password)

	// 判断手机是否存在 存在才能登录
	if !datebase.IsTelephoneExist(telephone) {
		// 存在下面的慈不了了
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		return
	}
	// 判断密码是否存在
	user := datebase.IsPasswordRight(telephone, password)
	if (user == model.User{}) {
		// 如果返回false
		c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 500, "msg": "密码错误"})
		return
	}
	// 同uo则放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "系统和错误"})
		log.Printf("token error:", err)
		return
	}
	// 返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")

}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjgsImV4cCI6MTY0ODEwNjkyNSwiaWF0IjoxNjQ3NTAyMTI1LCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.8OVuT_2jmxDuTgqKxjzuGPlqvVWvQzpjGTn7NUE1KdQ
