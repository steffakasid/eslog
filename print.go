package eslog

import (
	"context"
	"fmt"
	"log/slog"
)

// LevelPrint is a sentinel level used by Print/Printf/Println to
// indicate that the level attribute should be omitted from output.
// It is higher than default thresholds so it won't be filtered.
const LevelPrint = slog.Level(16)

// Print logs a simple, unstructured message to the configured Logger.
// The message is produced by fmt.Sprint over args and sent via context.Background.
// No level, timestamp, or attributes are set on the record; the Handler may supply defaults.
// Prefer structured or leveled logging APIs for richer context or severity control.
func Print(args ...any) {
	Logger.Print(args...)
}

// Printf formats and logs a message using the package's default Logger.
//
// It applies fmt.Sprintf to the provided format and arguments, then submits the
// resulting text to the Logger's Handler with a background context. This helper
// emits only the message and does not set a level or attach attributes; prefer
// structured, leveled logging APIs when additional context is required.
func Printf(format string, args ...any) {
	Logger.Printf(format, args...)
}

// Println logs a line by formatting args with fmt.Sprintln (space‑separated, with a trailing newline)
// and sending the resulting message to the configured Logger’s Handler using context.Background().
// It is a convenience wrapper similar to fmt.Println and does not set an explicit log level or add attributes.
func Println(args ...any) {
	Logger.Println(args...)
}

func (l eSlogLogger) Print(args ...any) {
	err := l.Handler().Handle(context.Background(), slog.Record{
		Level:   LevelPrint,
		Message: fmt.Sprint(args...),
	})
	if err != nil {
		panic(err)
	}
}

func (l eSlogLogger) Printf(format string, args ...any) {
	l.Print(fmt.Sprintf(format, args...))
}

func (l eSlogLogger) Println(args ...any) {
	args = append(args, "\n")
	l.Print(args...)
}
