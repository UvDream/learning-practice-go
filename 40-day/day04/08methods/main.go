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
type person struct {
	name string
	age  int
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法作用于特定的函数
func (d dog) wang() {
	fmt.Println("wangwangwang", d.name)
}

// 值类型(p person)为拷贝
func (p person) guonian() {
	p.age++
}

// (p *person)指针接受者
func (p *person) zGuoNian() {
	p.age++
}
func (p *person) dream() {
	fmt.Println("不上班也能赚钱")
}
func main() {
	d1 := newDog("aaa")
	d1.wang()
	p1 := newPerson("张三", 18)
	fmt.Println(p1)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zGuoNian()
	fmt.Println(p1.age)
}
