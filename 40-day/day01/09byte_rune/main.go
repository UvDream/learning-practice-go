package main

import (
	"fmt"
)

func main()  {
	s := "helloこんにちは"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	// 字符串修改
	s2:="卞萝卜"
	s3:=[]rune(s2)//字符串强制转换为rune
	s3[0]='红'//字符
	fmt.Println(string(s3))
	// 类型转换
	n:=10
	var f float64
	f=float64(n)
	fmt.Println(f)
	fmt.Printf("%T\n",f)

}