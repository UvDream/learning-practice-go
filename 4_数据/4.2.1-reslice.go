package main

import "fmt"

//所谓的reslice,是基于已有slice创建新slice对象,以便在cap允许范围内调整属性
func main() {
	s:=[]int{0,1,2,3,4,5,6,7,8,9}
	s1:=s[2:5]
	s2:=s1[2:6:7]
	//s3:=s2[3:6]
	fmt.Println(s1)
	fmt.Println(s2)
}
