package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jj-mon/testgen/internal/generator"
	"github.com/jj-mon/testgen/internal/goparser"
)

func GenerateTestsForFile(path string) error {
	fileName := filepath.Base(path)
	dirPath := filepath.Dir(path)

	suffix := ".go"
	fileName, ok := strings.CutSuffix(fileName, suffix)
	if !ok {
		return fmt.Errorf("failed cat suffix '%s' in '%s'", fileName, suffix)
	}

	testFileName := fmt.Sprintf("%s_test.go", fileName)
	file, err := os.Create(filepath.Join(dirPath, testFileName))
	if err != nil {
		return fmt.Errorf("failed create file: %v", err)
	}

	fileBody, err := generateTestsForFile(path)
	if err != nil {
		return fmt.Errorf("failed generate tests: %v", err)
	}

	_, err = file.WriteString(fileBody)
	if err != nil {
		return fmt.Errorf("failed write file: %v", err)
	}

	return nil
}

func generateTestsForFile(path string) (string, error) {
	fileModel, err := goparser.ParseGoFile(path)
	if err != nil {
		return "", err
	}

	fns, mtds, packageName := fileModel.Fns, fileModel.Mtds, fileModel.PackageName

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

	return file, nil
}
