package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student
)

func showAllstudent() {
	fmt.Println("获取全部学生")
	for k, v := range allStudent {
		fmt.Printf("学号:%d,姓名:%s\n", k, v.name)
	}
}

//构造函数造学生
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	//造学生
	newStu := newStudent(id, name)
	allStudent[id] = newStu

}
func deleteStudent() {
	var deleteID int64
	fmt.Print("请输入要删除的学生学号:")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}

// 函数版本学生系统
func main() {
	allStudent = make(map[int64]*student, 48)
	for {
		//1.打印菜单
		fmt.Println("欢迎来的学生管理系统")
		fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出
	`)
		fmt.Print("请输入你要干啥:")
		//2.等待用户选择要做啥
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//3.执行对应函数
		switch choice {
		case 1:
			showAllstudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}
}
