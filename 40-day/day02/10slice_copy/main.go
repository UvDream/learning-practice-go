package main

import "fmt"

func main() {
	// copy
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            // 使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]
	// 删除元素
	// 删除索引为1的元素(其实就是就是拼接原理),切完容量不变
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)                                         //[100 5]
	fmt.Printf("len(a1):%d cap(a1):%d\n", len(a1), cap(a1)) //len(a1):2 cap(a1):3

	// 切片容量
	fmt.Println("切片容量")
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Printf("%p\n", s1)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	s1 = append(s1[:1], s1[2:]...) //修改了底层数组
	fmt.Printf("%p\n", s1)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) // len(s1):2 cap(s1):3
	fmt.Println(x1)                                         // [1 5 5]
	// 通过切片可以修改底层数组
}
