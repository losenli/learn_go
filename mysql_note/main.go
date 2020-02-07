package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

/**
* @Author: Mr-Li
* @Date: 2020/2/6
 */

type DBUtil struct {
	*sql.DB
}

type User struct {
	id int
	name string
	age int
	addr string
}

// 初始化数据库连接池对象
func (d *DBUtil) newDB(conf string) *DBUtil {
	db ,err := sql.Open("mysql", conf)
	if err != nil{
		panic(err)
	}
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	log.Println("mysql连接成功")
	return &DBUtil{db}
}

// 根据id查询一条数据
func (d *DBUtil) findById(id int) User {
	var user User
	_sql := "select id,name,age,addr from user where id = ?"
	err := d.QueryRow(_sql, id).Scan(&user.id,&user.name,&user.age,&user.addr)
	if err != nil{
		log.Println(err)
	}
	return user
}

// 关闭db连接
func (d *DBUtil)closeDB()  {
	err := d.Close()
	if err != nil{
		panic(err)
	}
	log.Println("成功关闭数据库连接")
}

// 传入条件查询多条数据
func (d *DBUtil) query(queryStr ...string) []User {
	var user User
	var userList []User
	_sql := "select * from user where 1=1 and "
	_sql = _sql + strings.Join(queryStr, " and ")
	rows, err := d.Query(_sql)
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&user.id,&user.name,&user.age,&user.addr)
		if err != nil{
			log.Println(err)
		}
		userList = append(userList, user)
	}
	return userList
}

func (d *DBUtil)prepareQueryDemo()  {
	stmt, err := d.Prepare(`insert into user (name, age, addr) value (?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	fmt.Println(stmt)
}

// 根据id删除一条数据
func (d *DBUtil)deleteById(id int) int {
	_sql := "delete from user where id = ?"
	ret, err := d.Exec(_sql, id)
	if err != nil{
		panic(err)
	}
	n, _ := ret.RowsAffected()
	return int(n)
}

// 插入一条数据
func (d *DBUtil)insertOne(user *User) (int, int) {
	_sql := `insert into user (id, name, age, addr) value (?, ?, ?, ?)`
	ret, err := d.Exec(_sql, user.id, user.name, user.age, user.addr)
	if err != nil{
		panic(err)
	}
	n, _ := ret.RowsAffected()
	lastId, _ := ret.LastInsertId()
	return int(n), int(lastId)
}

func (d *DBUtil)transactionDemo()  {
	// 开启事务
	tx, err := d.Begin()
	if err != nil {
		panic(err)
	}
	_sql1 := `insert into user (name, age, addr) value ("li1", 99, "jx")`
	_sql2 := `insert into user1 (name, age, addr) value ("li1", 100, "jx" )`

	_, err = tx.Exec(_sql1)
	if err != nil {
		fmt.Println("sql1")
		_ = tx.Rollback()
	}
	_, err = tx.Exec(_sql2)
	if err != nil {
		fmt.Println("sql2")
		_ = tx.Rollback()
	}
	if err = tx.Commit(); err != nil{
		_ = tx.Rollback()
	}

}

func main() {
	var d DBUtil
	dbConf := "root:123456@tcp(localhost:3306)/learn_go"
	db := d.newDB(dbConf)

	//log.Println(db.findById(2))
	//defer db.closeDB()
	//userList := db.query("id>0", "age=23")
	//for _, user := range userList{
	//	log.Println(user)
	//}
	//log.Println("删除行数：", db.deleteById(1))

	//u := User{name:"losen", id:7, age:99, addr:"sh"}
	//n, lastId := db.insertOne(&u)
	//log.Println("插入行数：", n, lastId)

	start := time.Now()
	// 预处理时间比较
	_sql := `insert into user (name, age, addr) value (?, ?, ?)`
	// 批量操作使用prepare预处理节约时间
	stmt, err := db.Prepare(_sql)
	if err != nil {
		panic(err)
	}
	for i:=0;i<1000;i++{
		_, _ = stmt.Exec("lzh", 23, "jx")
	}
	end := time.Now()

	t := end.Sub(start)
	fmt.Println(t.Seconds()) // 1.117559388

	//start := time.Now()
	//// 预处理时间比较
	//_sql := `insert into user (name, age, addr) value (?, ?, ?)`
	////stmt, err := db.Prepare(_sql)
	////if err != nil {
	////	panic(err)
	////}
	//for i:=0;i<1000;i++{
	//	_, _ = db.Exec(_sql,"lzh", 23, "jx")
	//}
	//end := time.Now()
	//
	//t := end.Sub(start)
	//fmt.Println(t.Seconds())

	db.transactionDemo()

}
