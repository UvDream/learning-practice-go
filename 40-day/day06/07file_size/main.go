package main

import (
	"fmt"
	"os"
)

//	1.文件对象的类型
//	2.获取文件对象的详细信息
func main() {
	fileObj, err := os.Open("./40-day/day06/07file_size/main.go")
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	// 文件对象
	fmt.Printf("%T\n", fileObj)
	// 获取文件信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("获取失败,err:%v\n", err)
		return
	}
	fmt.Printf("文件大小是:%dB\n", fileInfo.Size())
}
