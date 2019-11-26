package main

import "fmt"

//不支持嵌套(nested).重载(overload)和默认参数(default parameter)
//1.无需声明原型
//2.支持不定长变参
//3.支持多返回值
//4.支持命名返回参数
//5.支持匿名函数和闭包

//func test(x,y int,s string)(int,string){//类型相同的相邻参数可合并
//	n:=x+y		//多返回值必须用括号
//	return n,fmt.Sprintf(s,n)
//}

//函数是第一类对象,可作为参数传递,建议将复杂签名定义为函数类型
func test(fn func() int)int{
	return fn()
}
type FormatFunc func(s string,x,y int) string
func format(fn FormatFunc,s string,x,y int) string{
	return fn(s,x,y)
}
func main() {
	s1:=test(func() int {
		return 100
	})
	s2:=format(func(s string, x, y int) string {
		return fmt.Sprintf(s,x,y)
	},"%d, %d",10,20)
	println(s1,s2)
//	有返回值的函数,必须有明确的终止语句,否则会引发编译错误.
}
