package eslog_test

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestSetOutputAndLogLevel(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Debug")
	assert.NoError(t, err)
	eslog.Logger.Infof("test output")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "test output")
}

func TestSetLogLevelFiltering(t *testing.T) {
	err := eslog.Logger.SetLogLevel("Info")
	assert.NoError(t, err)

	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.Debug("this should not appear")
	eslog.Info("this should appear")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.NotContains(t, string(out), "this should not appear")
	assert.Contains(t, string(out), "this should appear")
}

func TestCustomLogLevel(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.Logger.Log(context.Background(), eslog.LevelFatal, "fatal level log")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "FATAL")
	assert.Contains(t, string(out), "fatal level log")
}

func TestField(t *testing.T) {
	// Test Field creates a valid slog.Attr
	attr := eslog.Field("testkey", "testvalue")
	if attr.Key != "testkey" {
		t.Errorf("Expected key 'testkey', got %q", attr.Key)
	}
	if attr.Value.String() != "testvalue" {
		t.Errorf("Expected value 'testvalue', got %q", attr.Value.String())
	}

	// Test with different types
	intAttr := eslog.Field("number", 42)
	if intAttr.Value.Int64() != 42 {
		t.Errorf("Expected value 42, got %d", intAttr.Value.Int64())
	}
}

func TestParseText(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    slog.Level
		shouldError bool
	}{
		{"DEBUG uppercase", "DEBUG", slog.LevelDebug, false},
		{"debug lowercase", "debug", slog.LevelDebug, false},
		{"INFO uppercase", "INFO", slog.LevelInfo, false},
		{"info lowercase", "info", slog.LevelInfo, false},
		{"WARN uppercase", "WARN", slog.LevelWarn, false},
		{"warn lowercase", "warn", slog.LevelWarn, false},
		{"ERROR uppercase", "ERROR", slog.LevelError, false},
		{"error lowercase", "error", slog.LevelError, false},
		{"invalid level", "invalid", slog.Level(0), true},
		{"empty string", "", slog.Level(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := eslog.ParseText(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
