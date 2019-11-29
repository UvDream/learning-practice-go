package main

import (
	"fmt"
)

func main()  {
	// 数组
	var a1[3]bool
	fmt.Println(a1)
	// 数组的初始化
	//如果不初始化:默认值都是零值
	// 1.初始化方式1
	a1=[3]bool{true,true,true}
	fmt.Println(a1)
	// 2.初始化方式2
	// 根据初始化自动推断数组的长度
	a100:=[10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(a100)
	// 3.初始化方式3(根据索引)
	a3:=[5]int{0:1,2:6}
	fmt.Println(a3)
	// 数组遍历
	cityS := [...]string{"北京", "上海", "深圳"}
	// 1.根据索引遍历
	for i:=0;i<len(cityS);i++{
		fmt.Println(cityS[i])
	}
	// for range遍历
	for i,v:=range cityS{
		fmt.Println(i,v)
	}
}