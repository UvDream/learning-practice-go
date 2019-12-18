package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 往文件里面写日志相关代码
type FileLogger struct {
	Level       LogLevel
	filePath    string //路径
	fileName    string //名字
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 //大小
}

// NewFileLogger 文件日志构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 初始化文件以及打开
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// Closed 关闭
func (f *FileLogger) Closed() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 判断切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取失败,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

// 切割日志
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //拿到当前日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1.关闭当前文件
	file.Close()
	// 2.备份一下
	os.Rename(logName, newLogName)
	// 3.打开新的文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件错误,err:%v\n", err)
		return nil, err
	}
	// 4.将打开的新日志对象赋值给f.fileObj
	// f.fileObj = fileOb
	return fileObj, nil
}

// 日志信息
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		time := now.Format("2006-01-02 15:04:05")
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) //日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s][%s][文件名:%s方法名:%s行号:%d]%s\n", time, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.fileObj) {
				newFile, err := f.splitFile(f.errFileObj) //日志文件
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			// 如果记录的日志大于等于ERROR级别,还需要在error中记录一份
			fmt.Fprintf(f.errFileObj, "[%s][%s][文件名:%s方法名:%s行号:%d]%s\n", time, getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

// Debug 日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace 日志
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info 日志
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning 日志
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error 日志
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal 日志
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
