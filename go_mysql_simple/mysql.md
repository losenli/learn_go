# GO连接mysql数据库

#### 1、创建mysql连接+测试连接：

```
func (d *DBUtil) newDB(conf string) *DBUtil {
    // 1、open 创建 mysql连接
	db ,err := sql.Open("mysql", conf)
	if err != nil{
		panic(err)
	}
    // 2、测试mysql连接
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	log.Println("mysql连接成功")
	return &DBUtil{db}
}
// db配置格式
dbConf := "root:123456@tcp(localhost:3306)/learn_go"
```

#### 2、查询单条数据：

```
func (d *DBUtil) findById(id int) User {
	var user User
	_sql := "select id,name,age,addr from user where id = ?"
    // QueryRow 查询一行数据，Scan将数据赋值给对应结构体，并释放连接
	err := d.QueryRow(_sql, id).Scan(&user.id,&user.name,&user.age,&user.addr)
	if err != nil{
		log.Println(err)
	}
	return user
}
```

#### 3、查询多条数据：

```
func (d *DBUtil) query(queryStr ...string) []User {
	var user User
	var userList []User
	_sql := "select * from user where 1=1 and "
	_sql = _sql + strings.Join(queryStr, " and ")
    // Query:查询多条 --> rows
	rows, err := d.Query(_sql)
	if err != nil{
		panic(err)
	}
    // rows.close()释放连接
	defer rows.Close()
    // 遍历rows每行的数据
	for rows.Next(){
		err = rows.Scan(&user.id,&user.name,&user.age,&user.addr)
		if err != nil{
			log.Println(err)
		}
		userList = append(userList, user)
	}
	return userList
}
```

#### 4、关闭数据库连接:

```
func (d *DBUtil)closeDB()  {
	err := d.Close()
	if err != nil{
		panic(err)
	}
	log.Println("成功关闭数据库连接")
}
```

#### 5、插入+修改+删除数据：
```
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
```

#### 6、批量操作使用预处理：

```
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
fmt.Println(t.Seconds()) // 1.117559388s
```

#### 7、事务执行多条sql语句：

```
func (d *DBUtil)transactionDemo()  {
	// 开启事务
	tx, err := d.Begin()
	if err != nil {
		panic(err)
	}
	_sql1 := `insert into user (name, age, addr) value ("li1", 99, "jx")`
	_sql2 := `insert into user1 (name, age, addr) value ("li1", 100, "jx" )`
    // 执行语句
	_, err = tx.Exec(_sql1)
	if err != nil {
		fmt.Println("sql1") 
        // 异常回滚
		_ = tx.Rollback()
	}
	_, err = tx.Exec(_sql2)
	if err != nil {
		fmt.Println("sql2")
		_ = tx.Rollback()
	}
    // 提交操作
	if err = tx.Commit(); err != nil{
		_ = tx.Rollback()
	}
}
```