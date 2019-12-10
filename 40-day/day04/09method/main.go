package main

import "fmt"

// 任意类型添加方法
// 不能给别的包的类型添加方法,只能给自己包的类型添加方法

//myInt 将int定义为自定义myInt类型
type myInt int

func (i myInt) hello() {
	fmt.Println("hello")

}
func main() {
	m := myInt(100)
	m.hello()
}
