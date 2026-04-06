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
		assert.Equal(t, value, -1)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 2)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 4)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 4)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 7)

		h.Push(-5)
		h.Push(50)
		h.Push(8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, -5)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 9)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 50)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)
	})

	t.Run("max ints", func(t *testing.T) {
		h := NewMax([]int{4, 2, 7, 4, -1, 9}, func(value int) int { return value })

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 9)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 7)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 4)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 4)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 2)

		h.Push(50)
		h.Push(-5)
		h.Push(8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 50)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, -1)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, -5)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)
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
		assert.Equal(t, value, task{name: "alpha", difficulty: 1, done: true})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "beta", difficulty: 2, done: true})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "dishes", difficulty: 3, done: false})

		h.Push(task{name: "aardvark", difficulty: 9, done: true})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "aardvark", difficulty: 9, done: true})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "mop", difficulty: 5, done: false})

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, task{})
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
		assert.Equal(t, value, task{name: "mop", difficulty: 5, done: false})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "dishes", difficulty: 3, done: false})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "beta", difficulty: 2, done: true})

		h.Push(task{name: "zebra", difficulty: 0, done: false})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "zebra", difficulty: 0, done: false})

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, task{name: "alpha", difficulty: 1, done: true})

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, task{})
	})

	t.Run("nil starting values", func(t *testing.T) {
		h := NewMin(nil, func(value int) int { return value })
		assert.Equal(t, h.Length(), 0)

		value, found := h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)

		h.Push(3)
		h.Push(1)
		h.Push(2)
		assert.Equal(t, h.Length(), 3)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 1)
	})

	t.Run("empty starting values", func(t *testing.T) {
		h := NewMax([]int{}, func(value int) int { return value })
		assert.Equal(t, h.Length(), 0)

		value, found := h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)

		h.Push(3)
		h.Push(1)
		h.Push(2)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 3)
	})

	t.Run("single element", func(t *testing.T) {
		h := NewMin([]int{42}, func(value int) int { return value })
		assert.Equal(t, h.Length(), 1)

		value, found := h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 42)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)
	})

	t.Run("length", func(t *testing.T) {
		h := NewMin([]int{3, 1, 2}, func(value int) int { return value })
		assert.Equal(t, h.Length(), 3)

		h.Pop()
		assert.Equal(t, h.Length(), 2)

		h.Push(5)
		h.Push(6)
		assert.Equal(t, h.Length(), 4)

		h.Pop()
		h.Pop()
		h.Pop()
		h.Pop()
		assert.Equal(t, h.Length(), 0)
	})

	t.Run("peek", func(t *testing.T) {
		h := NewMin([]int{3, 1, 2}, func(value int) int { return value })

		value, found := h.Peek()
		assert.True(t, found)
		assert.Equal(t, value, 1)
		assert.Equal(t, h.Length(), 3)

		value, found = h.Peek()
		assert.True(t, found)
		assert.Equal(t, value, 1)
		assert.Equal(t, h.Length(), 3)

		h.Pop()
		value, found = h.Peek()
		assert.True(t, found)
		assert.Equal(t, value, 2)

		h.Pop()
		h.Pop()
		value, found = h.Peek()
		assert.False(t, found)
		assert.Equal(t, value, 0)
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
		assert.Equal(t, value, 9)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 8)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 5)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 3)

		value, found = h.Pop()
		assert.True(t, found)
		assert.Equal(t, value, 1)

		value, found = h.Pop()
		assert.False(t, found)
		assert.Equal(t, value, 0)
	})
}
