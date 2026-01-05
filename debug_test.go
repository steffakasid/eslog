package eslog_test

import (
	"io"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestDebugf(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test debug", []any{}, "test debug"},
		{"success with args", "test debug %s %s", []any{"arg1", "arg2"}, "test debug arg1 arg2"},
		{"success with args positions", "test debug %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test debug arg2 arg1 arg3"},
		{"success with error", "test debug %s", []any{os.ErrNotExist}, "test debug file does not exist"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Debug")
			assert.NoError(t, err)
			eslog.Debugf(tst.message, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestDebug_Nonf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Debug")
	assert.NoError(t, err)

	eslog.Debug("simple debug message")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "simple debug message")
}

func TestDebug_Ln(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Debug")
	assert.NoError(t, err)

	eslog.DebugLn("debugln message")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "debugln message")
	assert.Contains(t, string(out), "\n")
}

func TestDebugNonfAndLn(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Debug")
	assert.NoError(t, err)

	// non-format debug
	eslog.Debug("simple debug message")

	// DebugLn should append newline
	eslog.DebugLn("debugln message")

	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "simple debug message")
	assert.Contains(t, string(out), "debugln message")
}
