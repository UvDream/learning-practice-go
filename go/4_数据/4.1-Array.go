package main

import "fmt"

//数组
//1.数组是值类型
//2.数组长度必须是常量,且是类型的组成部分.[2]int和[3]int是不同类型
//3.支持"==","!="操作符,因为内存总是被初始化过的
//4.指针数组[n]*,数组指针*[n]




func main() {
	//可用复合语句初始化
	a :=[3]int{1,2}			//未初始化元素值为0
	b :=[...]int{1,2,3,4}  //通过初始化之确定数组长度
	c :=[5]int{2:100,4:200}//使用索引号初始化元素
	d :=[...]struct{
		name string
		age uint8
	}{
		{"user1",10},//可省略元素类型
		{"user2",20},//别忘记最后一行的逗号
	}
	//支持多维数组
	a1 :=[2][3]int{{1,2,3},{4,5,6}}
	b1 :=[...][2]int{{1,1},{2,2},{3,3}}
	fmt.Printf("b: %p\n",&b)
	fmt.Printf("c: %p\n",&c)
	fmt.Printf("d: %p\n",&d)
	fmt.Printf("a1: %p\n",&a1)
	fmt.Printf("b1: %p\n",&b1)

	fmt.Printf("a: %p\n",&a)
	a2 :=[2]int{}
	test(a2)
	fmt.Println(a2)
	//内置函数len和cap都返回数组长度(元素数量)
	println("a的长度",len(a),cap(a))
}
//值拷贝行为会造成性能问题,通常会建议使用slice,或数组指针
func test(x [2]int){
	fmt.Printf("x: %p\n",&x)
	x[1]=1000
}
