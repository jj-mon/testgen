package model

type Func struct {
	Name            string
	Args            []Arg
	BranchStmtCount int
}

type Arg struct {
	Name     string
	TypeName string
}

type Method struct {
	Name            string
	Struct          Struct
	Args            []Arg
	BranchStmtCount int
}

type Struct struct {
	Name    string
	IFields []IField
}

type IField struct {
	Name     string
	TypeName string
}
