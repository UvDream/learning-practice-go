package main

import (
	"fmt"
)

func main() {
	// make()函数创造切片
	s1 := make([]int, 5, 10)
	fmt.Printf("s=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 // s3和s4都指向同一个底层数组
	s4[0] = 100
	fmt.Println(s3, s4) //[100 3 5] [100 3 5]
	s3[0] = 200
	fmt.Println(s3, s4) //[200 3 5] [200 3 5]

	// 切片的遍历
	// 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// for range
	for _, v := range s3 {
		fmt.Println(v)
	}
}
