package types

import (
	"fmt"
	"lang/ast/statements"
)

type Func struct {
	Val *statements.FuncStatement
}

func (b *Func) Debug() {
	fmt.Println("Type Func -> ", b.Val)
}
