package logger

import (
	"io"
	"log"
	"os"
)

type defaultLogger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	mError  *log.Logger
}

func New(logName string) Logger {
	file, err := os.OpenFile(logName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	return &defaultLogger{
		log.New(io.MultiWriter(file, os.Stdout), "DEBUG: ", log.Ldate|log.Lmicroseconds|log.Lshortfile),
		log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Lmicroseconds|log.Lshortfile),
		log.New(io.MultiWriter(file, os.Stdout), "WARN: ", log.Ldate|log.Lmicroseconds|log.Lshortfile),
		log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Lmicroseconds|log.Lshortfile),
	}
}

func (m *defaultLogger) Debug(v ...interface{}) {
	m.debug.Println(v...)
}

func (m *defaultLogger) Debugf(format string, v ...interface{}) {
	m.debug.Printf(format, v...)
}

func (m *defaultLogger) Info(v ...interface{}) {
	m.info.Println(v...)
}

func (m *defaultLogger) Infof(format string, v ...interface{}) {
	m.info.Printf(format, v...)
}

func (m *defaultLogger) Warn(v ...interface{}) {
	m.warning.Println(v...)
}

func (m *defaultLogger) Warnf(format string, v ...interface{}) {
	m.warning.Printf(format, v...)
}

func (m *defaultLogger) Error(v ...interface{}) {
	m.mError.Println(v...)
}

func (m *defaultLogger) Errorf(format string, v ...interface{}) {
	m.mError.Printf(format, v...)
}
