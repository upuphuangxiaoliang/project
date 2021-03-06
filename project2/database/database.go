package datebase

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"project/project2/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Test() {
	// 插入
	// DB.Exec("insert into todos(title,status) values('123123',0)")

	// 另外一种查询 没有找到二时候会返回 错误
	row := DB.QueryRow("SELECT * FROM users where telephone=? ", "13")
	if row.Err() != nil {
		log.Println(row.Err())
	}
	fmt.Println("chengon")
	var id int
	var name string
	var telephone string
	var password string
	if err := row.Scan(&id, &name, &telephone, &password); err != nil {
		// 这里显示查没查到
		log.Fatal(err)
	}
	fmt.Println(id, name, telephone, password)

	fmt.Println("yes")

	// 查询  这种查询 是会显示多个行
	// rows, err := DB.Query("SELECT * FROM users where name=? ", "hzl")
	// rows, err := DB.Query("SELECT * FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if rows.Next() == false {
	// 	fmt.Println("无")
	// }

	// for rows.Next() {
	// 	// 一定要对应各个字段
	// 	var id int
	// 	var name string
	// 	var telephone string
	// 	var password string
	// 	if err := rows.Scan(&id, &name, &telephone, &password); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id, name, telephone, password)
	// }

	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)

	// }

}

func init() {
	//执行main之前 先执行init方法
	// 初始化数据库    这里是选择了数据库 相当于use 。。。 才有后面的操作  其中12345是密码
	dataSourceName := fmt.Sprintf("root:12345@tcp(localhost:3306)/project2?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
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

func CreateUser(user model.User) bool {
	fmt.Println("注册用户")
	_, err := DB.Exec("insert into users(name,telephone,password) values(?,?,?)", user.Name, user.Telephone, user.Password)
	if err != nil {
		fmt.Println("创建用户失败")
		return true
	}
	fmt.Println("创建用户成功  所有用户如下")
	AllUsers()
	return false
}
func AllUsers() {
	// 查询  这种查询 是会显示多个行
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		// 一定要对应各个字段
		var id int
		var name string
		var telephone string
		var password string
		var create_time string
		var uppdata_time string
		if err := rows.Scan(&id, &name, &telephone, &password, &create_time, &uppdata_time); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, telephone, password)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)

	}
}
func DropUser(id int) {
	fmt.Println("删除用户")
	_, err := DB.Exec("delete from users where id=?", id)
	if err != nil {
		fmt.Println("删除用户失败")
		return
	}
	fmt.Println("删除用户成功 所有用户如下")
	AllUsers()

}

