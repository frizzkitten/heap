package heap_test

import (
	"fmt"

	"github.com/frizzkitten/heap"
)

func ExampleNewMin() {
	h := heap.NewMin([]int{3, 1, 2}, func(v int) int { return v })

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	// Output:
	// 1 true
	// 2 true
	// 3 true
	// 0 false
}

func ExampleNewMax() {
	type Job struct {
		Name     string
		Priority int
	}

	jobs := []Job{
		{Name: "low", Priority: 1},
		{Name: "high", Priority: 10},
		{Name: "medium", Priority: 5},
	}

	getPriority := func(j Job) int { return j.Priority }

	h := heap.NewMax(jobs, getPriority)

	value, ok := h.Pop()
	fmt.Println(value.Name, ok)

	value, ok = h.Pop()
	fmt.Println(value.Name, ok)

	value, ok = h.Pop()
	fmt.Println(value.Name, ok)

	value, ok = h.Pop()
	fmt.Println(value.Name, ok)
	// Output:
	// high true
	// medium true
	// low true
	//  false
}
