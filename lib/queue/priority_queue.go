package queue

import "container/heap"

type Item[T any] struct {
	value    T
	priority int // 優先度（小さいほど優先される）
	index    int // ヒープ内でのインデックス
}

type PriorityQueue[T any] []*Item[T]

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}

// 要素の優先度を更新して再配置
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// for sort.Interface
func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

// for sort.Interface
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// for sort.Interface
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// for container/heap.Interface
func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.index = len(*pq)
	*pq = append(*pq, item)
}

// for container/heap.Interface
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // セーフティ
	*pq = old[0 : n-1]
	return item
}
