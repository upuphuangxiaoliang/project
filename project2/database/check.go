package datebase

import (
	"fmt"
	"log"
	"project/project2/model"

	"golang.org/x/crypto/bcrypt"
)

// 查询是电话是否存在
func IsTelephoneExist(t string) bool {
	fmt.Println("查询电话号码是否存在")
	row := DB.QueryRow("SELECT id FROM users where telephone=? ", t)
	if row.Err() != nil {
		log.Println(row.Err())

	}
	var id int
	if err := row.Scan(&id); err != nil {
		// 这里才是判断找没找到
		fmt.Println("不存在该用户 可以创建")
		return false
	}
	fmt.Println("已经存在")

	return true

}

// 查询密码是否正确  正确的话返回的是user
func IsPasswordRight(t string, p string) model.User {
	fmt.Println("查询密码是否正确")
	row := DB.QueryRow("SELECT password FROM users where telephone=? ", t)
	if row.Err() != nil {
		log.Println(row.Err())

	}
	var password string
	if err := row.Scan(&password); err != nil {
		// 这里才是判断找没找到
		fmt.Println("不存在该用户")
		return model.User{}
	}
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(p))
	if err != nil {
		fmt.Println("密码不正确 返回空的user")
		// 返回空
		return model.User{}
	}
	fmt.Println("密码正确,开始返回user")

	// 开始返回user
	rrow := DB.QueryRow("SELECT * FROM users where telephone=? ", t)
	if rrow.Err() != nil {
		log.Println(rrow.Err())
	}
	fmt.Println("chengon")
	var id int
	var name string
	var telephone string
	var password2 string
	var time1 string
	var time2 string
	if err := rrow.Scan(&id, &name, &telephone, &password2, &time1, &time2); err != nil {
		// 这里显示查没查到
		log.Fatal(err)
	}
	fmt.Println("返回user:", model.User{name, telephone, password2})
	return model.User{name, telephone, password2}

}

func ID(u model.User) int {
	fmt.Println("查询id")
	row := DB.QueryRow("SELECT id FROM users where name=? ", u.Name)
	if row.Err() != nil {
		log.Println(row.Err())

	}
	var id int
	if err := row.Scan(&id); err != nil {
		// 这里才是判断找没找到
		fmt.Println("不存在该用户")
		return -1
	}
	fmt.Println("id为", id)
	return id

}
