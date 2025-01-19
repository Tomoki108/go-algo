package deque

import "fmt"

// Deque構造
type Deque[T any] struct {
	data         []T // 固定サイズ配列
	head, tail   int // 先頭と末尾のインデックス
	capacity     int // 配列の容量
	elementCount int // 現在の要素数
}

// 新しいDequeを作成
func NewDeque[T any](capacity int) *Deque[T] {
	return &Deque[T]{
		data:     make([]T, capacity), // 配列の初期化
		capacity: capacity,
	}
}

// 先頭に要素を追加
func (d *Deque[T]) PushFront(value T) {
	if d.IsFull() {
		panic("deque is full")
	}
	// headを逆方向に進めて要素を追加
	d.head = (d.head - 1 + d.capacity) % d.capacity
	d.data[d.head] = value
	d.elementCount++
}

// 末尾に要素を追加
func (d *Deque[T]) PushBack(value T) {
	if d.IsFull() {
		panic("deque is full")
	}
	// 要素を追加し、tailを進める
	d.data[d.tail] = value
	d.tail = (d.tail + 1) % d.capacity
	d.elementCount++
}

// 先頭の要素を削除
func (d *Deque[T]) PopFront() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// 要素を取得し、headを進める
	value := d.data[d.head]
	d.head = (d.head + 1) % d.capacity
	d.elementCount--
	return value
}

// 末尾の要素を削除
func (d *Deque[T]) PopBack() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// tailを逆方向に進めて要素を取得
	d.tail = (d.tail - 1 + d.capacity) % d.capacity
	value := d.data[d.tail]
	d.elementCount--
	return value
}

// 先頭の要素を取得
func (d *Deque[T]) Front() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	return d.data[d.head]
}

// 末尾の要素を取得
func (d *Deque[T]) Back() T {
	if d.IsEmpty() {
		panic("deque is empty")
	}
	// tailの直前の要素を取得
	return d.data[(d.tail-1+d.capacity)%d.capacity]
}

// デックが空かどうか
func (d *Deque[T]) IsEmpty() bool {
	return d.elementCount == 0
}

// デックが満杯かどうか
func (d *Deque[T]) IsFull() bool {
	return d.elementCount == d.capacity
}

// デックの要素数
func (d *Deque[T]) Size() int {
	return d.elementCount
}

func main() {
	// デックの作成
	d := NewDeque[int](5)

	// 末尾に要素を追加
	d.PushBack(10)
	d.PushBack(20)
	// 先頭に要素を追加
	d.PushFront(5)

	// 先頭と末尾の要素を確認
	fmt.Println(d.Front()) // 5
	fmt.Println(d.Back())  // 20

	// 先頭と末尾の要素を削除
	fmt.Println(d.PopFront()) // 5
	fmt.Println(d.PopBack())  // 20

	// 要素数の確認
	fmt.Println("Size:", d.Size()) // 1
}
