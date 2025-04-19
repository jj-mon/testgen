package main

import (
	"log"
	"os"
	"testgen/internal/app"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("usage: testgen <file.go>")
	}

	filePath := os.Args[1]

	println(filePath)

	app.GenerateTestForFile(filePath)
}
