package a

import log "github.com/feixiao/log"

var logger  *log.Log
var tag string

func init() {
	// 模块 默认使用log4go输出info基本日志
	tag = "test"
	logger = log.NewLog4go(log.INFO)
	logger.SetTag(tag)
}


// SetLogger 外部控制模块日志输出
func SetLogger(log *log.Log) {
	logger = log
	logger.SetTag(tag)
}


// for test log

func PrintLog() {
	logger.Info("come for test module")
}