package main

import (
	"fmt"
)
// 常量
// 定义了常量之后不能修改
//在程序运行期间不能改变值
const pi=3.1314526
const (
	a=100
	b
	c
)
const(
	n1=iota//0
	n2
)
const(
	a1,a2=iota+1,iota+2
	b1,b2=iota+2,iota+2
	c1,c2
)
func main()  {
	fmt.Println(pi)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(a1,a2)
	fmt.Println(b1,b2)
	fmt.Println(c1,c2)


}