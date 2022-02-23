package main

import "fmt"

func main() {
	// 1.& 取地址
	n := 18
	fmt.Println(&n) //0xc000092008
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int
	// 2.* 根据地址取值
	m := *p
	fmt.Println(m)        //18
	fmt.Printf("%T\n", p) //*int

	//var a *int // 空指针
	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)
	

}
