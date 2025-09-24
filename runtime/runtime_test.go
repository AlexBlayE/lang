package runtime_test

import (
	"lang/ast/expressions"
	"lang/ast/expressions/operations"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
	"lang/runtime/dispatchers"
	"lang/runtime/memorymanager"
	"lang/stdlib"
	"testing"
)

func TestLetProgram(t *testing.T) {
	r := runtime.Runtime{
		StmDispatcher:  dispatchers.MapStatementDispatcher,
		ExprDispatcher: dispatchers.MapExpressionDispatcher,
		MemManager:     &memorymanager.MapMemManager{}}

	program := []runtime.Statement{
		&statements.LetStatement{Name: "var1", Elem: &expressions.Literal{
			Type: &types.Number{Val: 5},
		}},

		&statements.LetStatement{Name: "var2", Elem: &operations.AddOperation{
			L: &expressions.Literal{Type: &types.Number{Val: 5}},
			R: &expressions.Variable{Ref: "var1"},
		}},
	}

	err := r.RunProgram(program)
	if err != nil {
		t.Error("Runtime test failed at run program")
	}
}

func TestFunctionProgram(t *testing.T) {
	r := runtime.Runtime{
		StmDispatcher:  dispatchers.MapStatementDispatcher,
		ExprDispatcher: dispatchers.MapExpressionDispatcher,
		MemManager:     &memorymanager.MapMemManager{}}

	program := []runtime.Statement{
		&statements.LetStatement{
			Name: "count",
			Elem: &expressions.Literal{Type: &types.Number{Val: 0}},
		},

		&statements.FuncStatement{
			Name: "f1",
			Args: []statements.ArgsPair{
				{VarName: "num", Content: &types.Number{}},
			},
			Block: []runtime.Statement{
				&statements.AssignStatement{
					Name: "count",
					Elem: &expressions.Variable{Ref: "num"},
				},
			},
		},

		&statements.ExprStatement{Expr: &expressions.Call{
			Ref: "f1",
			Args: []runtime.Expression{
				&expressions.Literal{Type: &types.Number{Val: 200}},
			},
		}},

		&statements.TestAssertion{ //
			T:       t,
			VarName: "count",
			Type:    &types.Number{Val: 200},
		},
	}

	err := r.RunProgram(program)
	if err != nil {
		t.Error("Function runtime test failed")
	}
}

func TestStdlibPrint(t *testing.T) {
	mmanager := &memorymanager.MapMemManager{}

	mmanager.PushScope()
	defer mmanager.PopScope()
	stdlib.ImportPrint(mmanager)

	r := runtime.Runtime{
		StmDispatcher:  dispatchers.MapStatementDispatcher,
		ExprDispatcher: dispatchers.MapExpressionDispatcher,
		MemManager:     mmanager}

	program := []runtime.Statement{
		&statements.LetStatement{
			Name: "text",
			Elem: &expressions.Literal{Type: &types.String{Val: "Hello World"}},
		},

		&statements.ExprStatement{Expr: &expressions.Call{
			Ref: "print",
			Args: []runtime.Expression{
				&expressions.Variable{Ref: "text"},
			},
		}},
	}

	err := r.RunProgram(program)
	if err != nil {
		t.Error("Function runtime test failed")
	}
}

func TestReturnFunctions(t *testing.T) {
	r := runtime.Runtime{
		StmDispatcher:  dispatchers.MapStatementDispatcher,
		ExprDispatcher: dispatchers.MapExpressionDispatcher,
		MemManager:     &memorymanager.MapMemManager{}}

	program := []runtime.Statement{
		&statements.FuncStatement{
			Name: "f1",
			Args: []statements.ArgsPair{
				{VarName: "num", Content: &types.Number{}},
			},
			Block: []runtime.Statement{
				&statements.ReturnStatement{
					Returnable: &operations.AddOperation{
						L: &expressions.Literal{Type: &types.Number{Val: 12.5}},
						R: &expressions.Variable{Ref: "num"},
					},
				},
			},
		},

		&statements.LetStatement{
			Name: "res",
			Elem: &expressions.Call{
				Ref: "f1",
				Args: []runtime.Expression{
					&expressions.Literal{Type: &types.Number{Val: 12.5}},
				},
			},
		},

		&statements.TestAssertion{ //
			T:       t,
			VarName: "res",
			Type:    &types.Number{Val: 25},
		},
	}

	err := r.RunProgram(program)
	if err != nil {
		t.Error("Function runtime test failed")
	}
}

// TODO: test operations

// TODO: test for
