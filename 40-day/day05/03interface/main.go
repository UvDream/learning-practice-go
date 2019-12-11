package main

import "fmt"

// 引出接口的实例
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}
type person struct {
}
type cat struct {
}
type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵猫")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func (p person) speak() {
	fmt.Println("救命")
}

// 接受不定类型调用
func play(x speaker) {
	// 接受一个参数,传入进来就打印啥
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	play(c1)
	play(d1)
	play(p1)
}
