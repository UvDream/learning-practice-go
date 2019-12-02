package main

import "fmt"

func main() {
	// map,需要初始化后才能使用
	var m1 map[string]int
	// 初始化map
	m1 = make(map[string]int, 10)
	m1["北京"] = 1
	m1["上海"] = 2
	fmt.Println(m1)
	fmt.Println(m1["北京"])
	// 判断某个键是否存在
	v, ok := m1["呵呵"]
	if !ok {
		fmt.Println("查无此地")
	} else {
		fmt.Println(v)
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 删除
	delete(m1, "上海")
	fmt.Println(m1)
	delete(m1, "呵呵") //删除不存在的key
	
}
