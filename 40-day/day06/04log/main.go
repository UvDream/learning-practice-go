package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./40-day/day06/04log/xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开错误,err:%v\n", err)
		return
	}
	log.SetOutput(file)
	for {
		log.Println("这是一条日志")
		time.Sleep(time.Second * 3)
	}
}
