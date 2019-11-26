package main
//不能用容器对象接受多返回值.只能用多个变量,或"_"忽略
func test()(int,int){
	return 1,2
}
//多返回值可直接作为其他函数调用实参
func add(x,y int) int{
	return x+y

	//z=x+y  //命名返回参数可看做与形参类似的局部变量,最后由return隐式返回
	//return

	//{				//命名返回参数可被同名局部变量遮蔽,此时需要显示返回
	//	var z=x+y
	//	return z
	//}

	//defer func(){ //命名返回参数允许defer延迟调用通过闭包读取和修改
	//	z +=100
	//}()
	//z=x+y
	//return


}
func sum(n ...int) int{
	var x int
	for _,i:=range n{
		x +=i
	}
	return x
}
func main() {
	x,_:=test()
	println(x)
	println(add(test()))
	println(sum(test()))
}
