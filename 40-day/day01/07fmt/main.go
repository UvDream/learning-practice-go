package main

import (
	"fmt"
)
var a =100

func main()  {
	// 查看类型
	fmt.Printf("%T\n",a)
	fmt.Printf("%v\n",a)
	fmt.Printf("%b\n",a)
	fmt.Printf("%d\n",a)
	fmt.Printf("%o\n",a)
	fmt.Printf("%x\n",a)
	var s="Hello"
	fmt.Printf("%s",s)
}