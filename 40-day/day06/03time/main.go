package main

import (
	"fmt"
	"time"
)

// time
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())   //年
	fmt.Println(now.Month())  //月
	fmt.Println(now.Day())    //日
	fmt.Println(now.Hour())   //时
	fmt.Println(now.Minute()) //分
	fmt.Println(now.Second()) //秒
	// 时间戳
	fmt.Println(now.Unix())     //时间戳
	fmt.Println(now.UnixNano()) //纳秒
	// time.Unix
	ret := time.Unix(1576549339, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	// 时间间隔
	fmt.Println(time.Second)
	// 时间+24h
	fmt.Println(now.Add(24 * time.Hour))
	// 定时器
	// 一秒钟执行一次
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }
	// 时间格式化
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	// 将按照对应格式解析字符串类型的时间
	t, err := time.Parse("2006-01-02", "2019-08-03")
	if err != nil {
		fmt.Printf("错误%v\n", err)
		return
	}
	fmt.Println(t)
}
