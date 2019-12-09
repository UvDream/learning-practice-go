package main

import "fmt"

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	ret := adder()
	ret2 := ret(200)
	fmt.Println(ret2)
}
