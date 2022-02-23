package main

import "fmt"

// 结构体嵌套
type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address //匿名嵌套
	workPlace
}
type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "张三",
		age:  12,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.address.city)
}
