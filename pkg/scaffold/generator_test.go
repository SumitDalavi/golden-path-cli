package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
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

	// Verify output directory exists
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		t.Errorf("Generate() did not create output directory: %v", outDir)
	}

	// Because we don't have the real templates in the test context (since it depends on relative paths usually), 
	// this is a basic validation. In a real app we'd embed templates or mock the file system.
}
