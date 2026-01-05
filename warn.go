package eslog

import (
	"fmt"
	"strings"
)

// Warn logs at [LevelWarn]. Multiple args are joined with "  ".
func Warn(args ...any) {
	Logger.Warn(strings.Join(convertAnyToString(args...), " "))
}

// Fatalf logs at [LevelWarn]. The function uses fmt.Sprintf with given format and args
// and log it.
func Warnf(format string, args ...any) {
	Logger.Warnf(format, args...)
}

// Fatalf logs at [LevelWarn]. The function uses fmt.Sprintf with given format and args
// and log it.
func (l eSlogLogger) Warnf(format string, args ...any) {
	l.Warn(fmt.Sprintf(format, args...))
}

// WarnLn logs at [LevelWarn] and appends a newline.
func WarnLn(msg string, args ...any) {
	Logger.WarnLn(msg, args...)
}

func (l eSlogLogger) WarnLn(msg string, args ...any) {
	args = append(args, "\n")
	Logger.Warn(msg, args...)
}
