package operations

import (
	"lang/runtime"
)

type GreaterThanOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (gto *GreaterThanOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(gto)
}
