package controller

// ///////////////////////////////////////////////////这个是用来
import (
	"log"
	"math/rand"
	"net/http"
	datebase "project/project2/database"
	"project/project2/model"
	"project/project2/response"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// 获取参数
	Name, Telephone, Password := c.PostForm("name"), c.PostForm("telephone"), c.PostForm("password")
	// 验证数据
	if len(Telephone) != 11 {
		// r如果不是标准的 那么返回json数据  gin.H 是一个字典树组 value为任意值  key为字符串
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"cood": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(Password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "面膜不能少于6位")

		return

	}
	if len(Name) == 0 {
		// 如果没有传入 那么就会随机生成一个 随机的字符串   ch长度为10
		Name = RandomString(10)
	}

	// 后端打印
	log.Println(Name, Telephone, Password)

	// 判断手机是否存在
	if datebase.IsTelephoneExist(Telephone) {
		// 存在下面的慈不了了
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")

		return
	}
	// 创建用户
	// 密码加密
	hasedPassword, er := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if er != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "加密错误")

		return
	}
	// fmt.Println(Password, string(hasedPassword))
	user := model.User{Name, Telephone, string(hasedPassword)}
	err := datebase.CreateUser(user)
	if err {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "格式不催请重试")

		return
	}
	// 返回结果
	response.Success(c, nil, "注册成功")
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
