package eslog_test

import (
	"io"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{"success", "error %s", []any{os.ErrPermission}, "error permission denied"},
		{"nil error", "error %s", []any{}, ""},
		{"success with args", "error %s %s %s", []any{os.ErrInvalid, "test1", "test2"}, "error invalid argument test1 test2"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			eslog.Logger.Errorf(tst.format, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestLogIfErrorAndLogIfErrorf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	tErr := os.ErrExist
	eslog.LogIfError(tErr, eslog.Error, "an error occurred")
	eslog.LogIfErrorf(tErr, eslog.Errorf, "error: %s", "additional info")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "an error occurred")
	assert.Contains(t, string(out), "error: additional info")
}
func TestErrorf(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{"success", "error %s", []any{os.ErrPermission}, "error permission denied"},
		{"nil error", "error %s", []any{}, ""},
		{"success with args", "error %s %s %s", []any{os.ErrInvalid, "test1", "test2"}, "error invalid argument test1 test2"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			eslog.Logger.Errorf(tst.format, tst.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestError_Nonf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.Error("simple", "error")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "simple error")
}

func TestError_Ln(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.ErrorLn("errorln message")
	err = w.Close()
	assert.NoError(t, err)

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "errorln message")
	assert.Contains(t, string(out), "\n")
}
