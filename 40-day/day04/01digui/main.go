package main

import "fmt"

// 递归
func f(n uint64) uint64 {
	// 出口
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 走台阶
func f1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f1(n-1) + f1(n-2)
}

func main() {
	ret := f(5)
	fmt.Println(ret)
	m := f1(4)
	fmt.Println(m)
}
