//常量
package main

import "unsafe"

//常量的值必须是编译期可确定的数字,字符串,布尔值
const x, y int = 1, 2    //多常量初始化
const s = "Hello,World!" //类型推断
const (                  //常量组
	a, b      = 10, 100
	c    bool = false
)
const (
	s1 = "abc"
	x1 //在常量组中,如不提供类型和初始值,那么视作与上一常量相同
)
const ( //常量值还可以是len,cap,unsafe.Sizeof等编译期可确定结果的函数返回值
	a1 = "abc"
	b1 = "len(a1)"
	c1 = unsafe.Sizeof(b1)
)
const ( //如果常量类型足以储存初始值,那么不会引发溢出的错误
	a2 byte = 100
	b2 int  = 1e20
)

//枚举
const (
	Sunday    = iota //0   关键字iota定义常量组中从0开始按行计数的自增枚举值
	Monday           //1
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)
const ( //在同一常量组中,可以提供多个iota,它们各自增长
	A, B = iota, iota << 10 //0,0<<10
	C, D                    //1,1<<10
)
const ( //如果iota自增被打断,须显式回复
	A = iota //0
	B        //1
	C = "c"  //c
	D        //c,与上一行相同
	E = iota //4,显式回复,注意计数包含了C,D两行
	F        //5
)

//可以通过自定义类型来实现枚举类型限制
type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {}
func main() {
	const x = "xxx" //未使用的局部变量不会引起编译错误
	c := Black
	test(c)
	x1 := 1
	test(x1) //error
	test(1)  //常量会被编译器自动转换
}
