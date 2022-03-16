package views

import (
	"net/http"
	"project/project1/common"
	"project/project1/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {

	// 与之前的一样 需要把信息写进去

	// 先拿到页面
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)

}
