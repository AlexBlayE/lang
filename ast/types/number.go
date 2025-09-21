package types

import "fmt"

type Number struct {
	Val float64
}

func (n *Number) Debug() {
	fmt.Println("Type number -> ", n.Val)
}
