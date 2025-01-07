package unionfind

type UnionFind struct {
	parent []int // len(parent)分の要素を考え、各要素の親を記録している
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent, rank}
}

// O(α(N))　※定数時間。α(N)はアッカーマン関数の逆関数
// xの親を見つける（経路圧縮を適用）
func (uf *UnionFind) Find(xIdx int) int {
	if uf.parent[xIdx] != xIdx {
		uf.parent[xIdx] = uf.Find(uf.parent[xIdx])
	}
	return uf.parent[xIdx]
}

// O(α(N))　※定数時間。α(N)はアッカーマン関数の逆関数
// xとyを同じグループに統合する（ランクを考慮）
func (uf *UnionFind) Union(xIdx, yIdx int) {
	rootX := uf.Find(xIdx)
	rootY := uf.Find(yIdx)

	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// O(N)
func (uf *UnionFind) CountRoots() int {
	count := 0
	for i := range uf.parent {
		if uf.parent[i] == i {
			count++
		}
	}
	return count
}
