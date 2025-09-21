package types

import "fmt"

type String struct {
	Val string
}

func (n *String) Debug() {
	fmt.Println("Type String -> ", n.Val)
}
