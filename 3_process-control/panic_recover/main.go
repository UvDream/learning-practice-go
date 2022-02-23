package main

import "fmt"

func funcA() {
	fmt.Println("A")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover B")
		}
	}()
	panic("出现严重错误!") //程序崩溃退出

}
func funcC() {
	fmt.Println("C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
