package dispatchers

import (
	"lang/ast/statements"
	"lang/runtime"
	visitors "lang/runtime/visitors/statements"
	"reflect"
)

type memStmInjector = func(runtime.MemManager, runtime.ExprDispatcher) runtime.StmVisitor

var mapStatementDispatcherMap map[reflect.Type]memStmInjector = map[reflect.Type]memStmInjector{
	reflect.TypeOf(&statements.LetStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher) runtime.StmVisitor {
		return &visitors.LetVisitor{MemManager: m, ExprDispatcher: ed}
	},
}

func MapStatementDispatcher(
	s runtime.Statement,
	mm runtime.MemManager,
	expDis runtime.ExprDispatcher,
) runtime.StmVisitor {
	return mapStatementDispatcherMap[reflect.TypeOf(s)](mm, expDis)
}
