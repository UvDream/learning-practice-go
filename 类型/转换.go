package main

//引用类型包括slice,map和channel,它们有复杂的内部结构,除了申请内存外,还需要初始化相关属性

//类型转换,不支持隐式类型转换,即便是从窄向宽转换也不行
var b byte = 100

//var n int=b 隐式转换
var n int = int(b) //显式转换
func main() {
	//	内置函数new计算类型大大小,为其分配零值内存,返回指针,而make会被编译成具体的创建函数,由其分配内存和初始化成员结构,返回对象而非空指针
	a := []int{0, 0, 0} //提供初始化表达式
	a[1] = 10

	b := make([]int, 3) //make slice
	b[1] = 10

	c := new([]int)
	c[1] = 10 //error
}
