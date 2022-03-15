package mysqls

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMysql() (err error) {
	fmt.Println("连接mysql")
	dsn := "root:11165wzj@tcp(124.223.8.237)/uvdream"
	DB, err = sql.Open("mysql", dsn)
	fmt.Println(DB, err)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	fmt.Println("连接成功")
	//defer  db.Close()
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 单行查询
func QueryRow() {
	fmt.Println("查询一行")
	sqlStr := "select id,name,age from user where id=?"
	var u user
	err := DB.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("查询失败", err)
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
}

//多行查询
func QueryMultiRow() {
	fmt.Println("查询多行参数")
	sqlStr := "select id,name,age from user where id>?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("查询失败", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}

}

//插入
func InsertRow() {
	sqlStr := "insert into user(name,age) values(?,?)"
	ret, err := DB.Exec(sqlStr, "小明", 18)
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	ID, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("获取id失败", err)
		return
	}
	fmt.Println("插入成功", ID)
}

//更新数据
func UpdateRow() {
	sqlStr := "update user set name=? where id=?"
	res, err := DB.Exec(sqlStr, "小红", 1)
	if err != nil {
		fmt.Println("更新失败", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取影响行数失败", err)
		return
	}
	fmt.Println("更新成功", n)
}

//删除数据
func DeleteRow() {
	sqlStr := "delete from user where id=?"
	res, err := DB.Exec(sqlStr, 1)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取影响行数失败", err)
		return
	}
	fmt.Println("删除成功", n)
}

// 预处理
//1.把sql语句分成两部分,命令部分与数据部分
//2.先把命令部分发送给mysql服务端,mysql服务端进行sql预处理
//3.然后把数据部分发送给mysql服务端,mysql服务端对sql语句进行占位符替换
//4.mysql服务端执行完整的sql语句并将结果返回给客户端

func Pretreatment() {
	fmt.Println("预处理")
	sqlStr := "select id,name,age from user where id>?"
	stmt, err := DB.Prepare(sqlStr)
	fmt.Println(stmt, err)
	if err != nil {
		fmt.Println("预处理失败", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("查询失败", err)
			return
		}
		fmt.Println("id:", u.id, "name:", u.name, "age:", u.age)
	}

}

func PretreatmentInsert() {
	sqlStr := "insert into user(name,age) values(?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理失败", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小红", 18)
	if err != nil {
		fmt.Println("插入失败", err)
	}
	_, err = stmt.Exec("小明", 18)
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	fmt.Println("插入成功")
}

//事务
func Transaction() {
	tx, err := DB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
