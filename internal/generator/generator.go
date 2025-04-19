package generator

import (
	"github.com/jj-mon/testgen/internal/model"
)

func GenerateTestForFunction(fn model.Func) string {
	branchCount := fn.BranchStmtCount
	if branchCount < 2 {
		return generateSimpleTestForFunc(fn)
	}

	return generateTableTestForFunc(fn)
}

func GenerateTestForMethod(method model.Method) string {
	branchCount := method.BranchStmtCount
	if branchCount < 2 {
		return generateSimpleTestForMethod(method)
	}

	return generateTableTestForMethod(method)
}
