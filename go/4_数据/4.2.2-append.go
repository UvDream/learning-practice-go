package main

import "fmt"
//append往[]数据后面添加数据
func main() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2:=append(s,1)
	fmt.Printf("%p\n",&s2)
	fmt.Println(s,s2)
	data:=[...]int{0,1,2,3,4,5,6,7,8,9}
	s3:=data[:3]
	s4:=append(s3,100,200)
	fmt.Println(data)
	fmt.Println(s3)
	fmt.Println(s4)
	//	一旦超出原slice.cap限制,就会重新分配底层数组,即便原数组并未填满.
	fmt.Println("超过")
	a:=[...]int{0,1,2,3,4,10:0}
	b:=a[:2:3]
	fmt.Println(b)
	b=append(b,100,200)
	fmt.Println(b)
	fmt.Println(&b[0],&a[0])
}
