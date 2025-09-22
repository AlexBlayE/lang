package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type AssignVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (av *AssignVisitor) Accept(s runtime.Statement) error {
	a, ok := s.(*statements.AssignStatement)
	if !ok {
		return errors.New("invalid assign statement")
	}

	v := av.ExprDispatcher(a.Elem, av.MemManager, av.StmDispatcher)
	if v == nil {
		return errors.New("nil expression visitor")
	}

	t := a.Elem.Visit(v)

	av.MemManager.Set(a.Name, t)

	return nil
}
