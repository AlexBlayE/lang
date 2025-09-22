package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type ExprVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (ev *ExprVisitor) Accept(s runtime.Statement) error {
	es, ok := s.(*statements.ExprStatement)
	if !ok {
		return errors.New("isnt a expression")
	}

	v := ev.ExprDispatcher(es.Expr, ev.MemManager, ev.StmDispatcher)
	if v == nil {
		return errors.New("nil expression visitor")
	}

	es.Expr.Visit(v)

	return nil
}
