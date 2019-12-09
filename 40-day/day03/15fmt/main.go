package main

import "fmt"

func main() {
	var (
		name string
		age  int
		show bool
	)
	// 1
	// fmt.Scan(&name, &age, &show)
	// fmt.Println("扫描结果", name, age, show)
	// 2
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &show)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, show)
	// 3
	// fmt.Scanln(&name, &age, &show)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, show)

}
