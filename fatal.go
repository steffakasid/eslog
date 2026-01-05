package eslog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const LevelFatal = slog.Level(12)

// levelNames maps the LevelFatal to string "FATAL"
var levelNames = map[slog.Leveler]string{
	LevelFatal: "FATAL",
}

// Fatal logs at [LevelFatal].  Multiple args are joined with " ".
func Fatal(args ...any) {
	Logger.Fatal(strings.Join(convertAnyToString(args...), " "))
}

// Fatalf logs at [LevelFatal]. The function uses fmt.Sprintf with given format and args
// and log it. Also it calls os.Exit(1).
func Fatalf(format string, args ...any) {
	Logger.Fatalf(format, args...)
}

// Fatal logs at [LevelFatal]. Also it calls os.Exit(1).
func (l eSlogLogger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}

// Fatalf logs at [LevelFatal]. The function uses fmt.Sprintf with given format and args
// and log it. Also it calls os.Exit(1).
func (l eSlogLogger) Fatalf(format string, args ...any) {
	l.Fatal(fmt.Sprintf(format, args...))
}

// FatalLn logs at [LevelFatal] and appends a newline, then exits.
func FatalLn(msg string, args ...any) {
	Logger.FatalLn(msg, args...)
}

func (l eSlogLogger) FatalLn(msg string, args ...any) {
	args = append(args, "\n")
	Logger.Fatal(msg, args...)
}
