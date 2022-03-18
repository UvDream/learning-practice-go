package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	JSONStruct()
	TagStruct()
}
func JSONStruct() {
	type Person struct {
		Name   string
		Age    int64
		Weight float64
	}
	p1 := Person{
		Name:   "张三",
		Age:    18,
		Weight: 60.5,
	}
	fmt.Println(p1)
	p2, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("json.Marshal failed,err:", err)
		return
	}
	fmt.Println(p2)
	var p3 Person
	err = json.Unmarshal(p2, &p3)
	if err != nil {
		fmt.Println("json.Unmarshal failed,err:", err)
		return
	}
	fmt.Println(p3)
	fmt.Println(p3.Age)
}
func TagStruct() {
	fmt.Println("-----------------TagStruct-----------------")
	type Person struct {
		Name   string  `json:"name"`
		Age    int64   `json:"age"`
		Weight float64 `json:"-"`
	}
	p1 := Person{
		Name:   "张三",
		Age:    18,
		Weight: 60.5,
	}
	p2, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("json.Marshal failed,err:", err)
		return
	}
	fmt.Println(p2)
}
