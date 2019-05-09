package main

import "fmt"

//slice并不是数组或者数组指针,他通过内部指针和相关属性引用数组片段,以实现变长方案
//1.引用类型,但自身是结构体,值拷贝传递
//2.属性len表示可用元素数量,读写操作不能超过该限制
//3.属性cap表示最大扩张容量,不能超出数组限制
//4.如果Slice == nil,namelen,cap结果都等于0
func main() {
	data :=[...]int{0,1,2,3,4,5,6,7,8,9}
	//slice :=data[1:4:5]
	data1 :=data[:6:8]
	data2 :=data[5:]
	data3 :=data[:3]
	data4 :=data[:]
	fmt.Println(":6:8是:",data1)
	fmt.Println("5:是",data2)
	fmt.Println(data3)
	fmt.Println(data4)
//	读写操作实际目标是底层数组,只需要注意索引号的差别
	s :=data[2:4]
	s[0] +=100
	s[1] +=200
	fmt.Println(s)
	fmt.Println(data)
	s1 :=[]int{0,1,2,3,8:100}//通过初始化表达式构造,可使用索引号
	fmt.Println(s1,len(s1),cap(s1))
	s2 :=make([]int,6,8)//使用make创建,指定len和cap值
	fmt.Println(s2,len(s2),cap(s2))
	s3 :=make([]int,6)//省略cap,相当于cap=len
	fmt.Println(s3,len(s3),cap(s3))
//	使用make动态创建slice,避免了数组必须用常量做长度的麻烦,还可以用指针直接访问底层数组,退化成普通数组操作
	a :=[]int{0,1,2,3}
	p :=&a[2]
	*p +=100
	fmt.Println(a)
//	至于[][]T,是指元素类型为[]T
//	datas:=[][]int{
//		[]int{1,2,3},
//		[]int{100,200},
//		[]int{11,22,33,44},
//	}
//	可直接修改struct array/slice成员
	d :=[5]struct{
		x int
	}{}
	e :=d[:]
	d[1].x=10
	e[2].x=20
	fmt.Println("测试")
	fmt.Println(d)
	fmt.Printf("%p,%p\n",&d,&d[0])
}