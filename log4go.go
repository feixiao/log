package log

import log4go "github.com/feixiao/log4go"


type Log4go struct {

}


func (l *Log4go) Debug(format string, args ...interface{}) {
	log4go.Debug(format, args...)
}

func (l *Log4go) Trace(format string, args ...interface{}) {
	log4go.Trace(format, args...)
}


func (l *Log4go) Info(format string, args ...interface{}) {
	// log4go.Info(format, args)
	log4go.Info(format, args...)
}


func (l *Log4go) Warn(format string, args ...interface{}) {
	log4go.Warn(format, args...)
}

func (l *Log4go) Error(format string, args ...interface{}) {
	log4go.Error(format, args...)
}

// Error(format interface{}, args ...interface{})