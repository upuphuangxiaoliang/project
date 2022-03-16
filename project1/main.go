package main

import (
	"fmt"
	"log"
	"net/http"
	"project/project1/common"
	"project/project1/router"
)

// "encoding/json"

// // 定义index 数据的结构体              一个不要模板的实例
// type IndexData struct {
// 	Title string `json:"title"` // 转成想要的json 首字母不是大写的
// 	Desc  string
// }

// //定义 index 的内容
// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var indexdata IndexData
// 	indexdata.Title = "hello"
// 	indexdata.Desc = "入门教程"

// 	// 转换成json格式传输
// 	jsonStr, _ := json.Marshal(indexdata)
// 	w.Write(jsonStr)
// 	// 普通传输
// 	w.Write([]byte("asdasdasd"))
// }

func init() {
	// 程序刚开始的时候 进行模板加载
	common.LoadTemplate()
}
func main() {
	// web 入口
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}
	fmt.Println("ok")

	// 路由
	// router 文件下的 Router 方法
	router.Router()
	// 127.0.0.1:8888/home

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)

	}

	// dao.Test()

}
