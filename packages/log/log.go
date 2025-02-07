package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	var config zap.Config
	var zapLog *zap.Logger

	config = zap.NewProductionConfig()

	config.DisableStacktrace = true

	if gin.IsDebugging() {
		zapLog, _ = config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		zapLog, _ = config.Build(zap.AddCallerSkip(1))
	}

	log = logger{
		zapLog: zapLog,
	}
}

type ILogger interface {
	Sync()
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type logger struct {
	zapLog *zap.Logger
}

var log logger

func Logger() ILogger {
	return &log
}

func ZapLogger() *zap.Logger {
	return log.zapLog
}

func (l *logger) Sync() {
	_ = l.zapLog.Sync()
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLog.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zapLog.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zapLog.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.zapLog.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLog.Fatal(msg, fields...)
}
