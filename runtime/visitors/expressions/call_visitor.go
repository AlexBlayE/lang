package expressions_visitors

import (
	"errors"
	"lang/ast/expressions"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
)

type CallVisitor struct {
	runtime.ExprDispatcher
	runtime.MemManager
	runtime.StmDispatcher
}

func (cv *CallVisitor) Accept(s runtime.Expression) runtime.Type {
	c, ok := s.(*expressions.Call)
	if !ok {
		return nil
	}

	f, ok := cv.MemManager.Get(c.Ref).(*types.Func)
	if !ok {
		return nil
	}

	fs := f.Val

	cv.MemManager.PushScope()
	defer cv.MemManager.PopScope()

	for i, arg := range fs.Args {
		ex := c.Args[i]

		ed := cv.ExprDispatcher(ex, cv.MemManager, cv.StmDispatcher)
		if ed == nil {
			return nil
		}

		t := ex.Visit(ed)

		cv.MemManager.Set(arg.VarName, t)
	}

	r := runtime.Runtime{
		StmDispatcher:  cv.StmDispatcher,
		ExprDispatcher: cv.ExprDispatcher,
		MemManager:     cv.MemManager,
	}

	err := r.RunProgram(fs.Block)
	if errors.As(err, &statements.ReturnError{}) {
		re := err.(statements.ReturnError)
		return re.Res
	}

	if err != nil {
		return nil
	}

	return nil
}
