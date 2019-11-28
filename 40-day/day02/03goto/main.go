package main

import (
	"fmt"
)


func main()  {
// 跳出循环
	var flag=false
	for i:=0;i<10;i++{
		for j:='A';j<'Z';j++{
			if j=='C'{
				flag=true
				break //跳出内层循环
			}
			fmt.Printf("%v-%c\n",i,j)
		}
		if flag{
			break
		}
	}
	fmt.Println("goto")
	for i:=0;i<10;i++{
		for j:='A';j<'Z';j++{
			if j=='C'{
				goto xx
			}
			fmt.Printf("%v-%c\n",i,j)
		}
		
	}
xx:
	fmt.Println("结束")

}