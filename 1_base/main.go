package main

import (
	"fmt"
)
var (
	name string
	age int
	isOk bool
)

func main(){
	name="张三"
	age=18
	isOk=true
	fmt.Print(isOk,"\n")
	fmt.Printf("name:%s\n",name)	//%s:占位符,使用name这个变量的值去替换占位符
	fmt.Println(age)			//打印完指定的内容之后会在后面加上一个换行符
	// 声明变量同时赋值
	var s1 string="张三" //不规范
	fmt.Println(s1)
	// 类型推导
	var s2="王五"
	fmt.Println(s2)
	// 短变量声明
	s3:="李四"
	fmt.Println(s3)
}