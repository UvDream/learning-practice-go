package main

import "fmt"

// go语言如果标识符首字母大写就表示对外部可见

//Dog 这是一个狗的结构体//注释
// type Dog struct {
// 	name string
// }
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法作用于特定的函数
func (d dog) wang() {
	fmt.Println("wangwangwang", d.name)
}

func main() {
	d1 := newDog("aaa")
	d1.wang()
}
