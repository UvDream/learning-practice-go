package main

import "fmt"

// 结构体是值类型
type person struct {
	name string
	age  int
}

func f(x person) {
	x.age = 20
}
func f2(x *person) {
	//(*x).age = 30 //根据内存地址找到那个变量,修改的谁原来的变量
	x.age = 30 //语法糖,自动根据指针找对应变量
	fmt.Println("f2", *x)
}
func main() {
	var p person
	p.name = "呵呵"
	p.age = 10
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println("f2地址", p)
	var m = new(person) 
	fmt.Println(m)
}
