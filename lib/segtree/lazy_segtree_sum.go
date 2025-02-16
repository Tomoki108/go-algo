package segtree

// 区間和の遅延セグメント木
// 遅延セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type LazySegTreeSum struct {
	originSize int
	leafSize   int
	data       []int
	lazy       []int // 遅延伝搬用配列
}

// O(N) N: 元々の配列の要素数
func NewLazySegTreeSum(n int) *LazySegTreeSum {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}
	data := make([]int, 2*leafSize)
	lazy := make([]int, 2*leafSize)
	return &LazySegTreeSum{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
		lazy:       lazy,
	}
}

// O(N) N: 元々の配列の要素数
func (seg *LazySegTreeSum) Build(arr []int) {
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.leafSize] = arr[i]
	}
	for i := seg.leafSize - 1; i > 0; i-- {
		seg.data[i] = seg.data[2*i] + seg.data[2*i+1]
	}
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) に対して値 val を加算
func (seg *LazySegTreeSum) Update(originL, originR, val int) {
	seg.updateRec(originL, originR, val, 1, 0, seg.leafSize)
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) の和を取得
func (seg *LazySegTreeSum) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeSum) updateRec(originL, originR, val, currentNode, nl, nr int) {
	// 遅延情報を先に処理
	seg.push(currentNode, nl, nr)
	// 完全に区間外の場合
	if originR <= nl || nr <= originL {
		return
	}
	// 完全に区間内の場合
	if originL <= nl && nr <= originR {
		seg.lazy[currentNode] += val
		seg.push(currentNode, nl, nr)
		return
	}
	// 部分的に区間と重なる場合は子に伝搬
	mid := (nl + nr) / 2
	seg.updateRec(originL, originR, val, 2*currentNode, nl, mid)
	seg.updateRec(originL, originR, val, 2*currentNode+1, mid, nr)
	// 子の値から親の値を再計算
	seg.data[currentNode] = seg.data[2*currentNode] + seg.data[2*currentNode+1]
}

func (seg *LazySegTreeSum) queryRec(l, r, node, nl, nr int) int {
	seg.push(node, nl, nr)
	// 完全に区間外の場合
	if r <= nl || nr <= l {
		return 0 // 単位元
	}
	// 完全に区間内の場合
	if l <= nl && nr <= r {
		return seg.data[node]
	}
	// 部分的に重なる場合は子ノードに問い合わせ
	mid := (nl + nr) / 2
	left := seg.queryRec(l, r, 2*node, nl, mid)
	right := seg.queryRec(l, r, 2*node+1, mid, nr)
	return left + right
}

// 遅延情報を子ノードに伝搬し、現在のノードの値を更新
func (seg *LazySegTreeSum) push(node, nl, nr int) {
	if seg.lazy[node] != 0 {
		// 現在の区間の総和に対して、遅延値を反映
		seg.data[node] += seg.lazy[node] * (nr - nl)
		// 子ノードが存在する場合は、遅延情報を子へ伝搬
		if node < seg.leafSize {
			seg.lazy[2*node] += seg.lazy[node]
			seg.lazy[2*node+1] += seg.lazy[node]
		}
		seg.lazy[node] = 0
	}
}
