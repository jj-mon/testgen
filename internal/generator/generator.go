package generator

import (
	"github.com/jj-mon/testgen/internal/model"
)

const countForTable = 3

func GenerateTestForFunction(fn model.Func) string {
	branchCount := fn.BranchStmtCount
	if branchCount < countForTable {
		return generateSimpleTestForFunc(fn)
	}

	return generateTableTestForFunc(fn)
}

func GenerateTestForMethod(method model.Method) string {
	branchCount := method.BranchStmtCount
	if branchCount < countForTable {
		return generateSimpleTestForMethod(method)
	}

	return generateTableTestForMethod(method)
}
