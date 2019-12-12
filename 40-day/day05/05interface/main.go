package main

import "fmt"

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫")
}
func (c cat) eat(e string) {
	fmt.Printf("猫吃----%s\n", e)

}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡")
}
func (c chicken) eat(e string) {
	fmt.Printf("鸡吃----%s\n", e)
}

func main() {
	var a1 animal
	bc := cat{
		name: "小猫",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
	bc.eat("猫粮")
	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
}
