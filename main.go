package main

import (
	"log"
	"os"

	"github.com/jj-mon/testgen/internal/app"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("usage: testgen <file.go>")
	}

	filePath := os.Args[1]

	println(filePath)

	app.GenerateTestForFile(filePath)
}
