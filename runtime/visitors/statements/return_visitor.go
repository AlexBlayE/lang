package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type ReturnVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (rv *ReturnVisitor) Accept(s runtime.Statement) error {
	r, ok := s.(*statements.ReturnStatement)
	if !ok {
		return errors.New("invalid return statement")
	}

	ev := rv.ExprDispatcher(r.Returnable, rv.MemManager, rv.StmDispatcher)
	if ev == nil {
		return errors.New("havent a dispatcher")
	}

	res := r.Returnable.Visit(ev)

	return statements.ReturnError{Msg: "all ok", Res: res}
}
