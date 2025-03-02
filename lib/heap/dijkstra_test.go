package heap

import "testing"

func TestDijkstra(t *testing.T) {
	graph := [][][2]int{
		{{1, 1}, {2, 4}},
		{{2, 2}},
		{{3, 1}},
		{{0, 7}},
	}
	dists := Dijkstra(graph, 0)
	if dists[0] != 0 {
		t.Fatalf("got %v, expect %v", dists[0], 0)
	}
	if dists[1] != 1 {
		t.Fatalf("got %v, expect %v", dists[1], 1)
	}
	if dists[2] != 3 {
		t.Fatalf("got %v, expect %v", dists[2], 3)
	}
	if dists[3] != 4 {
		t.Fatalf("got %v, expect %v", dists[3], 4)
	}

}
