package util

import (
	"github.com/sirupsen/logrus"
)

//LogLevel represents possible log levels
type LogLevel uint

const (
	//ErrorLogLevel will only log error messages
	ErrorLogLevel LogLevel = iota + 1

	//WarningLogLevel will only log warnings
	WarningLogLevel

	//InfoLogLevel will only log infos
	InfoLogLevel

	//DebugLogLevel will only log debugs
	DebugLogLevel
)

//Logger is the log implementation that wraps underlying logging mechanism
type Logger struct {
	pkg       string
	component string
	log       *logrus.Entry
}

//NewLogger creates a log instance for the given component with the given log level enabled
func NewLogger(pkg string, component string, level LogLevel) *Logger {
	ll := xlateLevel(level)
	log := logrus.New()
	log.SetLevel(ll)
	log.SetFormatter(&logrus.TextFormatter{ForceColors: true, PadLevelText: true})

	e := log.WithFields(logrus.Fields{"component": component, "package": pkg})
	return &Logger{pkg: pkg, component: component, log: e}
}

//Info message logging
func (l *Logger) Info(msg string, args ...interface{}) {
	l._logIt(l.log.Infof, msg, args...)
}

func (l *Logger) _logIt(fn func(msg string, args ...interface{}), msg string, args ...interface{}) {
	fn(msg, args...)
}

func xlateLevel(level LogLevel) logrus.Level {
	switch level {
	case ErrorLogLevel:
		{
			return logrus.ErrorLevel
		}
	case WarningLogLevel:
		{
			return logrus.WarnLevel
		}
	case InfoLogLevel:
		{
			return logrus.InfoLevel
		}
	case DebugLogLevel:
		{
			return logrus.DebugLevel
		}
	default:
		return logrus.InfoLevel
	}

}