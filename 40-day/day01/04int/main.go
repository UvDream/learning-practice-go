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

}