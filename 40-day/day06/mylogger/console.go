package mylogger

import (
	"fmt"
	"time"
)

// NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}
func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	time := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s][%s][文件名:%s方法名:%s行号:%d]%s\n", time, getLogString(lv), fileName, funcName, lineNo, msg)
}

// Debug 日志
func (c ConsoleLogger) Debug(msg string) {
	if c.enable(DEBUG) {
		log(DEBUG, msg)
	}

}

// Trace 日志
func (c ConsoleLogger) Trace(msg string) {
	if c.enable(TRACE) {
		log(TRACE, msg)
	}

}

// Info 日志
func (c ConsoleLogger) Info(msg string) {
	if c.enable(INFO) {
		log(INFO, msg)
	}
}

// Warning 日志
func (c ConsoleLogger) Warning(msg string) {
	if c.enable(WARNING) {
		log(WARNING, msg)

	}
}

// Error 日志
func (c ConsoleLogger) Error(msg string) {
	if c.enable(ERROR) {
		log(ERROR, msg)
	}
}

// Fatal 日志
func (c ConsoleLogger) Fatal(msg string) {
	if c.enable(FATAL) {
		log(FATAL, msg)
	}
}
