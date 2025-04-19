package generator

import (
	"fmt"
	"strings"

	"github.com/jj-mon/testgen/internal/model"
)

func generateSimpleTestForMethod(method model.Method) string {
	s := method.Struct

	testFuncName := fmt.Sprintf("Test%s_%s", s.Name, method.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	code += "\tctrl := gomock.NewController(t)\n\n"

	for _, iField := range s.IFields {
		code += fmt.Sprintf("\tmock%s := NewMock%s(ctrl)\n", iField.Name, iField.TypeName)
	}
	code += "\n"

	structVarName := strings.ToLower(string(s.Name[0]))

	code += fmt.Sprintf("\t%s := &%s{\n", structVarName, s.Name)
	for _, iField := range s.IFields {
		code += fmt.Sprintf("\t\t%s: mock%s,\n", iField.Name, iField.TypeName)
	}
	code += "\t}\n\n"

	code += "\tvar (\n"
	for _, arg := range method.Args {
		code += fmt.Sprintf("\t\ttest%s %s\n", arg.Name, arg.TypeName)
	}
	code += "\t)\n\n"

	code += fmt.Sprintf("\t%s.%s(\n", structVarName, method.Name)
	for _, arg := range method.Args {
		code += fmt.Sprintf("\t\ttest%s,\n", arg.Name)
	}
	code += "\t)\n"

	code += "}\n"

	return code
}

func generateTableTestForMethod(method model.Method) string {
	s := method.Struct

	testFuncName := fmt.Sprintf("Test%s_%s", s.Name, method.Name)

	code := fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName)

	code += "\ttests := []struct {\n"
	code += "\t\tname\tstring\n"
	for _, arg := range method.Args {
		code += fmt.Sprintf("\t\t%s\t%s\n", arg.Name, arg.TypeName)
	}

	if len(s.IFields) > 0 {
		code += "\t\tsetupMock\tfunc(t *testing.T, ctrl *gomock.Controller, "
		for _, iField := range s.IFields {
			code += fmt.Sprintf("%s *Mock%s, ", iField.Name, iField.TypeName)
		}
		code += ")\n"
	}
	code += "\t}{}\n"
	code += "\tfor _, tt := range tests {\n"
	code += "\t\tt.Run(tt.name, func(t *testing.T) {\n"
	code += "\t\t\tctrl := gomock.NewController(t)\n"

	for _, iField := range s.IFields {
		code += fmt.Sprintf("\t\t\tmock%s := NewMock%s(ctrl)\n", iField.Name, iField.TypeName)
	}
	code += "\n"

	if len(s.IFields) > 0 {
		code += "\t\t\ttt.setupMock(t, ctrl, "
		for _, iField := range s.IFields {
			code += fmt.Sprintf("mock%s, ", iField.Name)
		}
		code += ")\n\n"
	}

	structVarName := strings.ToLower(string(s.Name[0]))

	code += fmt.Sprintf("\t\t\t%s := &%s{\n", structVarName, s.Name)
	for _, iField := range s.IFields {
		code += fmt.Sprintf("\t\t\t\t%s: mock%s,\n", iField.Name, iField.TypeName)
	}
	code += "\t\t\t}\n\n"

	code += fmt.Sprintf("\t\t\t%s.%s(\n", structVarName, method.Name)
	for _, arg := range method.Args {
		code += fmt.Sprintf("\t\t\t\ttt.%s,\n", arg.Name)
	}
	code += "\t\t\t)\n"

	code += "\t\t})\n"

	code += "\t}\n"

	code += "}\n"

	return code
}
