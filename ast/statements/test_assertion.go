package statements

import (
	"lang/runtime"
	"testing"
)

type TestAssertion struct {
	T       *testing.T
	VarName string
	Type    runtime.Type
}

func (stds *TestAssertion) Visit(v runtime.StmVisitor) error {
	return v.Accept(stds)
}
