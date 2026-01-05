package eslog

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestPrintAwareHandler_WithAttrs(t *testing.T) {
	var buf bytes.Buffer
	innerHandler := slog.NewTextHandler(&buf, nil)
	handler := &printAwareHandler{h: innerHandler, w: &buf}

	attrs := []slog.Attr{
		slog.String("key", "value"),
		slog.Int("count", 42),
	}

	newHandler := handler.WithAttrs(attrs)
	if newHandler == nil {
		t.Errorf("WithAttrs returned nil")
	}

	// Verify it returns a new handler
	if newHandler == handler {
		t.Errorf("WithAttrs should return a new handler instance")
	}

	// Verify it's the correct type
	if _, ok := newHandler.(*printAwareHandler); !ok {
		t.Errorf("WithAttrs should return a *printAwareHandler")
	}
}

func TestPrintAwareHandler_WithGroup(t *testing.T) {
	var buf bytes.Buffer
	innerHandler := slog.NewTextHandler(&buf, nil)
	handler := &printAwareHandler{h: innerHandler, w: &buf}

	newHandler := handler.WithGroup("testgroup")
	if newHandler == nil {
		t.Errorf("WithGroup returned nil")
	}

	// Verify it returns a new handler
	if newHandler == handler {
		t.Errorf("WithGroup should return a new handler instance")
	}

	// Verify it's the correct type
	if _, ok := newHandler.(*printAwareHandler); !ok {
		t.Errorf("WithGroup should return a *printAwareHandler")
	}
}
