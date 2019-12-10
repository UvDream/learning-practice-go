package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" db:"name" ini:"name"` //实现json里面小写
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "张三",
		Age:  300,
	}
	p, err := json.Marshal(p1) //序列化:go语言结构体变量==>json格式的字符串
	//反序列化:json格式的字符串==>go语言中能识别的结构体变量
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(p))
	fmt.Println(string(p))
	// 反序列化
	str := `{"name":"张三","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
