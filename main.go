package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"strings"
)

var loggerLevel = new(slog.LevelVar)
var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggerLevel}))

func main() {
	debugPtr := flag.Bool("debug", false, "for printing debug logs")
	flag.Parse()

	if *debugPtr {
		loggerLevel.Set(slog.LevelDebug)
	} else {
		loggerLevel.Set(slog.LevelError)
	}

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: remeta <path/to/package> <output-name>")
		return
	}

	packagePath := flag.Arg(0)
	outputName := flag.Arg(1)

	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		fmt.Printf("Could not find \"%s\"\n", packagePath)
		return
	} else if _, err := os.Stat(packagePath + "/go.mod"); os.IsNotExist(err) {
		fmt.Printf("Could not find a Go module in \"%s\"\n", packagePath)
		return
	}

	if !strings.HasSuffix(outputName, ".go") {
		outputName = outputName + ".go"
	}

	fmt.Printf("Package path: %s\n", packagePath)
	fmt.Printf("Output name: %s\n", outputName)

	fset := token.NewFileSet()

	packages, err := parser.ParseDir(fset, packagePath, nil, parser.AllErrors)
	if err != nil {
		logger.Error("%v", err)
		return
	}

	out := make([]string, 0)
	for _, pkgs := range packages {
		g := new(Grabber)
		for _, file := range pkgs.Files {
			g.GrabAll(file)
		}

		generator := NewGenerator(g, "test")
		generator.Generate()
		out = append(out, generator.OutputSource)
	}

	file, err := os.Open(outputName)
	if os.IsNotExist(err) {
		file, err = os.Create(outputName)
		if err != nil {
			logger.Error("%v", err)
			return
		}
	}
	if err != nil {
		logger.Error("%v", err)
		return
	}

	totalN := 0

	for _, content := range out {
		n, err := file.WriteString(content)
		if err != nil {
			logger.Error("%v", err)
			return
		}
		totalN += n
	}

	fmt.Printf("Wrote %d bytes to %s\n", totalN, outputName)
}
