package tmpl

import "github.com/jj-mon/testgen/internal/model"

var TmplSimpleTestForMethod = `func Test{{.StructName}}_{{.MethodName}}(t *testing.T) {
	{{- if .HasMocks }}
	ctrl := gomock.NewController(t)

	{{- range .Mocks }}
	mock{{.Name}} := NewMock{{.Type}}(ctrl)
	{{- end }}
	{{ end }}

	{{.StructVar}} := &{{.StructName}}{
		{{- range .Mocks }}
		{{.Name}}: mock{{.Name}},
		{{- end }}
	}

	{{- if .Args }}
	var (
		{{- range .Args }}
		test{{.Name}} {{.Type}}
		{{- end }}
	)
	{{ end }}

	{{- if .HasResults }}
	{{ range $i, $_ := .Results }}
	_{{if ne $i (sub1 $.ResultsCount)}}, {{end}}
	{{- end }}= 
	{{ end }}{{.StructVar}}.{{.MethodName}}(
		{{- range $i, .Args }}
		test{{.Name}}{{if ne $i (sub1 $.ArgsCount)}}, {{end}}
		{{- end }}
	)
}`

var TmplTableTestForMethod = `func Test{{.StructName}}_{{.MethodName}}(t *testing.T) {
	tests := []struct {
		name string
		{{- range .Args }}
		{{ .Name }} {{ .Type }}
		{{- end }}

		{{- if .HasMocks }}
		setupMock func(t *testing.T, ctrl *gomock.Controller{{range .Mocks}}, {{.VarName}} *Mock{{.Type}}{{end}})
		{{- end }}
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{{- if .HasMocks }}
			ctrl := gomock.NewController(t)

			{{- range .Mocks }}
			mock{{.Name}} := NewMock{{.Type}}(ctrl)
			{{- end }}

			tt.setupMock(t, ctrl{{range .Mocks}}, mock{{.Name}}{{end}})

			{{ end }}
			{{ .StructVar }} := &{{ .StructName }}{
				{{- range .Mocks }}
				{{ .Name }}: mock{{ .Name }},
				{{- end }}
			}

			{{- if .HasResults }}
			{{- range $i, $_ := .Results }}
			_{{if ne $i (sub1 $.ResultsCount)}}, {{end}}
			{{- end }}= 
			{{ end }}{{ .StructVar }}.{{ .MethodName }}(
				{{- range $i, .Args }}
				tt.{{ .Name }}{{if ne $i (sub1 $.ArgsCount)}}, {{end}}
				{{- end }}
			)
		})
	}
}`

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
