package tmpl

import (
	"fmt"

	"github.com/jj-mon/testgen/internal/model"
)

var FuncName = "Sub1"

var TmplSimpleTestForMethod = fmt.Sprintf(`func Test{{.StructName}}_{{.MethodName}}(t *testing.T) {
	{{- if .HasMocks }}
	ctrl := gomock.NewController(t)
	{{range .Mocks }}
	mock{{.Name}} := NewMock{{.Type}}(ctrl)
	{{- end }}
	{{- end }}

	{{.StructVar}} := &{{.StructName}}{
		{{- range .Mocks }}
		{{.Name}}: mock{{.Name}},
		{{- end }}
	}
	{{ if .Args }}
	var (
		{{- range .Args }}
		test{{.Name}} {{.Type}}
		{{- end }}
	)
	{{ end }}

	{{- if .HasResults }}
	{{range $i, $_ := .Results }}_{{if ne $i (%s $.ResultsCount)}}, {{end}}{{end}} = 
	{{- end }} {{.StructVar}}.{{.MethodName}}({{range $i, $_ := .Args}}test{{.Name}}{{if ne $i (%s $.ArgsCount)}}, {{end}}{{end}})
}`, FuncName, FuncName)

var TmplTableTestForMethod = fmt.Sprintf(`func Test{{.StructName}}_{{.MethodName}}(t *testing.T) {
	tests := []struct {
		name string
		{{- range .Args }}
		{{ .Name }} {{ .Type }}
		{{- end }}

		{{- if .HasMocks }}
		setupMock func(t *testing.T, ctrl *gomock.Controller,{{range .Mocks}}{{ .Name }} *Mock{{ .Type }},{{end}})
		{{- end }}
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{- if .HasMocks }}
			ctrl := gomock.NewController(t)
			{{ range .Mocks }}
			mock{{.Name}} := NewMock{{.Type}}(ctrl)
			{{- end }}

			tt.setupMock(t, ctrl, {{range .Mocks}} mock{{.Name}},{{end}})
			{{- end }}

			{{.StructVar}} := &{{.StructName}}{
				{{- range .Mocks }}
				{{.Name}}: mock{{.Name}},
				{{- end }}
			}
			{{ if .HasResults }}
			{{range $i, $_ := .Results }}_{{if ne $i (%s $.ResultsCount)}}, {{end}}{{end}} =
			{{- end }} {{ .StructVar }}.{{ .MethodName }}({{range $i, $_ := .Args}}tt.{{.Name}}{{if le $i (%s $.ArgsCount)}}, {{end}}{{end}})
		})
	}
}`, FuncName, FuncName)

type DataForMethod struct {
	StructName   string
	MethodName   string
	StructVar    string
	Args         []model.Arg
	ArgsCount    int
	Mocks        []model.IField
	Results      []struct{}
	ResultsCount int
	HasMocks     bool
	HasResults   bool
}

var TmplSimpleTestForFunc = fmt.Sprintf(`func Test{{.FuncName}}(t *testing.T) {
	{{- if .Args }}
	var (
		{{- range .Args }}
		test{{.Name}} {{.Type}}
		{{- end }}
	)
	{{ end }}

	{{- if .HasResults }}
	{{range $i, $_ := .Results }}_{{if ne $i (%s $.ResultsCount)}}, {{end}}{{end}} = 
	{{- end }} {{.FuncName}}({{range $i, $_ := .Args}}test{{.Name}}{{if ne $i (%s $.ArgsCount)}}, {{end}}{{end}})
}`, FuncName, FuncName)

var TmplTableTestForFunc = fmt.Sprintf(`func Test{{.FuncName}}(t *testing.T) {
	tests := []struct {
		name string
		{{- range .Args }}
		{{ .Name }} {{ .Type }}
		{{- end }}
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{- if .HasResults }}
			{{range $i, $_ := .Results }}_{{if ne $i (%s $.ResultsCount)}}, {{end}}{{end}} =
			{{- end }} {{.FuncName}}({{range $i, $_ := .Args}}tt.{{.Name}}{{if le $i (%s $.ArgsCount)}}, {{end}}{{end}})
		})
	}
}`, FuncName, FuncName)

type DataForFunc struct {
	FuncName     string
	Args         []model.Arg
	ArgsCount    int
	Results      []struct{}
	ResultsCount int
	HasResults   bool
}

func Sub1(x int) int {
	return x - 1
}
