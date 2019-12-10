package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("张三", 11)
	fmt.Println(p1.name)
}
