package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //登记一个goroutine
	go hello()
	fmt.Println("Hello, world!")

	wg.Wait()
	//time.Sleep(time.Second)

}
func hello() {
	fmt.Println("你好啊!")
	wg.Done()
}
