package runtime_test

import (
	"lang/ast/expressions"
	"lang/ast/expressions/operations"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
	"lang/runtime/dispatchers"
	memeorymanager "lang/runtime/memorymanager"
	"testing"
)

func TestLetProgram(t *testing.T) {
	r := runtime.Runtime{
		StmDispatcher:  dispatchers.MapStatementDispatcher,
		ExprDispatcher: dispatchers.MapExpressionDispatcher,
		MemManager:     &memeorymanager.MapMemManager{}}

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
		MemManager:     &memeorymanager.MapMemManager{}}

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
			CanReturn: false,
		},

		&statements.ExprStatement{Expr: &expressions.Call{
			Ref: "f1",
			Args: []runtime.Expression{
				&expressions.Literal{Type: &types.Number{200}},
			},
		}},
	}

	err := r.RunProgram(program)
	if err != nil {
		t.Error("Function runtime test failed")
	}
}

// TODO: stdlib print test
