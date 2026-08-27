package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	// We need to create dummy templates so scaffold.Generate doesn't fail
	os.MkdirAll("templates", 0755)
	defer os.RemoveAll("templates")
	os.WriteFile("templates/Dockerfile.tmpl", []byte("FROM alpine"), 0644)
	os.WriteFile("templates/k8s-deployment.yaml.tmpl", []byte("apiVersion: apps/v1"), 0644)
	os.WriteFile("templates/k8s-service.yaml.tmpl", []byte("apiVersion: v1"), 0644)
	os.WriteFile("templates/ci.yml.tmpl", []byte("name: CI"), 0644)
	os.WriteFile("templates/index.js.tmpl", []byte("console.log('test')"), 0644)
	os.WriteFile("templates/package.json.tmpl", []byte("{}"), 0644)

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

func TestRunGenerateError(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)

	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	// We DO NOT create templates here, which will cause scaffold.Generate to fail
	os.RemoveAll("templates")

	args := []string{"--out", outDir}
	exitCode := run(args, outStream, errStream)

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 due to missing templates, got %d", exitCode)
	}
}
