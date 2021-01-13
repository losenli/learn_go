package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
* @Author: Mr-Li
* @Date: 2020/2/7
 */

var db *sqlx.DB

type User struct {
	ID   int
	Name string
	Age  int
	Addr string
}

func initDB() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/learn_go"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var user User
	err := db.Get(&user, sqlStr, 12)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", user.ID, user.Name, user.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age, addr from user"
	var users []User
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func main() {
	_ = initDB()
	queryRowDemo()
}
