package stack

import (
	"GoDataToolkit/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int](5)
	for i := 0; i < 10; i++ {
		s.Push(i)
		assert.EqualFatalf(t, i+1, s.Len(), "")
	}
	for i := 9; i >= 0; i-- {
		v := s.Pop()
		assert.EqualFatalf(t, i, v, "")
	}
}
