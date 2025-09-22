package statements

import (
	"lang/runtime"
)

type ExprStatement struct {
	Expr runtime.Expression
}

func (ls *ExprStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(ls)
}
