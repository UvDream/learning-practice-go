package main

import "fmt"

// 结构体实现继承
type animal struct {
	name string
}

// 给动物实现动的方法

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 狗
type dog struct {
	foot   int8
	animal //继承了animal
}

// 给狗实现叫
func (d dog) wang() {
	fmt.Printf("%s在叫\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "小三",
		},
		foot: 4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
