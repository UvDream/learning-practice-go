package main

import (
	"fmt"
)

func main() {
	var ages [30]int
	// 1
	ages = [30]int{1, 2, 3}
	fmt.Println(ages)
	// 2
	var ages2 = [...]int{12, 3, 4}
	// 3
	var ages3 = [...]int{1: 100, 10: 200}
	fmt.Println(ages2)
	fmt.Println(ages3)
	var a1 = [...][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	fmt.Println(a1)
	// 数组值类型
	x := [3]int{1, 2, 3}
	y := x     //拷贝
	y[1] = 200 //修改的是y,不影响x
	f1(x)
	fmt.Println(x) //[1,2,3] 函数传参为拷贝

	fmt.Println("切片")
	// 切片slice
	var s1 []int
	// 1.初始化
	s1 = []int{10, 2, 3}
	fmt.Println(s1)
	// 2.make初始化
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)
	s3 := []int{1, 4, 6}
	s4 := s3
	fmt.Println(s3, s4)
	s4[1] = 200
	fmt.Println(s3, s4)
	// append
	s3 = append(s3, 200)
	fmt.Println("append", s3, s4)
	// copy
	var s5 = make([]int, 4, 4)
	copy(s5, s3)
	fmt.Println("copy", s3, s5)

	// 指针
	fmt.Println("指针")
	addr := "北京"
	addrP := &addr
	fmt.Println(addrP)
	fmt.Printf("%T\n", addrP)
	addrV := *addrP
	fmt.Println(addrV)

	// map
	fmt.Println("map")
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	fmt.Println(m1)
	m1["呵呵"] = 1
	fmt.Println(m1)

}
func f1(a [3]int) {
	a[1] = 100
}
