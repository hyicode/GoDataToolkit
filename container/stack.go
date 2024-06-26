package container

type sliceStack[T any] []T

func newSliceStack[T any](capSize int) sliceStack[T] {
	return make(sliceStack[T], 0, capSize)
}

func (s *sliceStack[T]) len() int { return len(*s) }

func (s *sliceStack[T]) pop() T {
	if len(*s) == 0 {
		var t T
		return t
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *sliceStack[T]) push(v T) {
	*s = append(*s, v)
}

type Stack[T any] struct {
	list       List[sliceStack[T]]
	bucketSize int
	len        int
}

func (s *Stack[T]) Init(bucketSize int) *Stack[T] {
	s.bucketSize = bucketSize
	s.list.Init()
	return s
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T]).Init(10)
}

func (s *Stack[T]) Len() int { return s.len }

func (s *Stack[T]) Pop() T {
	item := s.list.Back()
	v := item.Value.pop()
	if item.Value.len() <= 0 {
		s.list.Remove(item)
	}
	s.len--
	return v
}

func (s *Stack[T]) Push(v T) *Stack[T] {
	if s.len%s.bucketSize == 0 {
		item := s.list.PushBack(newSliceStack[T](s.bucketSize))
		item.Value.push(v)
		s.len++
		return s
	}
	item := s.list.Back()
	item.Value.push(v)
	s.len++
	return s
}

func (s *Stack[T]) Range(f func(v T) (stop bool)) {
	for e := s.list.Front(); e != nil; e = e.Next() {
		for _, v := range e.Value {
			if f(v) {
				return
			}
		}
	}
}

func (s *Stack[T]) RemoveFunc(num int, f func(v T) (stop bool)) {
	counter := 0
	for e := s.list.Front(); e != nil; e = e.Next() {
		var l sliceStack[T]
		for i, v := range e.Value {
			if f(v) {
				counter++
				s.len--
				if len(l) == 0 {
					l = append(l, e.Value[:i]...)
				}
				if num > 0 && counter >= num {
					l = append(l, e.Value[i+1:]...)
					e.Value = l
					if len(e.Value) == 0 {
						s.list.Remove(e)
					}
					return
				}
			} else if len(l) > 0 {
				l = append(l, v)
			}
		}
		if len(l) > 0 {
			e.Value = l
		}
	}
}
