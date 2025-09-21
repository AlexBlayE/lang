package statements

import (
	"lang/runtime"
)

type AssignStatement struct {
	Name string
	Elem runtime.Type
}

func (as *AssignStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(as)
}
