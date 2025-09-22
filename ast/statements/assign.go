package statements

import (
	"lang/runtime"
)

type AssignStatement struct {
	Name string
	Elem runtime.Expression
}

func (as *AssignStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(as)
}
