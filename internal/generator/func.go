package generator

import (
	"fmt"

	"github.com/jj-mon/testgen/internal/model"
)

func generateSimpleTestForFunc(fn model.Func) string {
	testFuncName := fmt.Sprintf("Test%s", fn.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	if len(fn.Args) > 0 {
		code += "\tvar (\n"
		for _, arg := range fn.Args {
			code += fmt.Sprintf("\t\ttest%s %s\n", arg.Name, arg.Type)
		}
		code += "\t)\n\n"
	}

	// Run testing function
	code += "\t"
	for i := range fn.LenResults {
		if i < fn.LenResults-1 {
			code += "_, "
		} else {
			code += "_ = "
		}
	}
	code += fmt.Sprintf("%s(", fn.Name)
	for _, arg := range fn.Args {
		code += fmt.Sprintf("test%s, ", arg.Name)
	}
	code += ")\n"

	code += "}\n"

	return code
}

func generateTableTestForFunc(fn model.Func) string {
	testFuncName := fmt.Sprintf("Test%s", fn.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	code += "\ttests := []struct {\n"
	code += "\t\tname\tstring\n"
	for _, arg := range fn.Args {
		code += fmt.Sprintf("\t\t%s\t%s\n", arg.Name, arg.Type)
	}
	code += "\t}{}\n"
	code += "\tfor _, tt := range tests {\n"
	code += "\t\tt.Run(tt.name, func(t *testing.T) {\n"

	code += "\n"

	// Run testing function
	code += "\t\t\t"
	for i := range fn.LenResults {
		if i < fn.LenResults-1 {
			code += "_, "
		} else {
			code += "_ = "
		}
	}
	code += fmt.Sprintf("%s(", fn.Name)
	for _, arg := range fn.Args {
		code += fmt.Sprintf("tt.%s, ", arg.Name)
	}
	code += ")\n"

	code += "\t\t})\n"

	code += "\t}\n"

	code += "}\n"

	return code
}
