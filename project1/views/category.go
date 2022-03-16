package views

import (
	"errors"
	"net/http"
	"project/project1/common"
	"project/project1/service"
	"strconv"
	"strings"
)

// 查看文章列表
func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//拿到路径 冲r这里拿
	//http://localhost:8080/c/1  1参数 分类的id
	path := r.URL.Path
	// 去掉前面的/c/ 得到后面的1
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别这个请求路径"))
		return
	}

	// 从前端获取 页数
	pageStr := r.Form.Get("page")

	if pageStr == "" {
		//转化成int类型
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	// 每页显示多少
	pagesize := 10

	categoryResponse, err := service.GetPostByCategoryId(cId, page, pagesize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}

	// 把数据写上去
	categoryTemplate.WriteData(w, categoryResponse)

}
