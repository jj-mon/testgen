package generator

import (
	"fmt"
	"strings"

	"github.com/jj-mon/testgen/internal/model"
)

func generateSimpleTestForMethod(method model.Method) string {
	s := method.Struct
	testFuncName := fmt.Sprintf("Test%s_%s", s.Name, method.Name)

	var code strings.Builder
	code.WriteString(fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName))
	if len(s.IFields) > 0 {
		code.WriteString("\tctrl := gomock.NewController(t)\n\n")
		for _, iField := range s.IFields {
			code.WriteString(fmt.Sprintf("\tmock%s := NewMock%s(ctrl)\n", iField.Name, iField.Type))
		}
		code.WriteString("\n")
	}

	structVarName := strings.ToLower(string(s.Name[0]))

	code.WriteString(fmt.Sprintf("\t%s := &%s{\n", structVarName, s.Name))
	for _, iField := range s.IFields {
		code.WriteString(fmt.Sprintf("\t\t%s: mock%s,\n", iField.Name, iField.Name))
	}
	code.WriteString("\t}\n\n")

	if len(method.Args) > 0 {
		code.WriteString("\tvar (\n")
		for _, arg := range method.Args {
			code.WriteString(fmt.Sprintf("\t\ttest%s %s\n", arg.Name, arg.Type))
		}
		code.WriteString("\t)\n\n")
	}

	// Run testing function
	code.WriteString("\t")
	for i := range method.LenResults {
		if i < method.LenResults-1 {
			code.WriteString("_, ")
		} else {
			code.WriteString("_ = ")
		}
	}
	code.WriteString(fmt.Sprintf("%s.%s(", structVarName, method.Name))
	for _, arg := range method.Args {
		code.WriteString(fmt.Sprintf("test%s, ", arg.Name))
	}
	code.WriteString(")\n")

	code.WriteString("}\n")

	return code.String()
}

func generateTableTestForMethod(method model.Method) string {
	s := method.Struct
	testFuncName := fmt.Sprintf("Test%s_%s", s.Name, method.Name)

	var code strings.Builder
	code.WriteString(fmt.Sprintf("func %s(t *testing.T) {\n", testFuncName))

	code.WriteString("\ttests := []struct {\n")
	code.WriteString("\t\tname\tstring\n")
	for _, arg := range method.Args {
		code.WriteString(fmt.Sprintf("\t\t%s\t%s\n", arg.Name, arg.Type))
	}

	if len(s.IFields) > 0 {
		code.WriteString("\t\tsetupMock\tfunc(t *testing.T, ctrl *gomock.Controller, ")
		for _, iField := range s.IFields {
			code.WriteString(fmt.Sprintf("%s *Mock%s, ", iField.Name, iField.Type))
		}
		code.WriteString(")\n")
	}
	code.WriteString("\t}{}\n")
	code.WriteString("\tfor _, tt := range tests {\n")
	code.WriteString("\t\tt.Run(tt.name, func(t *testing.T) {\n")
	if len(s.IFields) > 0 {
		code.WriteString("\t\t\tctrl := gomock.NewController(t)\n")

		code.WriteString("\n")

		for _, iField := range s.IFields {
			code.WriteString(fmt.Sprintf("\t\t\tmock%s := NewMock%s(ctrl)\n", iField.Name, iField.Type))
		}
		code.WriteString("\n")

		code.WriteString("\t\t\ttt.setupMock(t, ctrl, ")
		for _, iField := range s.IFields {
			code.WriteString(fmt.Sprintf("mock%s, ", iField.Name))
		}
		code.WriteString(")\n\n")
	}

	structVarName := strings.ToLower(string(s.Name[0]))

	code.WriteString(fmt.Sprintf("\t\t\t%s := &%s{\n", structVarName, s.Name))
	for _, iField := range s.IFields {
		code.WriteString(fmt.Sprintf("\t\t\t\t%s: mock%s,\n", iField.Name, iField.Name))
	}
	code.WriteString("\t\t\t}\n\n")

	// Run testing function
	code.WriteString("\t\t\t")
	for i := range method.LenResults {
		if i < method.LenResults-1 {
			code.WriteString("_, ")
		} else {
			code.WriteString("_ = ")
		}
	}
	code.WriteString(fmt.Sprintf("%s.%s(", structVarName, method.Name))
	for _, arg := range method.Args {
		code.WriteString(fmt.Sprintf("tt.%s, ", arg.Name))
	}
	code.WriteString(")\n")

	code.WriteString("\t\t})\n")

	code.WriteString("\t}\n")

	code.WriteString("}\n")

	return code.String()
}
