package goparser

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jj-mon/testgen/internal/model"
)

func ParseGoFile(path string) (model.File, error) {
	var packageName string
	// Read go file
	src, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	// parse entire package that contains go file
	pkgs, err := parser.ParseDir(fset, filepath.Dir(path), func(fi fs.FileInfo) bool {
		return !fi.IsDir() && !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	// collect all exported user structures in a package
	structs := make(map[string]model.Struct)
	for _, pkg := range pkgs {
		structs = collectStructs(pkg, fset)
		packageName = pkg.Name
		break
	}

	fset2 := token.NewFileSet()

	node, err := parser.ParseFile(fset2, path, src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	// собираем все экспортируемые функции и методы в исходном файле
	fns, mtds := collectFuncsAndMethods(node, structs, fset2)

	return model.File{
		Fns:         fns,
		Mtds:        mtds,
		PackageName: packageName,
	}, nil
}

func collectFuncsAndMethods(node *ast.File, structs map[string]model.Struct, fset *token.FileSet) ([]model.Func, []model.Method) {
	var (
		fns  []model.Func
		mtds []model.Method
	)
	// проходимся по всем декларациям в файле
	for _, decl := range node.Decls {
		// если очередная это декларация функции
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if !fn.Name.IsExported() {
				continue
			}

			fnName := fn.Name.Name
			// собираем аргументы функции
			args := collectFuncParams(fn, fset)
			// собираем количество ветвений в функции (if, case)
			branchCount := countFuncBranchStmt(fn)
			// количество возвращаемых значений
			lenResults := lenFuncResults(fn)

			if fn.Recv != nil {
				// если есть получатель, то это Метод какой то структуры
				var recvName string
				for _, recv := range fn.Recv.List {
					// находим имя структуры
					switch expr := recv.Type.(type) {
					case *ast.Ident:
						recvName = expr.Name
					case *ast.StarExpr:
						if ident, ok := expr.X.(*ast.Ident); ok {
							recvName = ident.Name
						}
					}
				}
				// ищем ранее сохраненную структуру
				if s, ok := structs[recvName]; ok {
					mtds = append(mtds, model.Method{
						Func: model.Func{
							Name:            fnName,
							Args:            args,
							BranchStmtCount: branchCount,
							LenResults:      lenResults,
						},
						Struct: s,
					})
				}
			} else {
				fns = append(fns, model.Func{
					Name:            fnName,
					Args:            args,
					BranchStmtCount: branchCount,
					LenResults:      lenResults,
				})
			}
		}
	}

	return fns, mtds
}

func collectFuncParams(fn *ast.FuncDecl, fset *token.FileSet) []model.Arg {
	var args []model.Arg

	for _, param := range fn.Type.Params.List {
		for _, name := range param.Names {
			argName := name.Name
			typeArgName := exprToString(param.Type, fset)
			args = append(args, model.Arg{
				Name: argName,
				Type: typeArgName,
			})
		}
	}

	return args
}

func countFuncBranchStmt(fn *ast.FuncDecl) int {
	branchCount := 0
	if fn.Body != nil {
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch stmt := n.(type) {
			case *ast.IfStmt:
				branchCount++
			case *ast.SwitchStmt:
				for _, s := range stmt.Body.List {
					if _, ok := s.(*ast.CaseClause); ok {
						branchCount++
					}
				}
			}
			return true
		})
	}

	return branchCount
}

func lenFuncResults(fn *ast.FuncDecl) int {
	if fn.Type.Results != nil {
		return len(fn.Type.Results.List)
	}

	return 0
}

func collectStructs(pkg *ast.Package, fset *token.FileSet) map[string]model.Struct {
	interfaces := collectUserInterfaces(pkg)
	structs := make(map[string]model.Struct)

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if !typeSpec.Name.IsExported() {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				var s model.Struct
				s.Name = typeSpec.Name.Name
				for _, field := range structType.Fields.List {
					for _, name := range field.Names {
						typeName := exprToString(field.Type, fset)
						_, ok := interfaces[typeName]
						if !ok {
							continue
						}
						s.IFields = append(s.IFields, model.IField{
							Name: name.Name,
							Type: typeName,
						})
					}
				}

				structs[s.Name] = s
			}
		}
	}

	return structs
}

func collectUserInterfaces(pkg *ast.Package) map[string]struct{} {
	interfaces := make(map[string]struct{})

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				_, ok = typeSpec.Type.(*ast.InterfaceType)
				if ok {
					interfaces[typeSpec.Name.Name] = struct{}{}
				}
			}
		}
	}

	return interfaces
}

func exprToString(expr ast.Expr, fset *token.FileSet) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return ""
	}
	return buf.String()
}
