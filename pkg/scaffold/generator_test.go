package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	// Create dummy templates for the test
	os.MkdirAll("templates", 0755)
	defer os.RemoveAll("templates")
	os.WriteFile("templates/Dockerfile.tmpl", []byte("FROM alpine"), 0644)
	os.WriteFile("templates/k8s-deployment.yaml.tmpl", []byte("apiVersion: apps/v1"), 0644)
	os.WriteFile("templates/k8s-service.yaml.tmpl", []byte("apiVersion: v1"), 0644)
	os.WriteFile("templates/ci.yml.tmpl", []byte("name: CI"), 0644)
	os.WriteFile("templates/index.js.tmpl", []byte("console.log('test')"), 0644)
	os.WriteFile("templates/package.json.tmpl", []byte("{}"), 0644)

	cfg := Config{
		AppName:    "test-app",
		Port:       8080,
		OutputPath: outDir,
	}

	err := Generate(cfg)
	if err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		t.Errorf("Generate() did not create output directory: %v", outDir)
	}
}

func TestGenerateMissingTemplate(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	// Missing templates will cause ParseFiles to fail
	os.RemoveAll("templates")

	cfg := Config{
		AppName:    "test-app",
		OutputPath: outDir,
	}

	err := Generate(cfg)
	if err == nil {
		t.Fatal("Expected error due to missing template, got nil")
	}
}

func TestGenerateBadOutputDir(t *testing.T) {
	// A file where a directory should be will cause MkdirAll to fail
	tmpFile, _ := os.CreateTemp("", "bad-dir")
	defer os.Remove(tmpFile.Name())

	cfg := Config{
		AppName:    "test-app",
		OutputPath: tmpFile.Name(), // This is a file, so MkdirAll should fail
	}

	err := Generate(cfg)
	if err == nil {
		t.Fatal("Expected error creating output dir, got nil")
	}
}

func TestGenerateExecuteError(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	os.MkdirAll("templates", 0755)
	defer os.RemoveAll("templates")
	// Malformed template causes Execute to fail
	os.WriteFile("templates/Dockerfile.tmpl", []byte("{{ .InvalidField }}"), 0644)

	cfg := Config{
		AppName:    "test-app",
		OutputPath: outDir,
	}

	err := Generate(cfg)
	if err == nil {
		t.Fatal("Expected error executing template, got nil")
	}
}

func TestGenerateCreateFileError(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	os.MkdirAll("templates", 0755)
	defer os.RemoveAll("templates")
	os.WriteFile("templates/Dockerfile.tmpl", []byte("FROM alpine"), 0644)

	// Create outDir and a directory named Dockerfile so os.Create fails
	os.MkdirAll(filepath.Join(outDir, "Dockerfile"), 0755)

	cfg := Config{
		AppName:    "test-app",
		OutputPath: outDir,
	}

	err := Generate(cfg)
	if err == nil {
		t.Fatal("Expected error creating file, got nil")
	}
}

func TestGenerateNestedDirError(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

	os.MkdirAll("templates", 0755)
	defer os.RemoveAll("templates")
	os.WriteFile("templates/k8s-deployment.yaml.tmpl", []byte("apiVersion: apps/v1"), 0644)

	os.MkdirAll(outDir, 0755)
	// Make outDir read-only so nested MkdirAll fails
	os.Chmod(outDir, 0500)
	defer os.Chmod(outDir, 0755) // restore so cleanup works

	cfg := Config{
		AppName:    "test-app",
		OutputPath: outDir,
	}

	err := Generate(cfg)
	if err == nil {
		t.Fatal("Expected error creating nested dir, got nil")
	}
}
