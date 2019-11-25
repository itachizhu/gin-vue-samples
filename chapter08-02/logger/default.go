package logger

import "sync"

var logg Logger
var once sync.Once

func init() {
	once.Do(func() {
		// 这个名称在后面可以使用配置来设置
		logg = newZap("localhost.log")
	})
}

func Debug(v ...interface{}) {
	logg.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logg.Debugf(format, v...)
}

func Info(v ...interface{}) {
	logg.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logg.Infof(format, v...)
}

func Warn(v ...interface{}) {
	logg.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	logg.Warnf(format, v...)
}

func Error(v ...interface{}) {
	logg.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logg.Errorf(format, v...)
}
