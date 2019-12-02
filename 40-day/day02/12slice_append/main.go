package main

import "fmt"

func main() {
	// 关于append删除切片中某个元素
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]
	// 删除索引为1
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
