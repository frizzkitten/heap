# heap

A generic, type-safe heap (priority queue) for Go.

Wraps `container/heap` with generics so you never have to write another `heap.Interface` implementation.

## Install

```
go get github.com/frizzkitten/heap
```

## Usage

### Min-heap

```go
h := heap.NewMin([]int{3, 1, 2}, func(v int) int { return v })

value, ok := h.Pop() // 1, true
value, ok = h.Pop()  // 2, true
value, ok = h.Pop()  // 3, true
value, ok = h.Pop()  // 0, false (empty)
```

### Max-heap with custom types

```go
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

value, ok := h.Pop() // {high 10}, true
value, ok = h.Pop()  // {medium 5}, true
value, ok = h.Pop()  // {low 1}, true
```

### Starting empty

```go
h := heap.NewMin(nil, func(v int) int { return v })

h.Push(5)
h.Push(2)
h.Push(8)

value, ok := h.Pop() // 2, true
```

## API

| Method | Description | Time complexity |
|--------|-------------|-----------------|
| `NewMin(startingValues, getPriority)` | Create a min-heap | O(n) |
| `NewMax(startingValues, getPriority)` | Create a max-heap | O(n) |
| `Push(value)` | Add an element | O(log n) |
| `Pop()` | Remove and return the highest-priority element | O(log n) |
| `Length()` | Return the number of elements | O(1) |

## Note

`startingValues` is reordered in place during heap construction. If you need to preserve the original order, pass a copy.

## License

[MIT](LICENSE)
