package statements

import (
	"lang/runtime"
)

type ForStatement struct {
	I    string
	Obj  runtime.Type
	Blok []runtime.Statement
}

func (fs *ForStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(fs)
}
