package router

import (
	"net/http"
	"project/project1/api"
	"project/project1/views"
)

// 写路由

func Router() {
	// 1 处理页面交给views 2 处理数据（json）交给 api 3静态资源

	// 跳到html里面   http://127.0.0.1:8888/
	// 从views里面的处理器
	http.HandleFunc("/", views.HTML.Index)
	// 跳到目录   http://127.0.0.1:8888/c/1
	http.HandleFunc("/c/1", views.HTML.Category)
	// 登录 http://127.0.0.1:8888/login   登录成功响应返回 http://127.0.0.1:8888/api/v1/login
	http.HandleFunc("/login",views.HTML.Login)

	// 文章的路由	 http://127.0.0.1:8888/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)

	// 数据  也交给对应的方法
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
http.HandleFunc("/api/v1/login",api.API.Login)
	// 静态资源的配置
	// 引导到自己的资源上   映射到自己的public下的resource上
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
