package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("这是f1")
	f()
}
func f2(x, y int) {
	fmt.Println("这是f2")
	fmt.Println(x + y)
}

// 定义一个函数对函数包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		fmt.Println("1111")
		f(m, n)
	}
	return tmp
}
func main() {
	f1(f3(f2, 2, 1))
}
