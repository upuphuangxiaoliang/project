package views

import (
	"errors"
	"log"
	"net/http"
	"project/project1/common"
	"project/project1/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {

	// 与之前的一样 需要把信息写进去

	// 先拿到页面
	detail := common.Template.Detail

	// 获取路径参数

	//拿到路径 冲r这里拿
	//http://localhost:8080/c/1  1参数 分类的id
	path := r.URL.Path
	// 去掉前面的/c/ 得到后面的1
	pIdStr := strings.TrimPrefix(path, "/p/")
	// 7.html  把后面的。html去掉
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	// 转成对应的int
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别这个请求路径"))
		return
	}

	postRes, err := service.GetPostDetal(pid)

	if err != nil {
		log.Println("获取数据出错：", err)
		detail.WriteError(w, errors.New("系统错误请联系管理员"))
	}

	detail.WriteData(w, postRes)

}
