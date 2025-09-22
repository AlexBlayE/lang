package expressions_visitors

import (
	"lang/ast/expressions"
	"lang/runtime"
)

type VariableVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (vv *VariableVisitor) Accept(e runtime.Expression) runtime.Type {
	v, ok := e.(*expressions.Variable)
	if !ok {
		return nil
	}

	return vv.MemManager.Get(v.Ref)
}
