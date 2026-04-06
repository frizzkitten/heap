package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type task struct {
	name       string
	difficulty int
	done       bool
}

func TestHeap(t *testing.T) {
	t.Run("min ints", func(t *testing.T) {
		h := NewMin([]int{4, 2, 7, 4, -1, 9}, func(value int) int { return value })

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, -1, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 2, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 4, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 4, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 7, value)

		h.Push(-5)
		h.Push(50)
		h.Push(8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, -5, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 8, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 9, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 50, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})

	t.Run("max ints", func(t *testing.T) {
		h := NewMax([]int{4, 2, 7, 4, -1, 9}, func(value int) int { return value })

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, 9, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 7, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 4, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 4, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 2, value)

		h.Push(50)
		h.Push(-5)
		h.Push(8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 50, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 8, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, -1, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, -5, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})

	t.Run("min structs by name", func(t *testing.T) {
		h := NewMin([]task{
			{name: "dishes", difficulty: 3, done: false},
			{name: "alpha", difficulty: 1, done: true},
			{name: "mop", difficulty: 5, done: false},
			{name: "beta", difficulty: 2, done: true},
		}, func(t task) string { return t.name })

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "alpha", difficulty: 1, done: true}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "beta", difficulty: 2, done: true}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "dishes", difficulty: 3, done: false}, value)

		h.Push(task{name: "aardvark", difficulty: 9, done: true})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "aardvark", difficulty: 9, done: true}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "mop", difficulty: 5, done: false}, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, task{}, value)
	})

	t.Run("max structs by name", func(t *testing.T) {
		h := NewMax([]task{
			{name: "dishes", difficulty: 3, done: false},
			{name: "alpha", difficulty: 1, done: true},
			{name: "mop", difficulty: 5, done: false},
			{name: "beta", difficulty: 2, done: true},
		}, func(t task) string { return t.name })

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "mop", difficulty: 5, done: false}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "dishes", difficulty: 3, done: false}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "beta", difficulty: 2, done: true}, value)

		h.Push(task{name: "zebra", difficulty: 0, done: false})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "zebra", difficulty: 0, done: false}, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, task{name: "alpha", difficulty: 1, done: true}, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, task{}, value)
	})

	t.Run("nil starting values", func(t *testing.T) {
		h := NewMin(nil, func(value int) int { return value })
		assert.Equal(t, 0, h.Length())

		value, found := h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)

		h.Push(3)
		h.Push(1)
		h.Push(2)
		assert.Equal(t, 3, h.Length())

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 1, value)
	})

	t.Run("empty starting values", func(t *testing.T) {
		h := NewMax([]int{}, func(value int) int { return value })
		assert.Equal(t, 0, h.Length())

		value, found := h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)

		h.Push(3)
		h.Push(1)
		h.Push(2)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 3, value)
	})

	t.Run("single element", func(t *testing.T) {
		h := NewMin([]int{42}, func(value int) int { return value })
		assert.Equal(t, 1, h.Length())

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, 42, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})

	t.Run("length", func(t *testing.T) {
		h := NewMin([]int{3, 1, 2}, func(value int) int { return value })
		assert.Equal(t, 3, h.Length())

		h.Pop()
		assert.Equal(t, 2, h.Length())

		h.Push(5)
		h.Push(6)
		assert.Equal(t, 4, h.Length())

		h.Pop()
		h.Pop()
		h.Pop()
		h.Pop()
		assert.Equal(t, 0, h.Length())
	})

	t.Run("peek", func(t *testing.T) {
		h := NewMin([]int{3, 1, 2}, func(value int) int { return value })

		value, found := h.Peek()
		assert.True(t, found)
		assert.Equal(t, 1, value)
		assert.Equal(t, 3, h.Length())

		value, found = h.Peek()
		assert.True(t, found)
		assert.Equal(t, 1, value)
		assert.Equal(t, 3, h.Length())

		h.Pop()
		value, found = h.Peek()
		assert.True(t, found)
		assert.Equal(t, 2, value)

		h.Pop()
		h.Pop()
		value, found = h.Peek()
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})

	t.Run("push only then pop all", func(t *testing.T) {
		h := NewMax(nil, func(value int) int { return value })

		h.Push(5)
		h.Push(3)
		h.Push(8)
		h.Push(1)
		h.Push(9)

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, 9, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 8, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 5, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 3, value)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, 1, value)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, 0, value)
	})
}
