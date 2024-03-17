package stack

type Stack[T any] []T

func NewStack[T any](expectSize int) Stack[T] {
	return make(Stack[T], 0, expectSize)
}

func (s *Stack[T]) Len() int { return len(*s) }

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var t T
		return t
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}
