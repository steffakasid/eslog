package eslog_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestPrint(t *testing.T) {
	cases := []struct {
		name     string
		args     []any
		expected string
	}{
		{"single string", []any{"test print"}, "test print"},
		{"multiple non-string args", []any{1, 2, 3}, "1 2 3"},
		{"string and error", []any{"error:", errors.New("something")}, "error:something"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)

			eslog.Print(c.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, string(out))
		})
	}
}

func TestPrintf(t *testing.T) {
	cases := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{"simple", "hello", nil, "hello"},
		{"with args", "hi %s %d", []any{"there", 7}, "hi there 7"},
		{"positional", "%[2]s %[1]s %[3]d", []any{"a", "b", 3}, "b a 3"},
		{"error arg", "err: %s", []any{errors.New("boom")}, "err: boom"},
		{"newline preserved", "l1\nl2", nil, "l1\nl2"},
	}

	for _, c := range cases {
		t.Run("TestPrintf_"+c.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			eslog.Printf(c.format, c.args...)
			err = w.Close()
			assert.NoError(t, err)

			out, err := io.ReadAll(r)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, string(out))
		})
	}
}

func TestPrintln(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{"single string", []any{"test println"}, "test println\n"},
		{"multiple non-string args", []any{1, 2, 3}, "1 2 3\n"},
		{"string and error", []any{"error:", errors.New("something")}, "error:something\n"},
		{"empty args", nil, "\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)
			eslog.Logger.SetOutput(w)

			eslog.Println(tt.args...)
			_ = w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(out))
		})
	}
}
