package generator

import (
	"bytes"
	"log"
	"text/template"

	"github.com/jj-mon/testgen/internal/model"
	"github.com/jj-mon/testgen/internal/tmpl"
)

func generateSimpleTestForFunc(fn model.Func) string {
	tmplStr := tmpl.TmplTableTestForFunc

	t, err := template.New("TmplTableTestForFunc").Funcs(template.FuncMap{
		tmpl.FuncName: tmpl.Sub1,
	}).Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	data := tmpl.DataForFunc{
		FuncName:     fn.Name,
		ArgsCount:    len(fn.Args),
		ResultsCount: fn.LenResults,
		HasResults:   fn.LenResults > 0,
	}

	data.Args = append(data.Args, fn.Args...)

	for range fn.LenResults {
		data.Results = append(data.Results, struct{}{})
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func generateTableTestForFunc(fn model.Func) string {
	tmplStr := tmpl.TmplTableTestForFunc

	t, err := template.New("TmplTableTestForFunc").Funcs(template.FuncMap{
		tmpl.FuncName: tmpl.Sub1,
	}).Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	data := tmpl.DataForFunc{
		FuncName:     fn.Name,
		ArgsCount:    len(fn.Args),
		ResultsCount: fn.LenResults,
		HasResults:   fn.LenResults > 0,
	}

	data.Args = append(data.Args, fn.Args...)

	for range fn.LenResults {
		data.Results = append(data.Results, struct{}{})
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
