package generator

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/jj-mon/testgen/internal/model"
	"github.com/jj-mon/testgen/internal/tmpl"
)

func generateSimpleTestForMethod(method model.Method) string {
	tmplStr := tmpl.TmplSimpleTestForMethod

	t, err := template.New("TmplSimpleTestForMethod").Funcs(template.FuncMap{
		tmpl.FuncName: tmpl.Sub1,
	}).Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	s := method.Struct
	structVar := strings.ToLower(string(s.Name[0]))

	data := tmpl.DataForMethod{
		StructName:   s.Name,
		MethodName:   method.Name,
		StructVar:    structVar,
		ArgsCount:    len(method.Args),
		ResultsCount: method.LenResults,
		HasMocks:     len(s.IFields) > 0,
		HasResults:   method.LenResults > 0,
	}

	data.Mocks = append(data.Mocks, s.IFields...)

	data.Args = append(data.Args, method.Args...)

	for range method.LenResults {
		data.Results = append(data.Results, struct{}{})
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func generateTableTestForMethod(method model.Method) string {
	tmplStr := tmpl.TmplTableTestForMethod

	t, err := template.New("TmplTableTestForMethod").Funcs(template.FuncMap{
		tmpl.FuncName: tmpl.Sub1,
	}).Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	s := method.Struct
	structVar := strings.ToLower(string(s.Name[0]))

	data := tmpl.DataForMethod{
		StructName:   s.Name,
		MethodName:   method.Name,
		StructVar:    structVar,
		ArgsCount:    len(method.Args),
		ResultsCount: method.LenResults,
		HasMocks:     len(s.IFields) > 0,
		HasResults:   method.LenResults > 0,
	}

	data.Mocks = append(data.Mocks, s.IFields...)

	data.Args = append(data.Args, method.Args...)

	for range method.LenResults {
		data.Results = append(data.Results, struct{}{})
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
