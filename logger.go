package eslog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// LevelFatal the constant to represent the fatal log level
const (
	LevelFatal = slog.Level(12)
	// LevelPrint is a sentinel level used by Print/Printf/Println to
	// indicate that the level attribute should be omitted from output.
	// It is higher than default thresholds so it won't be filtered.
	LevelPrint = slog.Level(16)
)

// levelNames maps the LevelFatal to string "FATAL"
var levelNames = map[slog.Leveler]string{
	LevelFatal: "FATAL",
}

// eSlogLogger is used to extend slog.
type eSlogLogger struct {
	*slog.Logger
}

// Logger is the default logger which extends slog.
var Logger *eSlogLogger

// logLevel holds the log level of the logger.
var logLevel = &slog.LevelVar{}

func init() {
	initLogger(os.Stdout, false)
}

// printAwareHandler wraps a slog.Handler and prints only the Record.Message
// when the level equals LevelPrint, omitting standard key-value formatting.
type printAwareHandler struct {
	h slog.Handler
	w io.Writer
}

func (p *printAwareHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// Delegate to underlying handler; LevelPrint is used via direct Handle calls.
	return p.h.Enabled(ctx, level)
}

func (p *printAwareHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level == LevelPrint {
		_, err := io.WriteString(p.w, r.Message)
		return err
	}
	return p.h.Handle(ctx, r)
}

func (p *printAwareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &printAwareHandler{h: p.h.WithAttrs(attrs), w: p.w}
}

func (p *printAwareHandler) WithGroup(name string) slog.Handler {
	return &printAwareHandler{h: p.h.WithGroup(name), w: p.w}
}

// initLogger initializes the Logger and enables LevelFatal. Also it sets the default log
// level to LevelDebug
func initLogger(writer io.Writer, overwrite bool) {
	if Logger == nil || overwrite {
		if logLevel == nil {
			logLevel.Set(slog.LevelDebug)
		}

		opts := &slog.HandlerOptions{
			Level: logLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					level := a.Value.Any().(slog.Level)
					// Drop level attribute for sentinel LevelPrint.
					if level == LevelPrint {
						return slog.Attr{}
					}
					levelLabel, exists := levelNames[level]
					if !exists {
						levelLabel = level.String()
					}
					a.Value = slog.StringValue(levelLabel)
				}
				return a
			},
		}

		Logger = &eSlogLogger{
			Logger: slog.New(&printAwareHandler{h: slog.NewTextHandler(writer, opts), w: writer}),
		}
	}
}

// SetOutput can be used to overwrite the default output writer os.Stdout. Can be used for
// testing purposes or to swich logging to os.Stderr. In fact that function reinitializes
// the Logger.
func (l *eSlogLogger) SetOutput(w io.Writer) {
	initLogger(w, true)
}

// SetLogLevel sets the LogLevel of the Logger
func (l *eSlogLogger) SetLogLevel(lvl string) error {
	return logLevel.UnmarshalText([]byte(lvl))
}

// Print logs a simple, unstructured message to the configured Logger.
// The message is produced by fmt.Sprint over args and sent via context.Background.
// No level, timestamp, or attributes are set on the record; the Handler may supply defaults.
// Prefer structured or leveled logging APIs for richer context or severity control.
func Print(args ...any) {
	Logger.Handler().Handle(context.Background(), slog.Record{
		Level:   LevelPrint,
		Message: fmt.Sprint(args...),
	})
}

// Printf formats and logs a message using the package's default Logger.
//
// It applies fmt.Sprintf to the provided format and arguments, then submits the
// resulting text to the Logger's Handler with a background context. This helper
// emits only the message and does not set a level or attach attributes; prefer
// structured, leveled logging APIs when additional context is required.
func Printf(format string, args ...any) {
	Logger.Handler().Handle(context.Background(), slog.Record{
		Level:   LevelPrint,
		Message: fmt.Sprintf(format, args...),
	})
}

// Println logs a line by formatting args with fmt.Sprintln (space‑separated, with a trailing newline)
// and sending the resulting message to the configured Logger’s Handler using context.Background().
// It is a convenience wrapper similar to fmt.Println and does not set an explicit log level or add attributes.
func Println(args ...any) {
	Logger.Handler().Handle(context.Background(), slog.Record{
		Level:   LevelPrint,
		Message: fmt.Sprintln(args...),
	})
}

// Debugf logs at [LevelDebug]. Multiple args are joined with "  ".
func Debug(args ...any) {
	Logger.Debug(strings.Join(convertAnyToString(args...), " "))
}

// Debugf logs at [LevelDebug]. The function uses fmt.Sprintf with given format and args
// and log it.
func Debugf(format string, args ...any) {
	Logger.Debugf(format, args...)
}

// Debugf logs at [LevelDebug]. The function uses fmt.Sprintf with given format and args
// and log it.
func (l eSlogLogger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

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

// LogIfError check the given error. If error is nil nothing is logged. If error is not
// nil the loggerFunc is used to log the args. The error is not automatically added to
// args except args are empty.
func LogIfError(err error, loggerFunc func(args ...any), args ...any) {
	if err != nil {
		if len(args) == 0 {
			loggerFunc(err)
		} else {
			loggerFunc(args...)
		}
	}
}

// LogIfErrorf checks the given error. If error is nil nothing is logged. If error is not
// nil the loggerFuncf is used with the given format to print the given args. The error is
// not automatically added to args except args are empty.
func LogIfErrorf(err error, loggerFuncf func(format string, args ...any), format string, args ...any) {
	if err != nil {
		if len(args) == 0 {
			loggerFuncf(format, err)
		} else {
			loggerFuncf(format, args...)
		}
	}
}

// convertAnyToString converts all args of type any to string and
// returns them as []string
func convertAnyToString(args ...any) (strArr []string) {
	strArr = []string{}

	for _, something := range args {
		strArr = append(strArr, fmt.Sprintf("%v", something))
	}

	return strArr
}
