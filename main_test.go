package main

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	args := []string{"--name", "test-app", "--port", "8080", "--out", outDir}
	exitCode := run(args, outStream, errStream)

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d. Stderr: %s", exitCode, errStream.String())
	}
}

func TestRunFlagParseError(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	// --invalid-flag should cause flag.Parse to fail
	args := []string{"--invalid-flag"}
	exitCode := run(args, outStream, errStream)

	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}
