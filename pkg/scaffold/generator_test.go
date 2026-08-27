package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

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

	// Verify a few files exist
	expectedFiles := []string{
		"Dockerfile",
		"package.json",
		"k8s/deployment.yaml",
	}
	for _, f := range expectedFiles {
		if _, err := os.Stat(filepath.Join(outDir, filepath.FromSlash(f))); os.IsNotExist(err) {
			t.Errorf("Generate() did not create expected file: %v", f)
		}
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

func TestGenerateCreateFileError(t *testing.T) {
	tmpDir := t.TempDir()
	outDir := filepath.Join(tmpDir, "output")

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
