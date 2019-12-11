package main

import (
	student "02student/student"
	"fmt"
	"os"
)

func showMenus() {
	fmt.Println("欢迎来到管理系统!")
	fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.修改学生
	4.删除学生
	5.退出`)
}

// 声明一个全局的学生的管理者
var smr student.Studentmanagement

func main() {
	smr = student.Studentmanagement{
		AllStudent: make(map[int64]student.Student, 100),
	}
	for {
		showMenus()
		fmt.Println("请输入序号:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是:", choice)
		switch choice {
		case 1:
			smr.ShowStudent()
		case 2:
			smr.AddStudent()
		case 3:
			smr.EditStudent()
		case 4:
			smr.DeleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}

}
