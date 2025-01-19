package deque

import "fmt"

// 先頭、末尾へのデータ追加、削除がO(1)で行える。インデックスアクセスもO(1)で可能
type Deque[T any] struct {
	data       []T
	head, tail int // 先頭と末尾のインデックス
	capacity   int // デックの容量
	size       int // 現在の要素数
}

func NewDeque[T any](initialCapacity int) *Deque[T] {
	return &Deque[T]{
		data:     make([]T, initialCapacity),
		capacity: initialCapacity,
	}
}

// O(1)
func (d *Deque[T]) PushFront(value T) {
	if d.IsFull() {
		d.resize()
	}
	// headを逆方向に進めて要素を追加
	d.head = (d.head - 1 + d.capacity) % d.capacity
	d.data[d.head] = value
	d.size++
}

// O(1)
func (d *Deque[T]) PushBack(value T) {
	if d.IsFull() {
		d.resize()
	}
	// 要素を追加し、tailを進める
	d.data[d.tail] = value
	d.tail = (d.tail + 1) % d.capacity
	d.size++
}

// O(1)
func (d *Deque[T]) PopFront() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// 要素を取得し、headを進める
	value := d.data[d.head]
	d.head = (d.head + 1) % d.capacity
	d.size--
	return value
}

// O(1)
func (d *Deque[T]) PopBack() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// tailを逆方向に進めて要素を取得
	d.tail = (d.tail - 1 + d.capacity) % d.capacity
	value := d.data[d.tail]
	d.size--
	return value
}

// O(1)
func (d *Deque[T]) Front() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	return d.data[d.head]
}

// O(1)
func (d *Deque[T]) Back() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// tailの直前の要素を取得
	return d.data[(d.tail-1+d.capacity)%d.capacity]
}

// O(1)
func (d *Deque[T]) At(index int) T {
	if index < 0 || index >= d.size {
		panic("index out of range")
	}

	physicalIndex := (d.head + index) % d.capacity
	return d.data[physicalIndex]
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) IsFull() bool {
	return d.size == d.capacity
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Dump() {
	fmt.Printf("head: %d, tail: %d, data: %v\n", d.head, d.tail, d.data)
}

// O(current size)
// デックの容量を2倍に拡張する
func (d *Deque[T]) resize() {
	newCapacity := d.capacity * 2
	newData := make([]T, newCapacity)

	// 現在のデータを新しい配列にコピー（リングバッファの順序を維持）
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.head+i)%d.capacity]
	}

	// 配列とインデックスを更新
	d.data = newData
	d.head = 0
	d.tail = d.size
	d.capacity = newCapacity
}
