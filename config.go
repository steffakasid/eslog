package eslog

import (
	"fmt"
	"io"
	"log/slog"
)

// Format represents the log format type.
type Format int

const (
	TextFormat Format = iota
	JSONFormat
)

// String returns the string representation of Format.
func (f Format) String() string {
	switch f {
	case JSONFormat:
		return "json"
	case TextFormat:
		return "text"
	default:
		return "unknown"
	}
}

// ParseFormat converts a string to Format.
// It returns an error if the format string is invalid.
func ParseFormat(format string) (Format, error) {
	switch format {
	case "json":
		return JSONFormat, nil
	case "text":
		return TextFormat, nil
	default:
		return TextFormat, fmt.Errorf("invalid format: %s", format)
	}
}

type Config struct {
	Level  slog.Level // Log level: debug, info, warn, error, fatal
	Format Format // Log format: TextFormat or JSONFormat
    out io.Writer
}