package expressions

import "lang/runtime"

type Call struct {
	Ref  string
	Args []runtime.Expression
}

func (v *Call) Visit(sv runtime.ExprVisitor) runtime.Type {
	return sv.Accept(v)
}
