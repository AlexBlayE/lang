package types

import (
	"fmt"
	"lang/runtime"
)

type Array struct {
	Val []runtime.Type
}

func (a *Array) String() string {
	return fmt.Sprint(a.Val)
}
