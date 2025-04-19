package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/jj-mon/testgen/internal/generator"
	"github.com/jj-mon/testgen/internal/goparser"
)

func GenerateTestForFile(path string) {
	fileName := filepath.Base(path)
	dirPath := filepath.Dir(path)

	fileName, ok := strings.CutSuffix(fileName, ".go")
	if !ok {
		return
	}

	testFileName := fmt.Sprintf("%s_test.go", fileName)
	file, err := os.Create(filepath.Join(dirPath, testFileName))
	if err != nil {
		return
	}

	fileBody := generateTestForFile(path)

	_, err = file.WriteString(fileBody)
	if err != nil {
		return
	}
}

func generateTestForFile(path string) string {
	fns, mtds, packageName := goparser.ParseGoFile(path)

	file := fmt.Sprintf("package %s\n", packageName)
	file += "\n"
	file += "import (\n"
	file += "\t\"testing\"\n"
	file += "\t\"go.uber.org/mock/gomock\"\n)\n"

	for _, fn := range fns {
		file += "\n"
		file += generator.GenerateTestForFunction(fn)
	}

	for _, mtd := range mtds {
		file += "\n"
		file += generator.GenerateTestForMethod(mtd)
	}

	return file
}
