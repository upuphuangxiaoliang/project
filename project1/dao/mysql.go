package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"project/project1/models"
	"time"

	// 不一定使用它  所以加上小横杆
	_ "github.com/go-sql-driver/mysql"
)

func Test() {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd = ?",
		"admin",
		"1254c27e22c2ec22e62855fd5878027")
	fmt.Println("123123123")
	if row.Err() != nil {
		// 如果没有找到返回nil
		log.Println(row.Err())

	}
	fmt.Println("chengon")
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("yes")
}

// golang 数据库先关的东西
var DB *sql.DB

func init() {
	//执行main之前 先执行init方法
	// 初始化数据库    这里是选择了数据库 相当于use 。。。 才有后面的操作  其中12345是密码
	dataSourceName := fmt.Sprintf("root:12345@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = db
}
