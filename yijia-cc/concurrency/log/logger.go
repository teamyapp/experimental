package log

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type LogLevel int

const (
	Fatal LogLevel = iota
	Error
	Warning
	Info
	Debug
)

var logLevelNames = map[LogLevel]string{
	Fatal : "Fatal",
	Error : "Error",
	Warning : "Warning",
	Info: "Info",
	Debug: "Debug",
}

type Logger struct {
	logLevel LogLevel
}

func (l *Logger) Log(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel > l.logLevel {
		return
	}

	message := fmt.Sprintf(format, args...)
	formattedTime := time.Now().Format(timeFormat)
	fmt.Printf("[%s] [%s] %s\n", logLevelNames[logLevel], formattedTime, message)
}

func NewLogger(logLevel LogLevel) Logger {
	return Logger{
		logLevel: logLevel,
	}
}
