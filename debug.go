package eslog

import "fmt"

// Debugf logs at [LevelDebug]. Multiple args are joined with "  ".
func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}

// Debugf logs at [LevelDebug]. The function uses fmt.Sprintf with given format and args
// and log it.
func Debugf(format string, args ...any) {
	Logger.Debugf(format, args...)
}

func DebugLn(msg string, args ...any) {
	Logger.DebugLn(msg, args...)
}

// Debugf logs at [LevelDebug]. The function uses fmt.Sprintf with given format and args
// and log it.
func (l eSlogLogger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l eSlogLogger) DebugLn(msg string, args ...any) {
	args = append(args, "\n")
	Logger.Debug(msg, args...)
}
