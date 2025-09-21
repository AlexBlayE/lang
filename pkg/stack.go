package pkg

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Pop() T {
	finalLength := len(s.elems) - 1
	lastElement := s.elems[finalLength]
	s.elems = s.elems[0:finalLength]
	return lastElement
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}

func (s *Stack[T]) Index(i int) T {
	return s.elems[i]
}
