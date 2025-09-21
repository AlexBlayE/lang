package expressions

import "lang/runtime"

type Literal struct {
	runtime.Type
}

func (l *Literal) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(l)
}
