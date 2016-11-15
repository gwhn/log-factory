package main

import (
	"errors"
	"os"
	"fmt"
)

type (
	Logger interface {
		Debug(format string, values ...interface{})
		Warning(format string, values ...interface{})
		Info(format string, values ...interface{})
		Error(format string, values ...interface{})
		Critical(format string, values ...interface{})
	}

	LoggerFactory interface {
		Make() Logger
	}

	LoggerFactoryFunc func() Logger
)

func (fn LoggerFactoryFunc) Make() Logger {
	return fn()
}

func main() {
	log := LoggerFactoryFunc(newConsoleLogger).Make()
	log.Debug("%d + %d = %d", 2, 3, 5)
	log.Warning("watch out!")
	log.Info("i'm here")
	log.Error("something's gone wrong with: %v", errors.New("ME"))
	log.Critical("flat line, no response with: %v", errors.New("YOU"))
}

type consoleLogger struct { file *os.File }

func newConsoleLogger() Logger { return &consoleLogger {file: os.Stdout} }

func (l *consoleLogger) Debug(format string, values ...interface{}) {
	fmt.Fprintf(l.file, "DEBUG: " + format + "\n", values...)
}

func (l *consoleLogger) Warning(format string, values ...interface{}) {
	fmt.Fprintf(l.file, "WARNING: " + format + "\n", values...)
}

func (l *consoleLogger) Info(format string, values ...interface{}) {
	fmt.Fprintf(l.file, "INFO: " + format + "\n", values...)
}

func (l *consoleLogger) Error(format string, values ...interface{}) {
	fmt.Fprintf(l.file, "ERROR: " + format + "\n", values...)
}

func (l *consoleLogger) Critical(format string, values ...interface{}) {
	fmt.Fprintf(l.file, "CRITICAL: " + format + "\n", values...)
}