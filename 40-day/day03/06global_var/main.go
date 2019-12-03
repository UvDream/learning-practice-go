package main

import "fmt"

// 变量的作用域
var x = 100

func f1() {
	// 函数中查找变量的顺序
	// 1.在自身函数内部查找
	// 2.找不到就在函数外部查找,一直找到全局
	fmt.Println(x)
}
func main() {
	f1()
}
