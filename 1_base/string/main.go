package main

import (
	"strings"
	"fmt"
)

func main(){
	path:="D:/user/local"
	fmt.Println(path)
	// 多行字符串
	s2:=`
第一行
第二行
第三行
	`
	fmt.Println(s2)
	// 字符串相关操作
	fmt.Println("长度",len(s2))
	// 字符串拼接
	name:="张三"
	status:="很帅"
	ss:=name+status
	fmt.Println(ss)
	ss1:=fmt.Sprintf("%s%s",name,status)
	fmt.Println(ss1)
	// 字符串分割
	s3:=`D:\\user\\local`
	ret:=strings.Split(s3,`\\`)
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss,"张三"))
	// 前缀/后缀判断
	fmt.Println(strings.HasPrefix(ss,"张三"))
	fmt.Println(strings.HasSuffix(ss,"张三"))
	// 字符串位置
	s4:="abcdefc"
	// 最先出现
	fmt.Println(strings.Index(s4,"c"))
	// 最后出现
	fmt.Println(strings.LastIndex(s4,"c"))
	// 拼接(先分割后拼接)
	fmt.Println(strings.Join(ret,"+"))



	// travelString()

}
func travelString()  {
	s:="aa你们好的呀!"
	for i:=0;i<len(s);i++{
		fmt.Printf("%v(%c)",s[i],s[i])
	}
	fmt.Println()
	for _,r:=range s{
		fmt.Printf("%v(%c)",r,r)
	}
}