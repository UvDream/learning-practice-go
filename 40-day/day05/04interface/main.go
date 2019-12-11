package main

import "fmt"

type car interface {
	run()
}
type falali struct {
	brand string
}
type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Println("法拉利700迈")
}
func (b baoshijie) run() {
	fmt.Println("保时捷900迈")
}
func drive(c car) {
	c.run()
}
func main() {
	var f1 = falali{
		brand: "法拉第",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}
