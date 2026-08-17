package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/your-username/golden-path-cli/pkg/scaffold"
)

func main() {
	name := flag.String("name", "my-secure-service", "Name of the microservice")
	port := flag.Int("port", 3000, "Port the application listens on")
	out := flag.String("out", "./output", "Output directory for the scaffolded project")

	flag.Parse()

	cfg := scaffold.Config{
		AppName:    *name,
		Port:       *port,
		OutputPath: *out,
	}

	fmt.Printf("Scaffolding Golden Path project '%s' on port %d into %s...\n", cfg.AppName, cfg.Port, cfg.OutputPath)

	err := scaffold.Generate(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scaffolding project: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Successfully scaffolded!")
	fmt.Println("\nNext steps:")
	fmt.Printf("  cd %s\n", cfg.OutputPath)
	fmt.Println("  npm install")
	fmt.Println("  npm start")
}
