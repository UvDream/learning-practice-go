package main

import "os"

//关键字defer用于注册延迟调用,这些调用直到ret前被执行,通常用于释放资源或错误处理
func test() error{
	f,err :=os.Create("test.txt")
	if err != nil {return err}
	defer f.Close()		//注册调用,而不是注册函数.必须提供参数,哪怕为空
	f.WriteString("Hello,World!")
	return nil
}
//多个defer注册,按FIL次序执行,哪怕函数或某个延迟调用发生错误,这些调用依旧会被执行
func test1(x int){
	defer println("a")
	defer println("b")
	defer func(){
		println(100 / x)
	}()
	defer  println("c")
}
//延迟调用参数在注册时求值或复制,可用指针或闭包"延迟"读取.
func test2()  {
	x,y:=10,20
	defer func(i int){
		println("defer:",i,y)//y闭包引用
	}(x)//x被复制
	x +=10
	y +=100
	println("x=",x,"y=",y)
}
//滥用defer可能导致性能问题,尤其在一个"大循环"里面
func main() {
	test()
	//test1(0)
	test2()
}
