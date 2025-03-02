package slwindow

import "testing"

func TestSlWindowSum(t *testing.T) {

	r1 := SlWindowSum([]int{4, 4, 4, 1, 5, 10, 2}, 17)
	if r1 != 3 {
		t.Fatalf("got %v, expect %v", r1, 3)
	}

	r2 := SlWindowSum([]int{1, 2, 29, 4, 11, 6, 2, 9, 9}, 100)
	if r2 != -1 {
		t.Fatalf("got %v, expect %v", r2, -1)
	}
}
