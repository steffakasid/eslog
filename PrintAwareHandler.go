package eslog

import (
	"context"
	"io"
	"log/slog"
)

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