package types

import (
	"fmt"
	"lang/ast/statements"
)

type Func struct {
	Val *statements.FuncStatement
}

func (f *Func) String() string {
	return fmt.Sprint(f.Val)
}
