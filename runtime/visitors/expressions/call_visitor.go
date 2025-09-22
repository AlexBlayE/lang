package expressions_visitors

import (
	"lang/ast/expressions"
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
		// TODO: fer que args sigui una expressió que es pugui resoldre
		cv.MemManager.Set(arg.VarName, c.Args[i])
	}

	r := runtime.Runtime{
		StmDispatcher:  cv.StmDispatcher,
		ExprDispatcher: cv.ExprDispatcher,
		MemManager:     cv.MemManager,
	}
	err := r.RunProgram(fs.Block)
	if err != nil {
		// TODO: crear un error i gestionar-ho
		return nil
	}

	return &types.Number{} // TODO: en un futur gestionar el return
}
