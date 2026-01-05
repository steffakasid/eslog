package eslog_test

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestDebug(t *testing.T) {
	// I want to table test the debug function
	tests := []struct {
		name     string
		message  string
		args     []any
		expected string
	}{
		{"success", "test debug", []any{}, "test debug"},
		{"success with args", "test debug %s %s", []any{"arg1", "arg2"}, "test debug arg1 arg2"},
		{"success with args positions", "test debug %[2]s %[1]s %[3]s", []any{"arg1", "arg2", "arg3"}, "test debug arg2 arg1 arg3"},
		{"success with error", "test debug %s", []any{errors.New("something")}, "test debug something"},
		{"success with error and args", "test debug %s %s", []any{errors.New("something"), "arg1"}, "test debug something arg1"},
		{"success with error and args positions", "test debug %[2]s %[1]s %[3]s", []any{errors.New("something"), "arg1", "arg2"}, "test debug arg1 something arg2"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			err = eslog.Logger.SetLogLevel("Debug")
			assert.NoError(t, err)
			eslog.Debugf(tst.message, tst.args...)
			w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestInfo(t *testing.T) {
	// I want to table test the info function
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
			w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}
func TestWarn(t *testing.T) {
	// I want to table test the warn function
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
			w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{"success", "error %s", []any{errors.New("something")}, "error something"},
		{"nil error", "error %s", []any{}, ""},
		{"success with args", "error %s %s %s", []any{errors.New("test"), "test1", "test2"}, "error test test1 test2"},
		{"success with args positions", "error %[2]s %[1]s %[3]s", []any{errors.New("test"), "test1", "test2"}, "error test1 test test2"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			eslog.Logger.SetOutput(w)
			eslog.Logger.Errorf(tst.format, tst.args...)
			w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)

			assert.Contains(t, string(out), tst.expected)
		})
	}
}

func TestFatalfWithFork(t *testing.T) {
	if os.Getenv("TEST_FATAL") == "1" {
		eslog.Logger.Fatalf("fatal %s", "test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalfWithFork")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err) // Expect an error because os.Exit was called
	assert.Contains(t, string(out), "fatal test")
}

func TestSetOutput(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	err = eslog.Logger.SetLogLevel("Debug")
	assert.NoError(t, err)
	eslog.Logger.Infof("test output")
	w.Close()

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "test output")
}

func TestSetLogLevel(t *testing.T) {
	err := eslog.Logger.SetLogLevel("Info")
	assert.NoError(t, err)

	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.Debug("this should not appear")
	eslog.Info("this should appear")
	w.Close()

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.NotContains(t, string(out), "this should not appear")
	assert.Contains(t, string(out), "this should appear")
}

func TestLogIfError(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	errToLog := errors.New("test error")
	eslog.LogIfError(errToLog, eslog.Error, "an error occurred")
	w.Close()

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "an error occurred")
}

func TestLogIfErrorf(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	errToLog := errors.New("test error")
	eslog.LogIfErrorf(errToLog, eslog.Errorf, "error: %s", "additional info")
	w.Close()

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "error: additional info")
}

func TestCustomLogLevel(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	eslog.Logger.SetOutput(w)
	eslog.Logger.Log(context.Background(), eslog.LevelFatal, "fatal level log")
	w.Close()

	out, err := io.ReadAll(r)
	assert.NoError(t, err)

	assert.Contains(t, string(out), "FATAL")
	assert.Contains(t, string(out), "fatal level log")
}

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
			w.Close()

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
			w.Close()

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
		{"string and error", []any{"error:", errors.New("something")}, "error: something\n"},
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
