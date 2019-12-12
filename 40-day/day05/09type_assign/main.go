package main

import "fmt"

// 类型断言
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("不是string")
	}
}
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)

	case bool:
		fmt.Println("bool", t)
	case int64:
		fmt.Println("int64", t)
	}
}
func main() {
	assign(199)
	assign2(100)
}
