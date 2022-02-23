package main

import "fmt"

// 结构体
// 定义
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var w person
	// 实例化
	w.name = "wzj"
	w.age = 18
	w.gender = "男"
	w.hobby = []string{"篮球", "足球"}
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	// 访问名字
	fmt.Println(w.name)

	// 匿名结构体
	fmt.Println("匿名结构体")
	var s struct {
		name string
		age  int
	}
	s.name = "张三"
	s.age = 19
	fmt.Println(s)
}
