package main

import (
	"project/project2/routes"

	"github.com/gin-gonic/gin"
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

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r = routes.CollectRoute(r)

	r.Run(":8000")
	panic(r.Run())

}
