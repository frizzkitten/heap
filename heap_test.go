package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		h := New([]int{4, 2, 7, 4, -1, 9}, func(value int) int { return value })
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
}
