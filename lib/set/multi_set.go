package set

import (
	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
)

func NewIntMultiSetAsc() *MultiSet[int] {
	return NewMultiSet[int](comparator.IntComparator)
}

func NewIntMultiSetDesc() *MultiSet[int] {
	return NewMultiSet[int](comparator.Reverse(comparator.IntComparator))
}

// gostlのNativeのMultiSetは、Erace()が同一の値を全て削除してしまうためこちらを使う。
type MultiSet[T any] struct {
	tree *rbtree.RbTree[T, int]
}

// O(1)
func NewMultiSet[T any](compare comparator.Comparator[T]) *MultiSet[T] {
	return &MultiSet[T]{
		tree: rbtree.New[T, int](compare),
	}
}

// O(log n)
func (ms *MultiSet[T]) Insert(value T) {
	if node := ms.tree.FindNode(value); node != nil {
		node.SetValue(node.Value() + 1)
	} else {
		ms.tree.Insert(value, 1)
	}
}

// Erase は、要素を1つだけ削除します。（カウントが1の場合はキーごと削除）
// O(log n)
func (ms *MultiSet[T]) Erase(value T) {
	if node := ms.tree.FindNode(value); node != nil {
		count := node.Value()
		if count > 1 {
			node.SetValue(count - 1)
		} else {
			ms.tree.Delete(node)
			return
		}
	}
}

// Count は、指定した要素の出現回数を返します。
// O(log n)
func (ms *MultiSet[T]) Count(value T) int {
	if node := ms.tree.FindNode(value); node != nil {
		return node.Value()
	}
	return 0
}

// Values は、マルチセット内の全ての要素をソートされた順序（比較関数に準ずる）で重複も含めてスライスとして返します。
// O(n + k)
//   - n はユニークなキーの数
//   - k は要素の総数（重複含む）
func (ms *MultiSet[T]) Values() []T {
	var result []T
	it := ms.tree.IterFirst()
	for it.IsValid() {
		key, count := it.Key(), it.Value()
		for i := 0; i < count; i++ {
			result = append(result, key)
		}

		it = it.Next().(*rbtree.RbTreeIterator[T, int])
	}
	return result
}

// Size は、マルチセットに含まれる要素の総数（重複含む）を返します。
// O(n)
//   - n はユニークなキーの数（RBTree 全体を走査）
func (ms *MultiSet[T]) Size() int {
	total := 0
	it := ms.tree.IterFirst()
	for it.IsValid() {
		_, count := it.Key(), it.Value()
		total += count

		it = it.Next().(*rbtree.RbTreeIterator[T, int])
	}
	return total
}

// Clear は、マルチセットを空にします。
// O(1) （実装上、要素数に依存しないクリア処理を行うため）
func (ms *MultiSet[T]) Clear() {
	ms.tree.Clear()
}

func (ms *MultiSet[T]) First() *rbtree.RbTreeIterator[T, int] {
	return ms.tree.IterFirst()
}

func (ms *MultiSet[T]) Last() *rbtree.RbTreeIterator[T, int] {
	return ms.tree.IterLast()
}
