package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() 失败\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	// 拿到路径最底层的文件
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f1()
}
