package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
)

type FuncVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (fv *FuncVisitor) Accept(s runtime.Statement) error {
	fs, ok := s.(*statements.FuncStatement)
	if !ok {
		return errors.New("invalid func statement")
	}

	fv.MemManager.Set(fs.Name, &types.Func{Val: fs})

	return nil
}
