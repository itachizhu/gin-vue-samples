package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

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
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), syncWriter), zap.NewAtomicLevelAt(zapcore.DebugLevel))
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
