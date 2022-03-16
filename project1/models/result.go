package models

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

// type UserInfo struct {
// 	Uid      int    `json:"uid"`
// 	UserName string `json:"userName"`
// 	Avatar   string `json:"avatar"`
// }
