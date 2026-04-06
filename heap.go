package heap

import (
	"cmp"
	containerHeap "container/heap"
)

const (
	min = iota
	max
)

type Heap[T any, S cmp.Ordered] struct {
	internalHeap hp[T, S]
}

func NewMin[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S) Heap[T, S] {
	return newHeap(startingValues, getPriority, min)
}

func NewMax[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S) Heap[T, S] {
	return newHeap(startingValues, getPriority, max)
}

func newHeap[T any, S cmp.Ordered](startingValues []T, getPriority func(T) S, heapType int) Heap[T, S] {
	internalHeap := hp[T, S]{
		getPriority: getPriority,
		values:      startingValues,
		heapType:    heapType,
	}

	containerHeap.Init(&internalHeap)

	return Heap[T, S]{
		internalHeap: internalHeap,
	}
}

func (h *Heap[T, S]) Push(value T) {
	containerHeap.Push(&h.internalHeap, value)
}

func (h *Heap[T, S]) Pop() (T, bool) {
	if len(h.internalHeap.values) == 0 {
		var empty T
		return empty, false
	}

	return containerHeap.Pop(&h.internalHeap).(T), true
}

func (h *Heap[T, S]) Length() int {
	return len(h.internalHeap.values)
}
