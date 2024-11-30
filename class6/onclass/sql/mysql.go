package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//下面是一个较完整配置的连接
	//其中后面的参数用于定制连接行为
	//charset：设置字符集，例如utf8mb4
	//parseTime：如果设为true，则MySQL中的DATE和DATETIME类型会被解析为Go的time.Time类型
	//loc：用于解析时间的位置信息，如Asia/Shanghai
	dsn := "root:123456@tcp(127.0.0.1:3306)/stus?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
	//可选参数有：
	//timeout：连接超时时间，例如30s。
	//readTimeout 和 writeTimeout：读写操作的超时时间。
	//collation：指定字符序。
	//maxAllowedPacket：允许的最大数据包大小。
	//tls：是否启用TLS/SSL加密，可选值包括skip-verify, preferred, required等
	//
	//第一部分：连接数据库
	db, err := sql.Open("mysql", dsn) //链接数据库
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("open database error: ", err)
	}
	//
	//第二部分：遍历查找
	rows, err := db.Query("select * from students where id= ?", 1)
	//在 Query 语句中 ? 表示占位符，在这里同样的效果就是
	// select * from stu where id= 1
	if err != nil {
		log.Println(err)
		return
	}
	// 延迟调用关闭rows释放持有的数据库链接
	defer rows.Close()
	var students struct {
		id     int
		name   string
		gender string
		age    int
	}
	// 迭代查询获取数据，必须调用
	for rows.Next() {
		err := rows.Scan(&students.id, &students.name, &students.age, &students.gender)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(students)
	}

	fmt.Println(students)
	row := db.QueryRow("select * from stu where id=?", 20)
	var user struct {
		name    string
		id      int
		math    int
		english int
	}
	// 需要注意在执行Scan的时候需要按照获取的元素个数获取
	// 同时需要加上&符号(取地址) 按照先后顺序查询
	err = row.Scan(&user.id, &user.name, &user.math, &user.english)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(user)

}
