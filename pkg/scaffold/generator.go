package scaffold

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type Config struct {
	AppName    string
	Port       int
	OutputPath string
}

func Generate(cfg Config) error {
	// Create output directory
	if err := os.MkdirAll(cfg.OutputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// Define template to output mapping
	files := map[string]string{
		"templates/Dockerfile.tmpl":          "Dockerfile",
		"templates/k8s-deployment.yaml.tmpl": "k8s/deployment.yaml",
		"templates/k8s-service.yaml.tmpl":    "k8s/service.yaml",
		"templates/ci.yml.tmpl":              ".github/workflows/ci.yml",
		"templates/index.js.tmpl":            "index.js",
		"templates/package.json.tmpl":        "package.json",
	}

	for tmplPath, outPath := range files {
		fullOutPath := filepath.Join(cfg.OutputPath, outPath)

		// Ensure parent dir exists for nested files
		if err := os.MkdirAll(filepath.Dir(fullOutPath), 0755); err != nil {
			return fmt.Errorf("failed to create parent dir for %s: %w", outPath, err)
		}

		// Parse the template from the embedded filesystem
		tmpl, err := template.ParseFS(templateFS, tmplPath)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
		}

		// Create the output file
		f, err := os.Create(fullOutPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outPath, err)
		}

		// Execute the template with our config variables
		err = tmpl.Execute(f, cfg)
		f.Close()
		if err != nil {
			return fmt.Errorf("failed to execute template %s: %w", tmplPath, err)
		}
	}

	return nil
}
