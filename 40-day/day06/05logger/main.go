package main

import (
	"hello/40-day/day06/mylogger"
	"time"
)

// 测试日志库
func main() {
	log := mylogger.NewLog("DEBUG")
	for {
		log.Debug("1,这是一条Debugger日志")
		log.Trace("2.TRACE")
		log.Info("3.这是一条Info日志")
		log.Warning("4.Waring")
		id := 1001
		name := "张三"
		log.Error("5.错误日志,id:%d,name%s", id, name)
		log.Fatal("6.fatal")
		time.Sleep(time.Second)
	}
}
