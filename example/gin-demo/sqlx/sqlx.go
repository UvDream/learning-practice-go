package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type User struct {
	ID   int
	Name string
	Age  int
}

func InitDB() {
	fmt.Println("连接数据库")
	dsn := "root:11165wzj@tcp(124.223.8.237)/uvdream?charset=utf8mb4&parseTime=True"
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("连接成功")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func QueryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
