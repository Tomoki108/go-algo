package slwindow

import "testing"

func TestExampleSlWindow(t *testing.T) {

	r1 := SlWindowExample([]int{1, 2, 29, 4, 11, 6, 2, 9, 9}, 17)
	if r1 != "found" {
		t.Fatalf("got %v, expect %v", r1, "found")
	}

	r2 := SlWindowExample([]int{1, 2, 29, 4, 11, 6, 2, 9, 9}, 100)
	if r2 != "not found" {
		t.Fatalf("got %v, expect %v", r2, "not found")
	}
}
