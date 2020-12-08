package log

import (
	"github.com/go-kratos/kratos/pkg/log"
)

type KratosLog struct {
}

func (l *KratosLog) Debug(format string, args ...interface{}) {
	log.Info(format, args...)
}

func (l *KratosLog) Trace(format string, args ...interface{}) {
	log.Info(format, args...)
}

func (l *KratosLog) Info(format string, args ...interface{}) {
	// log4go.Info(format, args)
	log.Info(format, args...)
}

func (l *KratosLog) Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func (l *KratosLog) Error(format string, args ...interface{}) {
	log.Error(format, args...)
}
