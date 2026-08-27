package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/your-username/golden-path-cli/pkg/scaffold"
)

func run(args []string, outStream, errStream io.Writer) int {
	flags := flag.NewFlagSet("golden-path-cli", flag.ContinueOnError)
	flags.SetOutput(errStream)

	name := flags.String("name", "my-secure-service", "Name of the microservice")
	port := flags.Int("port", 3000, "Port the application listens on")
	out := flags.String("out", "./output", "Output directory for the scaffolded project")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	cfg := scaffold.Config{
		AppName:    *name,
		Port:       *port,
		OutputPath: *out,
	}

	fmt.Fprintf(outStream, "Scaffolding Golden Path project '%s' on port %d into %s...\n", cfg.AppName, cfg.Port, cfg.OutputPath)

	err := scaffold.Generate(cfg)
	if err != nil {
		fmt.Fprintf(errStream, "Error scaffolding project: %v\n", err)
		return 1
	}

	fmt.Fprintln(outStream, "✅ Successfully scaffolded!")
	fmt.Fprintln(outStream, "\nNext steps:")
	fmt.Fprintf(outStream, "  cd %s\n", cfg.OutputPath)
	fmt.Fprintln(outStream, "  npm install")
	fmt.Fprintln(outStream, "  npm start")

	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
