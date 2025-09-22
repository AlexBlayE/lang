package dispatchers

import (
	"lang/ast/expressions"
	"lang/ast/expressions/operations"
	"lang/runtime"
	visitor "lang/runtime/visitors/expressions"
	"reflect"
)

type memMapExprDisInjector = func(runtime.ExprDispatcher, runtime.MemManager, runtime.StmDispatcher) runtime.ExprVisitor

var mapExpressionDispatcherMap map[reflect.Type]memMapExprDisInjector = map[reflect.Type]memMapExprDisInjector{
	reflect.TypeOf(&operations.AddOperation{}): func(ed runtime.ExprDispatcher, mm runtime.MemManager, stmd runtime.StmDispatcher) runtime.ExprVisitor {
		return &visitor.AddVisitor{ExprDispatcher: ed, MemManager: mm, StmDispatcher: stmd}
	},
	reflect.TypeOf(&expressions.Literal{}): func(ed runtime.ExprDispatcher, mm runtime.MemManager, stmd runtime.StmDispatcher) runtime.ExprVisitor {
		return &visitor.LiteralVisitor{ExprDispatcher: ed, MemManager: mm, StmDispatcher: stmd}
	},
	reflect.TypeOf(&expressions.Variable{}): func(ed runtime.ExprDispatcher, mm runtime.MemManager, stmd runtime.StmDispatcher) runtime.ExprVisitor {
		return &visitor.VariableVisitor{ExprDispatcher: ed, MemManager: mm, StmDispatcher: stmd}
	},
	reflect.TypeOf(&expressions.Call{}): func(ed runtime.ExprDispatcher, mm runtime.MemManager, stmd runtime.StmDispatcher) runtime.ExprVisitor {
		return &visitor.CallVisitor{ExprDispatcher: ed, MemManager: mm, StmDispatcher: stmd}
	},
}

func MapExpressionDispatcher(
	e runtime.Expression,
	mm runtime.MemManager,
	stmd runtime.StmDispatcher,
) runtime.ExprVisitor {
	return mapExpressionDispatcherMap[reflect.TypeOf(e)](MapExpressionDispatcher, mm, stmd)
}
