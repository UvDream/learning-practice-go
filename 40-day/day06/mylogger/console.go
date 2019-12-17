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

// 日志信息
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	time := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s][%s][文件名:%s方法名:%s行号:%d]%s\n", time, getLogString(lv), fileName, funcName, lineNo, msg)
}

// Debug 日志
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	if c.enable(DEBUG) {
		log(DEBUG, format, a...)
	}

}

// Trace 日志
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	if c.enable(TRACE) {
		log(TRACE, format, a...)
	}

}

// Info 日志
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	if c.enable(INFO) {
		log(INFO, format, a...)
	}
}

// Warning 日志
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	if c.enable(WARNING) {
		log(WARNING, format, a...)

	}
}

// Error 日志
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	if c.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

// Fatal 日志
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	if c.enable(FATAL) {
		log(FATAL, format, a...)
	}
}
