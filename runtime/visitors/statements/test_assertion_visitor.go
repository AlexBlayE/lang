package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/runtime"
)

type TestAssertionVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (tav *TestAssertionVisitor) Accept(s runtime.Statement) error {
	ta, ok := s.(*statements.TestAssertion)
	if !ok {
		return errors.New("invalid test assertion")
	}

	t := tav.MemManager.Get(ta.VarName)
	if t.String() != ta.Type.String() {
		ta.T.Error("invalid assertion: not equals")
		return errors.New("invalid assertion: not equals")
	}

	return nil
}
