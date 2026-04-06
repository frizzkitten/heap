// Package heap provides a generic, type-safe priority queue backed by
// container/heap. It supports both min-heaps and max-heaps with a
// user-defined priority function.
package heap

import (
	"cmp"
	containerHeap "container/heap"
)

// Heap is a generic priority queue. T is the element type and S is the
// priority type used for ordering.
type Heap[T any, S cmp.Ordered] struct {
	internalHeap hp[T, S]
}

// NewMin creates a min-heap where the lowest priority value is popped first.
// startingValues will be reordered in place; pass nil to start empty.
// Runs in O(n) time where n is len(startingValues).
func NewMin[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S) Heap[T, S] {
	return newHeap(startingValues, getPriority, func(a, b S) bool { return a < b })
}

// NewMax creates a max-heap where the highest priority value is popped first.
// startingValues will be reordered in place; pass nil to start empty.
// Runs in O(n) time where n is len(startingValues).
func NewMax[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S) Heap[T, S] {
	return newHeap(startingValues, getPriority, func(a, b S) bool { return a > b })
}

func newHeap[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S, less func(S, S) bool) Heap[T, S] {
	internalHeap := hp[T, S]{
		getPriority: getPriority,
		less:        less,
		values:      startingValues,
	}

	containerHeap.Init(&internalHeap)

	return Heap[T, S]{
		internalHeap: internalHeap,
	}
}

// Push adds a value to the heap. Runs in O(log n) time.
func (h *Heap[T, S]) Push(value T) {
	containerHeap.Push(&h.internalHeap, value)
}

// Pop removes and returns the highest-priority element.
// Returns false if the heap is empty. Runs in O(log n) time.
func (h *Heap[T, S]) Pop() (T, bool) {
	if len(h.internalHeap.values) == 0 {
		var zero T
		return zero, false
	}

	return containerHeap.Pop(&h.internalHeap).(T), true
}

// Length returns the number of elements in the heap. Runs in O(1) time.
func (h *Heap[T, S]) Length() int {
	return len(h.internalHeap.values)
}
