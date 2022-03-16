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

//查询
func QueryRow() {
	fmt.Println("-----------------QueryRow-----------------")
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
func QueryMultiRow() {
	fmt.Println("-----------------QueryMultiRow-----------------")
	sqlStr := "select id, name, age from user where id>?"
	var user []User
	err := db.Select(&user, sqlStr, 1)
	if err != nil {
		fmt.Println("查询失败")
	}
	fmt.Println(user)
}

//插入
func InsertRow() {
	fmt.Println("-----------------InsertRow-----------------")
	sqlStr := "insert into user(name,age) values(?,?)"
	res, err := db.Exec(sqlStr, "新冠病毒", 18)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("查询插入数据的ID失败")
	}
	fmt.Println("插入数据的ID:", id)
}

//更新数据
func UpdateRow() {
	fmt.Println("-----------------UpdateRow-----------------")
	sqlStr := "update user set name=? where id=?"
	res, err := db.Exec(sqlStr, "新冠病毒", 2)
	if err != nil {
		fmt.Println("更新失败")
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("查询影响行数失败")
		return
	}
	fmt.Println("影响行数:", n)
}

//删除数据
func DeleteRow() {
	sqlStr := "delete from user where id=?"
	res, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("查询影响行数失败")
		return
	}
	fmt.Println("影响行数:", n)
}
