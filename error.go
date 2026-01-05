package eslog

import (
	"fmt"
	"strings"
)

// Error logs at [LevelError]. Multiple args are joined with "  ".
func Error(args ...any) {
	Logger.Error(strings.Join(convertAnyToString(args...), " "))
}

// Errorf logs at [LevelError]. The function uses fmt.Sprintf with given format and args
// and log it.
func Errorf(format string, args ...any) {
	Logger.Errorf(format, args...)
}

// Errorf logs at [LevelError]. The function uses fmt.Sprintf with given format and args
// and log it.
func (l eSlogLogger) Errorf(format string, args ...any) {
	Logger.Error(fmt.Sprintf(format, args...))
}

// ErrorLn logs at [LevelError] and appends a newline.
func ErrorLn(msg string, args ...any) {
	Logger.ErrorLn(msg, args...)
}

func (l eSlogLogger) ErrorLn(msg string, args ...any) {
	args = append(args, "\n")
	Logger.Error(msg, args...)
}
