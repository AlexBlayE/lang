package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type StdVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (sv *StdVisitor) Accept(s runtime.Statement) error {
	ss, ok := s.(*statements.StdStatement)
	if !ok {
		return errors.New("isnt a std visitor")
	}

	ss.Fn(sv.MemManager)

	return nil
}
