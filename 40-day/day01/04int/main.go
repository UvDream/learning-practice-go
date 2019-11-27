package main

import (
	"fmt"
)

func main()  {
	
	var i1=101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%b\n",i1) //十进制转换二进制
	fmt.Printf("%o\n",i1) //十进制转换八进制
	fmt.Printf("%x\n",i1)//十进制转换为十六进制
	// 八进制
	i2:=077
	fmt.Printf("%d\n",i2)
	// 十六进制
	i3:=0x1234567
	fmt.Printf("%d\n",i3)
	// 查看变量类型%T
	fmt.Printf("%T\n",i3)
	// 声明init8类型的变量
	i4:=int8(9)//明确指定类型,否则默认为int
	fmt.Printf("%T\n",i4)
}