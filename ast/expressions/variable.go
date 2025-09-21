package expressions

import "lang/runtime"

type Variable struct {
	Ref string
}

func (v *Variable) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(v)
}
