package operations

import (
	"lang/runtime"
)

type MulOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (mu *MulOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(mu)
}
