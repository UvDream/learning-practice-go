package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3 cap(s1):3
	// s1[3]="南京" //错误写法,索引越界
	// append,调用必须使用原来的切片接受返回值,底层数组放不下的时候,原来的底层数组换一个位置3->6->12
	s1 = append(s1, "南京")
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):4 cap(s1):6
	fmt.Println(s1)                                         //[北京 上海 深圳 南京]
	// 多个元素
	s1 = append(s1, "合肥", "杭州", "无锡")
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1)) //len(s1):7 cap(s1):12
	// 追加切片
	a := []string{"成都", "重庆"}
	s1 = append(s1, a...) //a...拆开切片
	fmt.Println(s1)       //[北京 上海 深圳 南京 合肥 杭州 无锡 成都 重庆]

}
