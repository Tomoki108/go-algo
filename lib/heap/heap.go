package heap

import "container/heap"

// 最小ヒープ（小さい値が優先）
type HeapItem interface {
	Priority() int
}

type Heap[T HeapItem] []T

func NewHeap[T HeapItem]() *Heap[T] {
	return &Heap[T]{}
}

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

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *Heap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.

	*h = append(*h, x.(T))
}

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeap struct {
	iarr        []int // iarrは昇順/降順になっているとは限らないため、インデックスアクセスしないこと。
	IntHeapType IntHeapType
}

func NewIntHeap(t IntHeapType) *IntHeap {
	return &IntHeap{
		iarr:        make([]int, 0),
		IntHeapType: t,
	}
}

type IntHeapType int

const (
	MinIntHeap IntHeapType = iota // 小さい方が優先して取り出される
	MaxIntHeap                    // 大きい方が優先して取り出される
)

// O(logN)
func (h *IntHeap) PushI(i int) {
	heap.Push(h, i)
}

// O(logN)
func (h *IntHeap) PopI() int {
	return heap.Pop(h).(int)
}

// to implement sort.Interface
func (h *IntHeap) Len() int { return len(h.iarr) }
func (h *IntHeap) Less(i, j int) bool {
	if h.IntHeapType == MaxIntHeap {
		return h.iarr[i] > h.iarr[j]
	} else {
		return h.iarr[i] < h.iarr[j]
	}
}
func (h *IntHeap) Swap(i, j int) { h.iarr[i], h.iarr[j] = h.iarr[j], h.iarr[i] }

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.iarr = append(h.iarr, x.(int))
}

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *IntHeap) Pop() any {
	oldiarr := h.iarr
	n := len(oldiarr)
	x := oldiarr[n-1]
	h.iarr = oldiarr[0 : n-1]
	return x
}
