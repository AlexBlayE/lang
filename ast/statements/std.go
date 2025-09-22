package statements

import (
	"lang/runtime"
)

type StdStatement struct {
	Fn func(runtime.MemManager)
}

func (stds *StdStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(stds)
}
