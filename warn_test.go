package eslog_test

import (
	"io"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestWarn(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test warn", []any{}, "test warn"},
		{"success with args", "test warn %s %s", []any{"arg1", "arg2"}, "test warn arg1 arg2"},
		{"success with args positions", "test warn %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test warn arg2 arg1 arg3"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Warn")
			assert.NoError(t, err)
			eslog.Warnf(tst.message, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}
func TestWarnf(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test warn", []any{}, "test warn"},
		{"success with args", "test warn %s %s", []any{"arg1", "arg2"}, "test warn arg1 arg2"},
		{"success with args positions", "test warn %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test warn arg2 arg1 arg3"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Warn")
			assert.NoError(t, err)
			eslog.Warnf(tst.message, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestWarn_Nonf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Warn")
	assert.NoError(t, err)

	eslog.Warn("simple", "warn")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "simple warn")
}

func TestWarn_Ln(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Warn")
	assert.NoError(t, err)

	eslog.WarnLn("warnln message")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "warnln message")
	assert.Contains(t, string(out), "\n")
}
