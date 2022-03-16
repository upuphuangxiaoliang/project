package views

import (
	"errors"
	"log"
	"net/http"
	"project/project1/common"
	"project/project1/service"
	"strconv"
)

// 这里是首页里面的全部内容  
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// zhh这里是拿到模板
	index := common.Template.Index
	// cc传入数据给前端     需要给全前端需要的数据都要有定义
	// 再models这个文件夹下找到相对应的数结构体  因为这个结构体才符合前端的参数需求
	// 假数据
	//页面上涉及到的所有的数据，必须有定义

	// 数据库查询 交给service 层次 数据库中的分类类别  获取所有的信息
	// 页数的处理
	if err := r.ParseForm(); err != nil {
		log.Println("获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误请联系管理员"))
		return
	}
	// 从前端获取 页数
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		//转化成int类型
		page, _ = strconv.Atoi(pageStr)
	}

	// 把页数 与每页显示的数量全进去

	pagesize := 5 //表示5页一张纸
	hr, err := service.GetAllIndexInfo(page, pagesize)
	if err != nil {
		log.Println("获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误请联系管理员"))
	}

	// 这个操作是填充数据
	index.WriteData(w, hr)

	//

}
