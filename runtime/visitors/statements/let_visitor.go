package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type LetVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (lv *LetVisitor) Accept(s runtime.Statement) error {
	l, ok := s.(*statements.LetStatement)
	if !ok {
		return errors.New("invalid let statement in let visitor")
	}

	ev := lv.ExprDispatcher(l.Elem, lv.MemManager, lv.StmDispatcher)

	t := l.Elem.Visit(ev) // t can be nil
	lv.MemManager.Set(l.Name, t)

	return nil
}
