package deque

import (
	"testing"
)

// Test Resultタブでログが見られ、デック内部のデータの挙動を確認できる
func Test_Deque(t *testing.T) {
	d := NewDeque[int](4)
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	t.Log("data:", d.data, "head:", d.head, "tail:", d.tail)

	if d.PopFront() != 1 {
		t.Fatal("failed test")
	}
	t.Log("data:", d.data, "head:", d.head, "tail:", d.tail)

	d.PushFront(100)
	if d.Front() != 100 {
		t.Fatal("failed test")
	}
	t.Log("data:", d.data, "head:", d.head, "tail:", d.tail)

	d.PushFront(200)
	if d.Front() != 200 {
		t.Fatal("failed test")
	}
	t.Log("data:", d.data, "head:", d.head, "tail:", d.tail)
}
