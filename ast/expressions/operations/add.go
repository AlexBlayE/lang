package operations

import (
	"lang/runtime"
)

type AddOperation struct {
	L runtime.Expression
	R runtime.Expression
}

func (ao *AddOperation) Visit(ev runtime.ExprVisitor) runtime.Type {
	return ev.Accept(ao)
}
