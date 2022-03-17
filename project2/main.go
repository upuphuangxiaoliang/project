package main

import (
	"os"
	"project/project2/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// 数据库测试
	// 删除asda
	// datebase.DropUser("hhh")
	// 查看
	// datebase.AllUsers()
	// 创建
	// datebase.CreateUser("hhh", "13076868372", "123123")
	// fmt.Println(datebase.IsTelephoneExist("13076868375"))
	// 查看id
	// user := model.User{"ngjhufoogt", "13123", "asdasd"}
	// fmt.Println(datebase.ID(user))
	// 同过id查看user
	// datebase.ID2User(4)

	// 读取配置
	InitConfig()
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r = routes.CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
