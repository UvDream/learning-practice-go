package main

import (
	"fmt"
)

// if
func main() {
	age := 19
	if age > 18 {
		fmt.Println("你都成年了!")
	} else {
		fmt.Println("你还未成年")
	}
	// for range
	s := "hello南京"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
	// 9*9乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, i*j)
		}
		fmt.Println()
	}
}
