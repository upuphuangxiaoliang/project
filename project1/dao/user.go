package dao

import (
	"log"
	"project/project1/models"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	// 查询到之后取出来
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd = ?",
		username,
		passwd)
	if row.Err() != nil {
		// 如果没有找到返回nil
		log.Println(row.Err())
		return nil
	}
	// 查询到之后取出来
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user

}
