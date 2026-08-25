package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
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

	// Verify output directory exists
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		t.Errorf("Generate() did not create output directory: %v", outDir)
	}

	// Because we don't have the real templates in the test context (since it depends on relative paths usually), 
	// this is a basic validation. In a real app we'd embed templates or mock the file system.
}
