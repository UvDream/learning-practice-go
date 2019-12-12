package main

import "fmt"

// 使用值接受者和指针接受者的区别
type animal interface {
	move()
}
type cat struct {
	name string
	feet int8
}

// 使用值接受者实现
// func (c cat) move() {
// 	fmt.Println("猫")
// }

// 使用指针实现
func (c *cat) move() {
	fmt.Println("猫")
}

func main() {
	var a animal
	c1 := cat{"tom", 4}
	//a = c1 //不可以接受c1
	a = &c1
	fmt.Println(a)        //&{tom 4}
	fmt.Printf("%T\n", a) //*main.cat
}
