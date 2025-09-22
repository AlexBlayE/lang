package types

import "fmt"

type Boolean struct {
	Val bool
}

func (b *Boolean) String() string {
	return fmt.Sprint(b.Val)
}
