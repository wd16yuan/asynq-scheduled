package internal

import (
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Logger struct {
	base *zap.Logger
}

func NewLogger(l *zap.Logger) asynq.Logger {
	return &Logger{base: l}
}

func (l *Logger) Debug(args ...interface{}) {
	msg, fields := l.convert(args...)
	l.base.Debug(msg, fields...)
}

func (l *Logger) Info(args ...interface{}) {
	msg, fields := l.convert(args...)
	l.base.Info(msg, fields...)
}

func (l *Logger) Warn(args ...interface{}) {
	msg, fields := l.convert(args...)
	l.base.Warn(msg, fields...)
}

func (l *Logger) Error(args ...interface{}) {
	msg, fields := l.convert(args...)
	l.base.Error(msg, fields...)
}

func (l *Logger) Fatal(args ...interface{}) {
	msg, fields := l.convert(args...)
	l.base.Fatal(msg, fields...)
}

func (l *Logger) convert(args ...interface{}) (string, []zap.Field) {
	messages := make([]interface{}, 0)
	fields := make([]zap.Field, 0)
	for _, a := range args {
		switch v := a.(type) {
		case zap.Field:
			fields = append(fields, v)
		default:
			messages = append(messages, v)
		}
	}
	return fmt.Sprint(messages...), fields
}
