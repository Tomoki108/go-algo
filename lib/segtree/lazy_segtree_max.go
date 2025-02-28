package segtree

// 区間最大値の遅延セグメント木
// 遅延セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type LazySegTreeMax struct {
	originSize int
	leafSize   int
	data       []int
	lazy       []int // 遅延伝搬用配列
}

// O(N) N: 元々の配列の要素数
func NewLazySegTreeMax(n int) *LazySegTreeMax {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}
	data := make([]int, 2*leafSize)
	lazy := make([]int, 2*leafSize)
	return &LazySegTreeMax{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
		lazy:       lazy,
	}
}

// O(N) N: 元々の配列の要素数
func (seg *LazySegTreeMax) Build(arr []int) {
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.leafSize] = arr[i]
	}
	// arrの要素がない部分は単位元（0）で初期化
	for i := len(arr); i < seg.leafSize; i++ {
		seg.data[i+seg.leafSize] = 0
	}
	// 上位ノードの値を下から構築
	for i := seg.leafSize - 1; i > 0; i-- {
		seg.data[i] = max(seg.data[2*i], seg.data[2*i+1])
	}
}

func (seg *LazySegTreeMax) Update(originL, originR, val int) {
	seg.updateRec(originL, originR, val, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeMax) updateRec(originL, originR, val, currentNode, nl, nr int) {
	// まず遅延情報を反映
	seg.push(currentNode)
	// 現在の区間 [nl, nr) がクエリ区間 [originL, originR) と全く重ならない場合
	if originR <= nl || nr <= originL {
		return
	}
	// 現在の区間 [nl, nr) がクエリ区間 [originL, originR) に完全に含まれる場合
	if originL <= nl && nr <= originR {
		seg.lazy[currentNode] += val
		seg.push(currentNode)
		return
	}
	// 一部だけ重なる場合
	seg.updateRec(originL, originR, val, 2*currentNode, nl, (nl+nr)/2)
	seg.updateRec(originL, originR, val, 2*currentNode+1, (nl+nr)/2, nr)
	seg.data[currentNode] = max(seg.data[2*currentNode]+seg.lazy[2*currentNode], seg.data[2*currentNode+1]+seg.lazy[2*currentNode+1])
}

func (seg *LazySegTreeMax) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeMax) queryRec(originL, originR, currentNode, nl, nr int) int {
	// 遅延情報を反映
	seg.push(currentNode)
	// 完全に区間外の場合
	if originR <= nl || nr <= originL {
		return 0
	}
	// 完全に区間内の場合
	if originL <= nl && nr <= originR {
		return seg.data[currentNode] + seg.lazy[currentNode]
	}
	// 部分的に区間と重なる場合
	lRes := seg.queryRec(originL, originR, 2*currentNode, nl, (nl+nr)/2)
	rRes := seg.queryRec(originL, originR, 2*currentNode+1, (nl+nr)/2, nr)
	return max(lRes, rRes)
}

func (seg *LazySegTreeMax) push(node int) {
	if node < seg.leafSize {
		seg.lazy[2*node] += seg.lazy[node]
		seg.lazy[2*node+1] += seg.lazy[node]
	}
	seg.data[node] += seg.lazy[node]
	seg.lazy[node] = 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
