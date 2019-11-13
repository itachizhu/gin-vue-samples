package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

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

type zapLogger struct {
	zLogger *zap.SugaredLogger
}

func NewZap(logName string) Logger {
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  logName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), syncWriter), zap.NewAtomicLevelAt(zapcore.DebugLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &zapLogger{zLogger: logger.Sugar()}
}

func (m *zapLogger) Debug(v ...interface{}) {
	m.zLogger.Debug(v...)
}

func (m *zapLogger) Debugf(format string, v ...interface{}) {
	m.zLogger.Debugf(format, v...)
}

func (m *zapLogger) Info(v ...interface{}) {
	m.zLogger.Info(v...)
}

func (m *zapLogger) Infof(format string, v ...interface{}) {
	m.zLogger.Infof(format, v...)
}

func (m *zapLogger) Warn(v ...interface{}) {
	m.zLogger.Warn(v...)
}

func (m *zapLogger) Warnf(format string, v ...interface{}) {
	m.zLogger.Warnf(format, v...)
}

func (m *zapLogger) Error(v ...interface{}) {
	m.zLogger.Error(v...)
}

func (m *zapLogger) Errorf(format string, v ...interface{}) {
	m.zLogger.Errorf(format, v...)
}
