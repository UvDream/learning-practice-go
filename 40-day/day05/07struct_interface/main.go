package main

import "fmt"

// 同一个结构体可以实现多个接口
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("猫")
}
func (c *cat) eat(s string) {
	fmt.Printf("猫吃---:%s", s)
}

func main() {
	var a animal
	c1 := cat{name: "猫", feet: 1}
	a = &c1
	fmt.Println(a)
	a.eat("猫粮")

}
