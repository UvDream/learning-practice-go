package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// LogLevel Log类型
type LogLevel uint16

const (
	// UNKNOWN UNKNOWN
	UNKNOWN LogLevel = iota
	// DEBUG debug
	DEBUG
	// TRACE TRACE
	TRACE
	// INFO INFO
	INFO
	// WARNING WARNING
	WARNING
	// ERROR ERROR
	ERROR
	// FATAL FATAL
	FATAL
)

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}
func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

// Debug 日志
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][DEBUG]%s\n", time, msg)
	}

}

// Trace 日志
func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][DEBUG]%s\n", time, msg)
	}

}

// Info 日志
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][INFO]%s\n", time, msg)
	}
}

// Warning 日志
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}

// Error 日志
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}

// Fatal 日志
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}
