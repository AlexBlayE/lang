package statements

import (
	"lang/runtime"
)

type ArgsPair = struct {
	VarName string
	Content runtime.Type
}

type FuncStatement struct {
	Name      string
	Args      []ArgsPair
	Block     []runtime.Statement
	CanReturn bool
}

func (ls *FuncStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(ls)
}
