package topolosort

// O(V + E) (V: 頂点の数, E: 辺の数)
// トポロジカルソートを行う
// graphにはDAG（有向非巡回グラフ）を渡すこと
func TopologicalSort(graph [][]int) []int {
	N := len(graph)
	visited := make([]bool, N)
	ret := make([]int, 0, N)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, adj := range graph[node] {
			if visited[adj] {
				continue
			}
			dfs(adj)
		}
		ret = append(ret, node)
	}

	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	return ret
}

// O(V + E) (V: 頂点の数, E: 辺の数)
// startNodeから始まるトポロジカルソートを行う
// startNodeのタスクの前提となるタスクを順番に列挙する. 戻り値の末尾はstartNode自身
// graphにはDAG（有向非巡回グラフ）を渡すこと
func TopologicalSortFrom(graph [][]int, startNode int) []int {
	N := len(graph)

	visited := make([]bool, N)
	ret := make([]int, 0, N)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, adj := range graph[node] {
			if visited[adj] {
				continue
			}
			dfs(adj)
		}
		ret = append(ret, node)
	}
	dfs(startNode)

	return ret[:len(ret)-1] // remove startNode
}
