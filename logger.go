package eslog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

// eSlogLogger is used to extend slog.
type eSlogLogger struct {
	*slog.Logger
	config *Config
}

// Logger is the default logger which extends slog.
var Logger *eSlogLogger

// logLevel holds the log level of the logger.
var logLevel = &slog.LevelVar{}

func New(cfg *Config) *eSlogLogger {
	return initLogger(cfg)
}

func init() {

	cfg := &Config{
		Level:  slog.LevelDebug,
		Format: TextFormat,
		out:    os.Stdout,
	}

	// Initialize the default Logger on package load.
	Logger = New(cfg)
}

// initLogger initializes the Logger and enables LevelFatal. Also it sets the default log
// level to LevelDebug
func initLogger(cfg *Config) *eSlogLogger {

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

	return &eSlogLogger{
		Logger: slog.New(&printAwareHandler{h: slog.NewTextHandler(cfg.out, opts), w: cfg.out}),
		config: cfg,
	}
}

func Field(key string, value any) slog.Attr {
	return slog.Any(key, value)
}

// SetOutput can be used to overwrite the default output writer os.Stdout. Can be used for
// testing purposes or to swich logging to os.Stderr. In fact that function reinitializes
// the Logger.
func (l *eSlogLogger) SetOutput(w io.Writer) {
	cfg := Config{
		out: w,
	}
	Logger = initLogger(&cfg)
}

// SetLogLevel sets the LogLevel of the Logger
func (l *eSlogLogger) SetLogLevel(lvl string) error {
	return logLevel.UnmarshalText([]byte(lvl))
}

func ParseText(text string) (slog.Level, error) {
	var level slog.Level
	err := level.UnmarshalText([]byte(text))
	return level, err
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
