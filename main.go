package main

import (
	"flag"
	"log"
	"os"

	"github.com/jj-mon/testgen/internal/app"
	"github.com/jj-mon/testgen/internal/config"
)

var (
	conditions = flag.Int("c", 3, "number of conditions for table-driven tests")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		log.Print("usage: testgen <file.go>")
		os.Exit(1)
	}

	filePath := os.Args[len(os.Args)-1]

	cfg := &config.Config{
		Conditions: *conditions,
	}

	a := app.New(cfg)

	if err := a.GenerateTestsForFile(filePath); err != nil {
		log.Printf("Failed: %v", err)
	}
}
