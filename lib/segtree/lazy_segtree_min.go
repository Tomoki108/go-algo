package segtree

// INF は単位元として使う十分に大きな数
const INF int = 1 << 60

// 区間最小値の遅延セグメント木
// 遅延セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type LazySegTreeMin struct {
	originSize int
	leafSize   int
	data       []int
	lazy       []int // 遅延伝搬用配列
}

// O(N) N: 元々の配列の要素数
func NewLazySegTreeMin(n int) *LazySegTreeMin {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}
	data := make([]int, 2*leafSize)
	lazy := make([]int, 2*leafSize)
	// dataの全要素を単位元（INF）で初期化
	for i := 0; i < len(data); i++ {
		data[i] = INF
	}
	return &LazySegTreeMin{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
		lazy:       lazy,
	}
}

// O(N) N: 元々の配列の要素数
func (seg *LazySegTreeMin) Build(arr []int) {
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.leafSize] = arr[i]
	}
	// arrの要素がない部分は単位元INF
	for i := len(arr); i < seg.leafSize; i++ {
		seg.data[i+seg.leafSize] = INF
	}
	// 上位ノードの値を下から構築
	for i := seg.leafSize - 1; i > 0; i-- {
		seg.data[i] = min(seg.data[2*i], seg.data[2*i+1])
	}
}

// O(log N)
// [originL, originR)に対して、値 val を加算
func (seg *LazySegTreeMin) Update(originL, originR, val int) {
	seg.updateRec(originL, originR, val, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeMin) updateRec(originL, originR, val, currentNode, nl, nr int) {
	// まず遅延情報を反映
	seg.push(currentNode)
	// 現在の区間 [nl, nr) がクエリ区間 [originL, originR) と全く重ならない場合
	if originR <= nl || nr <= originL {
		return
	}
	// 現在の区間がクエリ区間に完全に含まれる場合
	if originL <= nl && nr <= originR {
		seg.lazy[currentNode] += val
		seg.push(currentNode)
		return
	}
	// 部分的に重なる場合は子ノードへ伝搬
	mid := (nl + nr) / 2
	seg.updateRec(originL, originR, val, 2*currentNode, nl, mid)
	seg.updateRec(originL, originR, val, 2*currentNode+1, mid, nr)
	// 子ノードの値から現在のノードの値を再計算（最小値を取る）
	seg.data[currentNode] = min(seg.data[2*currentNode], seg.data[2*currentNode+1])
}

// O(log N)
// [l, r) の最小値を取得
func (seg *LazySegTreeMin) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeMin) queryRec(l, r, node, nl, nr int) int {
	seg.push(node)
	// 現在の区間 [nl, nr) がクエリ区間 [l, r) と重ならない場合
	if r <= nl || nr <= l {
		return INF // 単位元
	}
	// 現在の区間がクエリ区間に完全に含まれる場合
	if l <= nl && nr <= r {
		return seg.data[node]
	}
	// 部分的に重なる場合は子ノードから結果を取得
	mid := (nl + nr) / 2
	left := seg.queryRec(l, r, 2*node, nl, mid)
	right := seg.queryRec(l, r, 2*node+1, mid, nr)
	return min(left, right)
}

// 遅延情報を子ノードに伝搬し、現在のノードの値に反映する
func (seg *LazySegTreeMin) push(node int) {
	if seg.lazy[node] != 0 {
		// 区間全体に加算していた遅延値を反映
		seg.data[node] += seg.lazy[node] // ある区間の最小値の確定は、暫定の値プラス遅延評価の値
		// 葉でなければ子ノードに遅延情報を伝搬
		if node < seg.leafSize {
			seg.lazy[2*node] += seg.lazy[node]
			seg.lazy[2*node+1] += seg.lazy[node]
		}
		seg.lazy[node] = 0
	}
}
