package statements_visitors

import (
	"errors"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
)

type ForVisitor struct {
	runtime.MemManager
	runtime.ExprDispatcher
	runtime.StmDispatcher
}

func (ev *ForVisitor) Accept(s runtime.Statement) error {
	f, ok := s.(*statements.ForStatement)
	if !ok {
		return errors.New("isnt a for statement")
	}

	switch f.Obj.(type) {
	case *types.Number:
		ev.doNumber(f)
	case *types.Array:
		return errors.New("todo")
	default:
		return errors.New("need a type number or array")
	}

	return nil
}

func (ev *ForVisitor) doNumber(f *statements.ForStatement) error {
	r := runtime.Runtime{
		StmDispatcher:  ev.StmDispatcher,
		ExprDispatcher: ev.ExprDispatcher,
		MemManager:     ev.MemManager,
	}

	for i := range int(f.Obj.(*types.Number).Val) + 1 {
		ev.MemManager.Set(f.I, &types.Number{Val: float64(i)})

		err := r.RunProgram(f.Blok)
		if err != nil {
			return err
		}
	}

	return nil
}
