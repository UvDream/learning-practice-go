package main

import "fmt"

// defer 延迟到函数即将返回的时候再执行
// 1.返回值赋值
// defer
// 2.真正的RET返回
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("呵呵")
	defer fmt.Println("嘿嘿")
	fmt.Println("end")
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=y=x=5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值=x=5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
