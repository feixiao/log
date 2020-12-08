package log_test

import (
	"github.com/feixiao/log"
	"github.com/feixiao/log4go"
	"testing"
)

// type Option func(cfg *Config)

// func WithLogger(l log.Logger) Option {
// 	return func(cfg *Config) {
// 			cfg.l = l
// 	}
// }

//  ModuleA : 模拟模块A
type ModuleA struct {
	l log.Logger
}

func NewModuleA() *ModuleA {
	var m ModuleA
	m.l = log.NewLog4go(log.DEBUG)
	return &m
}

func (m *ModuleA) DoSomething() {
	m.l.Trace("Trace")
	m.l.Debug("Debug")
	m.l.Info("Info")
	m.l.Warn("Warn")
	m.l.Error("Error")
}

func (m *ModuleA) SetLogger(l log.Logger) {
	m.l = l
}

func TestModuleA(t *testing.T) {
	defer log4go.Close()

	// 默认log4go使用全局的，所以这边只设置了级别
	// 通常一个应该程序，共享一个logger， 然后将该logger传递给模块 NewLog(log Logger, level Level)
	l := log.NewLog4go(log.INFO)

	m := NewModuleA()
	m.SetLogger(l)
	m.DoSomething()
}
