package log

import "fmt"

type Level int

const (
	DEBUG = iota + 1
	TRACE
	INFO
	WARNING
	ERROR
)

// 默认日志输出接口
var logger Log

func init() {
	logger.log = NewLog4go(DEBUG)  //  默认的输出最低级别,模块日志输出级别才有效
}

// SetLogger 设置日志的具体输出
// level 日志输出级别
// log 日志的具体输出实现
func SetLogger(level Level, log Logger) {
	logger.log = log
	logger.level = level
}

// SetLoggerLevel 只修改默认的输出日志级别
func SetLoggerLevel(level Level) {
	logger.level = level
}


func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	logger.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

/**************************************** 下面代码为了控制模块日志输出 ***************************************************/

// 模块使用
type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type Log struct {
	level Level 	// 设置日志输出级别
	log   Logger	// 实际的日志输出
	tag   string    // 日志输出添加tag
}

// NewLog 实例化Log对象
func NewLog(log Logger, level Level) *Log {
	return &Log{
		log:   log,
		level: level,
	}
}

// NewLog4go 基于log4go的实现
func NewLog4go(level Level) *Log {
	return &Log{
		log:   &Log4go{},
		level: level,
	}
}


// NewKratosLog
func NewKratosLog(level Level) *Log {
	return &Log{
		log:   &KratosLog{},
		level: level,
	}
}


func (l *Log) SetTag(tag string) {
	l.tag = tag
}

func (l *Log) appendTag(format string) string {
	if l.tag == "" {
		return format
	}

	return fmt.Sprintf("[%s] %s", l.tag, format)
}


// Trace 输出TRACE级别日志，前提是level不高于TRACE才输出
func (l *Log) Trace(format string, args ...interface{}) {
	if l.level > TRACE {
		return
	}
	format = l.appendTag(format)
	l.log.Trace(format, args...)
}

// Trace 输出TRACE级别日志，前提是level不高于TRACE才输出
func (l *Log) Debug(format string, args ...interface{}) {
	if l.level > DEBUG {
		return
	}
	format = l.appendTag(format)
	l.log.Debug(format, args...)
}

// Info 输出INFO级别日志，前提是level不高于INFO才输出
func (l *Log) Info(format string, args ...interface{}) {
	if l.level > INFO {
		return
	}
	format = l.appendTag(format)
	l.log.Info(format, args...)
}

// Warn 输出WARN级别日志，前提是level不高于WARNING才输出
func (l *Log) Warn(format string, args ...interface{}) {
	if l.level > WARNING {
		return
	}
	format = l.appendTag(format)
	l.log.Warn(format, args...)
}

// Error 输出ERROR级别日志，前提是level不高于ERROR才输出
func (l *Log) Error(format string, args ...interface{}) {
	if l.level > ERROR {
		return
	}
	format = l.appendTag(format)
	l.log.Error(format, args...)
}
