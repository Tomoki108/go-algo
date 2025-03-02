package heap

type pqItem struct {
	node    int
	distSum int
}

func (p pqItem) Priority() int {
	return p.distSum
}

// O(E * log V) (E: 辺の数, V: 頂点の数)
// ダイクストラ法で、始点から各頂点への最短距離を求める
func Dijkstra(graph [][][2]int, startNode int) (dists []int) {
	N := len(graph)
	dists = make([]int, N)

	for i := 0; i < N; i++ {
		dists[i] = 1<<63 - 1 // INT_MAX
	}
	fixed := make([]bool, N)

	pq := Heap[pqItem]{}
	pq.PushItem(pqItem{0, 0})
	for len(pq) > 0 {
		item := pq.PopItem()
		if fixed[item.node] {
			continue
		}
		dists[item.node] = item.distSum
		fixed[item.node] = true

		adjacents := graph[item.node]
		for _, adj := range adjacents {
			pq.PushItem(pqItem{adj[0], item.distSum + adj[1]})
		}
	}

	return dists
}
