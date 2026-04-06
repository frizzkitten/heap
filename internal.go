package heap

import "cmp"

type hp[T any, S cmp.Ordered] struct {
	getPriority func(T) S
	less        func(S, S) bool
	values      []T
}

func (h hp[T, S]) Len() int { return len(h.values) }

func (h hp[T, S]) Less(i, j int) bool {
	return h.less(h.getPriority(h.values[i]), h.getPriority(h.values[j]))
}

func (h hp[T, S]) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *hp[T, S]) Push(x any) {
	h.values = append(h.values, x.(T))
}

func (h *hp[T, S]) Pop() any {
	finalIndex := len(h.values) - 1
	value := h.values[finalIndex]

	// zero out the value in case it's a pointer,
	// otherwise it would be a memory leak
	var zero T
	h.values[finalIndex] = zero

	h.values = h.values[:finalIndex]
	return value
}
