package html

import (
	"fmt"
)

// 跳出for循环
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("结束")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("结束")
	//	for range
	ins := []int{1, 2, 3, 4, 5}
	for i := range ins {
		fmt.Println(i)
	}
}
