package types

import "fmt"

type String struct {
	Val string
}

func (s *String) String() string {
	return fmt.Sprint(s.Val)
}
