package model

type Func struct {
	Name            string
	Args            []Arg
	BranchStmtCount int
	LenResults      int
}

type Arg struct {
	Name string
	Type string
}

type Method struct {
	Func
	Struct Struct
}

type Struct struct {
	Name    string
	IFields []IField
}

type IField struct {
	Name string
	Type string
}

type File struct {
	PackageName string
	Fns         []Func
	Mtds        []Method
}
