package eslog_test

import (
	"io"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestInfo(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test info", []any{}, "test info"},
		{"success with args", "test info %s %s", []any{"arg1", "arg2"}, "test info arg1 arg2"},
		{"success with args positions", "test info %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test info arg2 arg1 arg3"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Info")
			assert.NoError(t, err)

			eslog.Infof(tst.message, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}
func TestInfof(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test info", []any{}, "test info"},
		{"success with args", "test info %s %s", []any{"arg1", "arg2"}, "test info arg1 arg2"},
		{"success with args positions", "test info %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test info arg2 arg1 arg3"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Info")
			assert.NoError(t, err)

			eslog.Infof(tst.message, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestInfo_Nonf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Info")
	assert.NoError(t, err)

	eslog.Info("simple", "info")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "simple info")
}

func TestInfo_Ln(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Info")
	assert.NoError(t, err)

	eslog.InfoLn("infoln message")
	err = w.Close()
	assert.NoError(t, err)
	out, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Contains(t, string(out), "infoln message")
	assert.Contains(t, string(out), "\n")
}
