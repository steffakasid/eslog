package eslog

import (
	"fmt"
	"strings"
)

// Infof logs at [LevelInfo]. Multiple args are joined with "  ".
func Info(args ...any) {
	Logger.Info(strings.Join(convertAnyToString(args...), " "))
}

// Infof logs at [LevelInfo]. The function uses fmt.Sprintf with given format and args
// and log it.
func Infof(format string, args ...any) {
	Logger.Infof(format, args...)
}

// Infof logs at [LevelInfo]. The function uses fmt.Sprintf with given format and args
// and log it.
func (l eSlogLogger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

// InfoLn logs at [LevelInfo] and appends a newline.
func InfoLn(msg string, args ...any) {
	Logger.InfoLn(msg, args...)
}

func (l eSlogLogger) InfoLn(msg string, args ...any) {
	args = append(args, "\n")
	Logger.Info(msg, args...)
}
