package service

import (
	"errors"
	"project/project1/dao"
	"project/project1/models"
	"project/project1/utils"
)

// 这里处理的时登录的时候的信息 返回的是models.LohinRes
// func Login(userName, passwd string) (*models.LoginRes, error) {
// 	// 进行加密  让密码更加安全
// 	passwd = utils.Md5Crypt(passwd, "mszlu")
// 	user := dao.GetUser(userName, passwd)
// 	if user == nil {
// 		// meiy没有找到
// 		return nil, errors.New("账号密码不正确")
// 	}
// 	uid := user.Uid
// 	// 生成token  jwt的技术 进行生成 相当于一个令牌
// 	token, err := utils.Award(&uid)
// 	if err != nil {
// 		return nil, errors.New("token未能生成")
// 	}
// 	var userInfo models.UserInfo
// 	userInfo.Uid = user.Uid
// 	userInfo.UserName = user.UserName
// 	userInfo.Avatar = user.Avatar
// 	var lr = &models.LoginRes{
// 		token,
// 		userInfo,
// 	}
// 	return lr, nil
// }

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token  jwt技术进行生成 令牌  A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
