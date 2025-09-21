package types

import "fmt"

type Boolean struct {
	Val bool
}

func (b *Boolean) Debug() {
	fmt.Println("Type Boolean -> ", b.Val)
}
