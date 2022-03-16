package main

import (
	// "log"

	"log"
	"math/rand"
	"net/http"

	// "net/http"
	"time"

	// "github.com/gin-gonic/gin"
	datebase "project/project2/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// 数据库测试
	// 删除
	// datebase.DropUser("hhh")
	// 查看
	// datebase.AllUsers()
	// 创建
	// datebase.CreateUser("hhh", "13076868372", "123123")
	// fmt.Println(datebase.IsTelephoneExist("13076868375"))

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数

		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		// 验证数据
		if len(telephone) != 11 {
			// r如果不是标准的 那么返回json数据  gin.H 是一个字典树组 value为任意值  key为字符串
			c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "密码规格不对"})
			return

		}
		if len(name) == 0 {
			// 如果没有传入 那么就会随机生成一个 随机的字符串   ch长度为10
			name = RandomString(10)
		}

		// 后端打印
		log.Println(name, telephone, password)

		// 判断手机是否存在
		if datebase.IsTelephoneExist(telephone) {
			// 存在下面的慈不了了
			c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "用户已存在"})
			return
		}
		// 创建用户
		err := datebase.CreateUser(name, telephone, password)
		if err {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "格式不催请重试"})
			return
		}
		// 返回结果
		c.JSON(200, gin.H{"msg": "注册成功"})

	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")

	// panic(r.Run())

	// 数据库
}

// 这是一个随机生成 字符串的 函数
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())

	for i := range result {
		// 这里是然result中的每一个i 填入 letters数组中的税金及一个
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
