package statements

import (
	"fmt"
	"lang/runtime"
)

type ReturnError struct {
	Msg string
	Res runtime.Type
}

func (e ReturnError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

type ReturnStatement struct {
	Returnable runtime.Expression
}

func (ls *ReturnStatement) Visit(v runtime.StmVisitor) error {
	return v.Accept(ls)
}
