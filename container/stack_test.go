package container

import (
	"github.com/hyicode/GoDataToolkit/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	count := 100
	for i := 0; i < count; i++ {
		s.Push(i)
		assert.EqualFatalf(t, i+1, s.Len(), "push")
	}
	expectV := 0
	s.Range(func(v int) (stop bool) {
		assert.EqualFatalf(t, expectV, v, "range")
		expectV++
		return false
	})
	assert.EqualFatalf(t, expectV, count, "expectV")
	for i := count; i > 0; i-- {
		v := s.Pop()
		assert.EqualFatalf(t, i-1, v, "pop")
	}
}

func TestStack_RemoveFunc(t *testing.T) {
	tests := []struct {
		name     string
		stack    *Stack[int]
		num      int
		f        func(v int) bool
		expected []int
	}{
		{
			name:     "Remove first occurrence",
			stack:    NewStack[int]().Push(1).Push(2).Push(3),
			num:      1,
			f:        func(v int) bool { return v == 2 },
			expected: []int{1, 3},
		},
		{
			name:     "Remove multiple occurrences",
			stack:    NewStack[int]().Push(1).Push(2).Push(3).Push(2).Push(4),
			num:      2,
			f:        func(v int) bool { return v == 2 },
			expected: []int{1, 3, 4},
		},
		{
			name:     "Remove all occurrences",
			stack:    NewStack[int]().Push(1).Push(2).Push(3).Push(2).Push(4),
			num:      0,
			f:        func(v int) bool { return v == 2 },
			expected: []int{1, 3, 4},
		},
		{
			name:     "Remove non-existent element",
			stack:    NewStack[int]().Push(1).Push(2).Push(3),
			num:      1,
			f:        func(v int) bool { return v == 5 },
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty stack",
			stack:    NewStack[int](),
			num:      1,
			f:        func(v int) bool { return v == 2 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.RemoveFunc(tt.num, tt.f)
			var result []int
			tt.stack.Range(func(v int) bool {
				result = append(result, v)
				return false
			})
			assert.EqualSliceErrorf(t, tt.expected, result, "")
		})
	}
}
