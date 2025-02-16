package segtree

import "testing"

func TestSegTreeSum(t *testing.T) {
	st := NewSegTreeSum(5)
	st.Build([]int{1, 2, 3, 4, 5})

	if got, want := st.Query(0, 3), 6; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(2, 10)
	if got, want := st.Query(0, 3), 13; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
