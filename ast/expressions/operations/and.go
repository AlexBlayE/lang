package operations

import (
	"lang/runtime"
)

type AndOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (ao *AndOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(ao)
}
