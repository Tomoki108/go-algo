package heap

import "testing"

func Test_IntHeap(t *testing.T) {
	h := NewIntHeap(MinIntHeap)
	h.Push(3)
	h.Push(2)
	h.Push(1)

	if h.Pop() != 1 {
		t.Errorf("Pop() should return 1")
	}
	if h.Pop() != 2 {
		t.Errorf("Pop() should return 2")
	}
	if h.Pop() != 3 {
		t.Errorf("Pop() should return 3")
	}
}
