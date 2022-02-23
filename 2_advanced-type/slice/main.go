package main

import (
	"fmt"
)

func main() {
	// 切片定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"上海", "北京", "深圳"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 4, 5, 6, 7, 9, 10}
	s3 := a1[0:4]   //基于数组切割左包含,右不含[)
	fmt.Println(s3) //[1 3 4 5]
	s4 := a1[1:6]
	fmt.Println(s4) //[3 4 5 6 7]
	s5 := a1[:4]
	fmt.Println(s5) //[1 3]
	s6 := a1[:]
	fmt.Println(s6) //[1 3 4 5 6 7 9]
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) //len(s5):4 cap(s5):7
	// 底层数组从切片的第一个元素指向到最后一个元素的数量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s4), cap(s4)) //len(s5):5 cap(s5):6

	//	切片再切片
	s8 := s4[3:]
	fmt.Println("这是s4", s4)                                 // [3 4 5 6 7]
	fmt.Println("这是s8", s8)                                 //[6 7]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8)) //len(s8):2 cap(s8):4
}
