package main

import "fmt"

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	// 函数内部不能声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 2)
	// 立即执行函数
	func(x, y int) {
		fmt.Println(x, y)
	}(100, 200)

}
