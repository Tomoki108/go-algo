package unionfind

type UnionFind struct {
	parent []int // len(parent)分のノードを考え、各ノードの親を記録している
	size   []int // そのノードを頂点とする部分木の頂点数
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	s := make([]int, size)
	for i := range parent {
		parent[i] = i
		s[i] = 1
	}
	return &UnionFind{parent, s}
}

// O(α(N))　※定数時間。α(N)はアッカーマン関数の逆関数
// xの親を見つける
func (uf *UnionFind) Find(xIdx int) int {
	if uf.parent[xIdx] != xIdx {
		uf.parent[xIdx] = uf.Find(uf.parent[xIdx]) // 経路圧縮
	}
	return uf.parent[xIdx]
}

// O(α(N))
// xとyを同じグループに統合する（サイズが大きい方に統合）
func (uf *UnionFind) Union(xIdx, yIdx int) {
	rootX := uf.Find(xIdx)
	rootY := uf.Find(yIdx)

	if rootX != rootY {
		if uf.size[rootX] < uf.size[rootY] {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		} else if uf.size[rootX] > uf.size[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		} else {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		}
	}
}

// O(1)
func (uf *UnionFind) IsRoot(xIdx int) bool {
	return uf.parent[xIdx] == xIdx
}

// O(α(N))
func (uf *UnionFind) IsSameRoot(xIdx, yIdx int) bool {
	return uf.Find(xIdx) == uf.Find(yIdx)
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

// O(N)
func (uf *UnionFind) Roots() []int {
	roots := make([]int, 0)
	for i := range uf.parent {
		if uf.parent[i] == i {
			roots = append(roots, i)
		}
	}
	return roots
}

// O(α(N))
func (uf *UnionFind) GroupSize(xIdx int) int {
	return uf.size[uf.Find(xIdx)]
}
