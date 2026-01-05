package eslog_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/steffakasid/eslog"
	"github.com/steffakasid/eslog/internal/assert"
)

func TestFatalfWithFork(t *testing.T) {
	if os.Getenv("TEST_FATAL") == "1" {
		eslog.Logger.Fatalf("fatal %s", "test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalfWithFork")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "fatal test")
}

func TestFatalAndFatalLnWithFork(t *testing.T) {
	if os.Getenv("TEST_FATAL") == "1" {
		eslog.Logger.Fatal("fatal test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalWithFork")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "fatal test")
}

func TestFatalLnWithFork(t *testing.T) {
	if os.Getenv("TEST_FATAL_LN") == "1" {
		eslog.Logger.FatalLn("fatal ln test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalLnWithFork")
	cmd.Env = append(os.Environ(), "TEST_FATAL_LN=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "fatal ln test")
}

func TestFatalWithFork(t *testing.T) {
	if os.Getenv("TEST_FATAL") == "1" {
		eslog.Logger.Fatal("fatal test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalWithFork")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "fatal test")
}

// Package-level Fatal function tests
func TestPackageFatalWithFork(t *testing.T) {
	if os.Getenv("TEST_PKG_FATAL") == "1" {
		eslog.Fatal("package fatal test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestPackageFatalWithFork")
	cmd.Env = append(os.Environ(), "TEST_PKG_FATAL=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "package fatal test")
}

func TestPackageFatalfWithFork(t *testing.T) {
	if os.Getenv("TEST_PKG_FATALF") == "1" {
		eslog.Fatalf("package fatal %s", "formatted")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestPackageFatalfWithFork")
	cmd.Env = append(os.Environ(), "TEST_PKG_FATALF=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "package fatal formatted")
}

func TestPackageFatalLnWithFork(t *testing.T) {
	if os.Getenv("TEST_PKG_FATALLN") == "1" {
		eslog.FatalLn("package fatalln test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestPackageFatalLnWithFork")
	cmd.Env = append(os.Environ(), "TEST_PKG_FATALLN=1")
	out, err := cmd.CombinedOutput()

	assert.IsError(t, err)
	assert.Contains(t, string(out), "package fatalln test")
}
