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

func TestSegTreeMin(t *testing.T) {
	st := NewSegTreeMin(5)
	st.Build([]int{1, 2, 3, 4, 5})

	if got, want := st.Query(0, 3), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(2, 10)
	if got, want := st.Query(0, 3), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLazySegTreeSum(t *testing.T) {
	st := NewLazySegTreeSum(5)
	st.Build([]int{1, 2, 3, 4, 5})

	if got, want := st.Query(0, 3), 6; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(0, 3, 10)
	if got, want := st.Query(0, 3), 36; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLazySegTreeMin(t *testing.T) {
	st := NewLazySegTreeMin(5)
	st.Build([]int{1, 2, 3, 4, 5})

	if got, want := st.Query(0, 3), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	st.Update(0, 3, 10)
	if got, want := st.Query(0, 3), 11; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
