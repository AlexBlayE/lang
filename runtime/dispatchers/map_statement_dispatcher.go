package dispatchers

import (
	"lang/ast/statements"
	"lang/runtime"
	visitors "lang/runtime/visitors/statements"
	"reflect"
)

type memStmInjector = func(runtime.MemManager, runtime.ExprDispatcher, runtime.StmDispatcher) runtime.StmVisitor

var mapStatementDispatcherMap map[reflect.Type]memStmInjector = map[reflect.Type]memStmInjector{
	reflect.TypeOf(&statements.LetStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.LetVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.FuncStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.FuncVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.AssignStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.AssignVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.ExprStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.ExprVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.StdStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.StdVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.TestAssertion{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.TestAssertionVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.ReturnStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.ReturnVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
	reflect.TypeOf(&statements.ForStatement{}): func(m runtime.MemManager, ed runtime.ExprDispatcher, sd runtime.StmDispatcher) runtime.StmVisitor {
		return &visitors.ForVisitor{MemManager: m, ExprDispatcher: ed, StmDispatcher: sd}
	},
}

func MapStatementDispatcher(
	s runtime.Statement,
	mm runtime.MemManager,
	expDis runtime.ExprDispatcher,
) runtime.StmVisitor {
	return mapStatementDispatcherMap[reflect.TypeOf(s)](mm, expDis, MapStatementDispatcher)
}
