package operations

import (
	"lang/runtime"
)

type DivOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (do *DivOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(do)
}
