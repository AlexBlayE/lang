package stdlib

import (
	"fmt"
	"lang/ast/statements"
	"lang/ast/types"
	"lang/runtime"
)

type FuncInjector func(runtime.MemManager)

func ImportPrint(mm runtime.MemManager) {
	mm.Set(
		"print",
		&types.Func{Val: &statements.FuncStatement{
			Name: "print",
			Args: []statements.ArgsPair{{VarName: "s", Content: &types.String{}}},
			Block: []runtime.Statement{
				&statements.StdStatement{
					Fn: func(mm runtime.MemManager) {
						v := mm.Get("s")
						fmt.Println(v)
					},
				},
			},
		}},
	)
}
