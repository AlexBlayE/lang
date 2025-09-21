package operations

import (
	"lang/runtime"
)

type LowerThanOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (lto *LowerThanOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(lto)
}
