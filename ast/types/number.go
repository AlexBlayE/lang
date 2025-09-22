package types

import "fmt"

type Number struct {
	Val float64
}

func (n *Number) String() string {
	return fmt.Sprint(n.Val)
}
