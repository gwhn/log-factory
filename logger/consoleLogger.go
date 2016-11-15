package logger

import (
	"os"
	"fmt"
)

func init() {
	New = newConsoleLogger
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