package expressions_visitors

import (
	"lang/ast/expressions/operations"
	"lang/ast/types"
	"lang/runtime"
)

type AddVisitor struct {
	runtime.ExprDispatcher
	runtime.MemManager
}

func (av *AddVisitor) Accept(s runtime.Expression) runtime.Type {
	ao := s.(*operations.AddOperation)

	l := ao.L
	lv := av.ExprDispatcher(l, av.MemManager)
	lt, ok := l.Visit(lv).(*types.Number)
	if !ok {
		return nil
	}

	r := ao.R
	rv := av.ExprDispatcher(r, av.MemManager)
	rt, ok := r.Visit(rv).(*types.Number)
	if !ok {
		return nil
	}

	// FIXME: com sumo lt + rt si per algún casual no son nuemros?(de moment obligo a que signuin nuemros)

	return &types.Number{Val: lt.Val + rt.Val}
}
