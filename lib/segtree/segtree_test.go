package segtree

import "testing"

func TestSegTreeSum(t *testing.T) {
	st := NewSegTree([]int{1, 2, 3, 4, 5}, func(a, b int) int { return a + b }, 0)

	if got, want := st.Query(0, 3), 6; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(2, 10)
	if got, want := st.Query(0, 3), 13; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSegTreeMin(t *testing.T) {
	st := NewSegTree([]int{1, 2, 3, 4, 5}, func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}, 1<<63-1)

	if got, want := st.Query(0, 3), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(2, 10)
	if got, want := st.Query(0, 3), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
