package heap

import "container/heap"

type HeapItem interface {
	Priority() int
}

type Heap[T HeapItem] []T

func (h *Heap[T]) PushItem(item T) {
	heap.Push(h, item)
}

func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

// to implement sort.Interface
func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// to implement heap.Interface
func (h *Heap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(T))
}

// to implement heap.Interface
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
