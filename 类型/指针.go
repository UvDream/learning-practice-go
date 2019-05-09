package main

import (
	"fmt"
	"unsafe"
)

//返回局部变量指针是安全的,编译器会根据需要将其分配在GC Heap上
func test() *int {
	x := 100
	return &x //在对上分配x内存,但在内联时,也可能直接分配在目标栈
}

//指针类型*,指针的指针**,以及包含包名前缀的*<package>.T
//默认值为nil,没有NULL常量
//操作符"&"取变量地址,"*"透过指针访问目标对象
//不支持指针运算,不支持"->"运算符,直接用"."访问目标成员
func main() {
	type data struct {
		a int
	}
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p,%v\n", p, p.a) //直接用指针访问目标对象成员,无须转换
	//不能对指针做加减运算
	//可以在unsafe.Pointer和任意类型指针间进行换算
	x := 0x12345678
	p1 := unsafe.Pointer(&x) //*int ->Pointer
	n := (*[4]byte)(p1)      //Pointer ->*[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X--", n[i])
	}

	//	将Pointer转换成uintptr可变相实现指针计算
	x1 := struct {
		s string
		k int
	}{"abc", 100}
	y := uintptr(unsafe.Pointer(&x1)) //*struct -> Pointer ->uintptr
	y += unsafe.Offsetof(x1.k)        //uintptr + offset
	y2 := unsafe.Pointer(y)           //uintptr ->Pointer
	yx := (*int)(y2)                  //Pointer ->*int
	*yx = 200
	fmt.Printf("\n%#v\n", x1)
	//GC把uintptr当成普通整数对象,它无法阻止"关联"对象被回收
}
