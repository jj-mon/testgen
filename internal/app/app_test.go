package app

import "testing"

func TestGenerateTestForFile(t *testing.T) {
	GenerateTestForFile("../generator/generator.go")
}
