package main

import "fmt"

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	fmt.Println(m1)
	m1["name"] = "张三"
	m1["age"] = 200
	m1["married"] = false
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
}
