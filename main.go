package main

import (
	"log"
	"os"

	"github.com/jj-mon/testgen/internal/app"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("usage: testgen <file.go>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	if err := app.GenerateTestsForFile(filePath); err != nil {
		log.Printf("Failed: %v", err)
	}
}
