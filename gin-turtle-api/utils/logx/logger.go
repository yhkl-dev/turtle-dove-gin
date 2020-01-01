package logx

import (
	"fmt"
	"os"
	"time"
)

const (
	LevelError = iota
	LevelWarning
	LevelInfomational
	LevelDebug
)

var logger *Logger

// Logger log struct
type Logger struct {
	level int
}

// Println print log
func (l *Logger) Println(msg string) {
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
}

// Panic extreme error
func (l *Logger) Panic(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	message := fmt.Sprintf("[Panic] "+format, v...)
	l.Println(message)
	os.Exit(0)
}

// Error errors
func (l *Logger) Error(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	message := fmt.Sprintf("[Error] "+format, v...)
	l.Println(message)
}

// Warning warnings for log
func (l *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > l.level {
		return
	}
	message := fmt.Sprintf("[Warning] "+format, v...)
	l.Println(message)
}

// Info infos
func (l *Logger) Info(format string, v ...interface{}) {

	if LevelInfomational > l.level {
		return
	}
	message := fmt.Sprintf("[Info] "+format, v...)
	l.Println(message)
}

// Debug check
func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}
	message := fmt.Sprintf("[Debug] "+format, v...)
	l.Println(message)
}

// BuildLogger construct logger
func BuildLogger(level string) {
	intLevel := LevelError

	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInfomational
	case "debug":
		intLevel = LevelDebug
	}

	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// Log return log object
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
