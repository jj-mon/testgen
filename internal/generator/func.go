package generator

import (
	"fmt"
	"github.com/jj-mon/testgen/internal/model"
)

func generateSimpleTestForFunc(fn model.Func) string {
	testFuncName := fmt.Sprintf("Test%s", fn.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	code += "\tvar (\n"
	for _, arg := range fn.Args {
		code += fmt.Sprintf("\t\ttest%s %s\n", arg.Name, arg.TypeName)
	}
	code += "\t)\n\n"

	code += fmt.Sprintf("\t%s(\n", fn.Name)
	for _, arg := range fn.Args {
		code += fmt.Sprintf("\t\ttest%s,\n", arg.Name)
	}
	code += "\t)\n"

	code += "}\n"

	return code
}

func generateTableTestForFunc(fn model.Func) string {
	testFuncName := fmt.Sprintf("Test%s", fn.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	code += "\ttests := []struct {\n"
	code += "\t\tname\tstring\n"
	for _, arg := range fn.Args {
		code += fmt.Sprintf("\t\t%s\t%s\n", arg.Name, arg.TypeName)
	}
	code += "\t}{}\n"
	code += "\tfor _, tt := range tests {\n"
	code += "\t\tt.Run(tt.name, func(t *testing.T) {\n"

	code += "\n"

	code += fmt.Sprintf("\t\t\t%s(\n", fn.Name)
	for _, arg := range fn.Args {
		code += fmt.Sprintf("\t\t\t\ttt.%s,\n", arg.Name)
	}
	code += "\t\t\t)\n"

	code += "\t\t})\n"

	code += "\t}\n"

	code += "}\n"

	return code
}
