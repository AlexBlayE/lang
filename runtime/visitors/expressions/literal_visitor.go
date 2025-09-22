package expressions_visitors

import (
	"lang/ast/expressions"
	"lang/runtime"
)

type LiteralVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (vv *LiteralVisitor) Accept(e runtime.Expression) runtime.Type {
	l, ok := e.(*expressions.Literal)
	if !ok {
		return nil
	}

	return l.Type
}
