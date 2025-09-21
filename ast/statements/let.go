package statements

import (
	"lang/runtime"
)

type LetStatement struct {
	Name string
	Elem runtime.Expression
}

func (ls *LetStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(ls)
}
