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
	value := h.values[len(h.values)-1]
	h.values = h.values[0 : len(h.values)-1]
	return value
}
