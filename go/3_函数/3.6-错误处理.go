package main

import "fmt"

//没有结构化异常,使用panic抛出错误,recover捕获错误!
func test(){
	defer func(){
		if err := recover();err!=nil{
			println(err.(string))
		}
	}()
	panic("panic error!")
}
//由于panic,recover参数类型为interface{},因此可抛出任何类型对象
//延迟调用中引发的错误,可被后续延迟调用捕获,但仅最后一个错误可被捕获
func test2(){
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

//捕获函数recover只有在延迟调用内直接调用才会终止错误,否则总是返回nil,任何未捕获的错误都会沿用堆栈向外传递
func test3()  {
	defer recover()  //无效
	defer fmt.Println(recover())//无效
	defer func() {
		println("defer inner")
		recover()   //无效
	}()

}
func main() {
	test()
	test2()
}
