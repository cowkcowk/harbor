package log

import (
	"fmt"
	"time"
)

var logger = New()

type Logger struct {
	lvl Level
}

func New() *Logger {
	return &Logger{}
}

func DefaultLogger() *Logger {
	return logger
}

// Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.lvl <= DebugLevel {
		record := NewRecord(time.Now(), fmt.Sprintf(format, v...), "", DebugLevel)
		_ = l.output(record)
	}
}