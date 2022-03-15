package main

import (
	"fmt"
	"hello/4_function/service"
)

func main() {
	r := sum(10, 20)
	fmt.Println(r)
	sum1(1, 2)
	r2 := sum2()
	fmt.Println(r2)
	r3 := sum3(4, 1)
	fmt.Println(r3)
	r4, r5 := sum4()
	fmt.Println(r4, r5)
	r6 := sum6(1, 7)
	fmt.Println(r6)
	sum7("下雨了", 1, 2, 3)
	service.HtmlFunc()

}
func sum(x int, y int) (ret int) {
	return x + y
}
func sum1(x int, y int) {
	fmt.Println(x + y)
}
func sum2() int {
	return 10
}

// 返回值命名
func sum3(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func sum4() (int, string) {
	return 1, "北京"
}

// 多类型简写
func sum6(x, y int) int {
	return x + y
}

// 可变长参数
func sum7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //[1 2 3]

}
