package generator

import (
	"github.com/jj-mon/testgen/internal/model"
)

func GenerateTestForFunction(fn model.Func, countForTable int) string {
	branchCount := fn.BranchStmtCount
	if branchCount < countForTable {
		return generateSimpleTestForFunc(fn)
	}

	return generateTableTestForFunc(fn)
}

func GenerateTestForMethod(method model.Method, countForTable int) string {
	branchCount := method.BranchStmtCount
	if branchCount < countForTable {
		return generateSimpleTestForMethod(method)
	}

	return generateTableTestForMethod(method)
}
