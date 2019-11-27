package main

import "fmt"

//匿名函数可复制给变量,作为结构字段,或者在channel里传送

func main() {
	//function variable
	fn :=func(){println("Hello world!")}
	fn()

	//function collection
	fns :=[](func(x int)int){
		func(x int) int {
			return x+1
		},
		func(x int) int {
			return x+2
		},
	}
	println(fns[0](100))

	//function as field
	d:= struct {
		fn func() string
	}{
		fn: func() string {return "Hello,World!"},
	}
	println(d.fn)

	//channel of function
	fc:=make(chan func() string,2)
	fc <- func() string {return "Hello,world!"}
	println((<-fc)())

//	闭包复制的是原对象指针,这就很容易解释延迟引用现象
	f:=test()
	f()
}
func test() func(){
	x:= 100
	fmt.Printf("x (%p)=%d\n",&x,x)
	return func() {
		fmt.Printf("x (%p)=%d\n",&x,x)
	}
}
