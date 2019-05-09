//变量
package main

import "fmt"

//go是静态类型语言,不能在运行期间改变变量的类型
var x, y, z int
var s, n = "abc", 123
var (
	a int
	b float32
)

//注意:go语言会将未使用的全局变量当做错误
func main() {
	//在函数内部,可以更简略的":="定义变量
	n, s := 0x1234, "Hello Word!"
	println(x, s, n)
	//	多变量赋值时,先计算所有相关值,然后从左到右赋值
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100
	fmt.Print("\ni的值为:", i)
	fmt.Print("\ndata的值为:", data)
	//	特殊只写变量"_",用户忽略值占位符
	_, s1 := test()
	println("\ns1的值为", s1)
	test2()
}
func test() (int, string) {
	return 1, "abc"
}

//注意重新赋值与定义新同名变量的区别
func test2() {
	s2 := "abc"
	fmt.Print("\n重新赋值和同名")
	println(&s2)
	s2, y := "hello", 20 //重新赋值,与前s在同一层次的代码块中,且有新的变量被定义
	println(&s2, y)      //通常函数多返回值err会被重复使用
	{
		s2, z := 1000, 30 //定义新同名变量,不在同一层次代码块
		println(&s2, z)
	}
}
