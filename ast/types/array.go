package types

import (
	"fmt"
	"lang/runtime"
)

type Array struct {
	Val []runtime.Type
}

func (b *Array) Debug() {
	fmt.Println("Type Array -> ", b.Val)
}
