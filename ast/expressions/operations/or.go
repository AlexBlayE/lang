package operations

import (
	"lang/runtime"
)

type OrOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (oo *OrOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(oo)
}
