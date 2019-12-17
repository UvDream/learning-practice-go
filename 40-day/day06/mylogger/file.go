package mylogger

import (
	"fmt"
	"time"
)

// FileLogger 往文件里面写日志相关代码
type FileLogger struct {
	Level       LogLevel
	filePath    string //路径
	fileName    string //名字
	maxFileSize int64
}

// NewFileLogger NewFileLogger
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// Debug 日志
func (f *FileLogger) Debug(msg string) {
	if f.enable(DEBUG) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][DEBUG]%s\n", time, msg)
	}

}

// Trace 日志
func (f *FileLogger) Trace(msg string) {
	if f.enable(TRACE) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][DEBUG]%s\n", time, msg)
	}

}

// Info 日志
func (f *FileLogger) Info(msg string) {
	if f.enable(INFO) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][INFO]%s\n", time, msg)
	}
}

// Warning 日志
func (f *FileLogger) Warning(msg string) {
	if f.enable(WARNING) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}

// Error 日志
func (f *FileLogger) Error(msg string) {
	if f.enable(ERROR) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}

// Fatal 日志
func (f *FileLogger) Fatal(msg string) {
	if f.enable(FATAL) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]%s\n", time, msg)
	}
}
