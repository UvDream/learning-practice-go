package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f2()
}

// 替换
func f2() {
	// 读取操作文件
	fileObj, err := os.OpenFile("./40-day/day06/02file/log.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("读取失败,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// 读取缓冲文件
	tempFile, err := os.OpenFile("./40-day/day06/02file/log1.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("读取失败,err:%v\n", err)
		return
	}
	defer tempFile.Close()

	// 读源文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("读取失败,err:%v\n", err)
		return
	}
	// 写入临时文件
	tempFile.Write(ret[:n])
	//再写入插入文件
	var s []byte
	s = []byte{'c'}
	tempFile.Write(s)
	// 接着把源文件内容写入临时文件
	var x [1024]byte
	for {
		m, err := fileObj.Read(x[:])
		if err == io.EOF {
			tempFile.Write(x[:m])
			break
		}
		if err != nil {
			fmt.Printf("读取失败,err:%v\n", err)
			return
		}
		tempFile.Write(x[:m])
	}
	// 源文件后续写入临时文件
	fileObj.Close()
	tempFile.Close()
	// 重命名
	os.Rename("./40-day/day06/02file/log1.txt", "./40-day/day06/02file/log.txt")
}
