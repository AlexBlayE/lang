package operations

import (
	"lang/runtime"
)

type MinOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (mo *MinOperation) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(mo)
}
