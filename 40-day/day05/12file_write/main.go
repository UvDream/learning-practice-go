package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func main() {
	// write()
	// write2()
	write3()
}
func write3() {
	str := "再次来写"
	err := ioutil.WriteFile("./40-day/day05/12file_write/log.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
func write2() {
	fileObj, err := os.OpenFile("./40-day/day05/12file_write/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("又来写了") //写入缓存
	wr.Flush()             //将缓存中的写入文件
}

func write() {
	fileObj, err := os.OpenFile("./40-day/day05/12file_write/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	var str = "字符串\n"
	// 写入文件
	// Write
	fileObj.Write([]byte(str)) //写入字节切片数据
	// WriteString
	fileObj.WriteString("嘻嘻\n") //直接写入字符串数据
}
