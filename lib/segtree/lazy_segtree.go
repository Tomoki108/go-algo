package segtree

// LazySegmentTree は区間加算と区間和取得を行う遅延セグメント木の例です。
type LazySegmentTree struct {
	n, size int
	data    []int // 各区間の和を保持（ノードの値）
	lazy    []int // 遅延伝搬用配列
}

// NewLazySegmentTree は要素数 n の遅延セグメント木を初期化します。
// 遅延セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
func NewLazySegmentTree(n int) *LazySegmentTree {
	size := 1
	for size < n {
		size *= 2
	}
	// 2*size は完全2分木のノード数の上限
	data := make([]int, 2*size)
	lazy := make([]int, 2*size)
	return &LazySegmentTree{
		n:    n,
		size: size,
		data: data,
		lazy: lazy,
	}
}

// build は元の配列 arr からセグメント木を構築します。
func (seg *LazySegmentTree) Build(arr []int) {
	// 葉ノードに値を設定
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.size] = arr[i]
	}
	// 足りない葉は単位元（ここでは0）で初期化
	for i := len(arr); i < seg.size; i++ {
		seg.data[i+seg.size] = 0
	}
	// 内部ノードの値を構築（ここでは和）
	for i := seg.size - 1; i > 0; i-- {
		seg.data[i] = seg.data[2*i] + seg.data[2*i+1]
	}
}

// Update は区間 [l, r) に対して値 val を加算する外部インターフェースです。
func (seg *LazySegmentTree) Update(l, r, val int) {
	seg.updateRec(l, r, val, 1, 0, seg.size)
}

// Query は区間 [l, r) の和を取得する外部インターフェースです。
func (seg *LazySegmentTree) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.size)
}

// updateRec は区間 [l, r) に対して値 val を加算します。（再帰処理）
func (seg *LazySegmentTree) updateRec(l, r, val, node, nl, nr int) {
	// 遅延情報を先に処理
	seg.push(node, nl, nr)
	// 完全に区間外の場合
	if r <= nl || nr <= l {
		return
	}
	// 完全に区間内の場合
	if l <= nl && nr <= r {
		seg.lazy[node] += val
		seg.push(node, nl, nr)
		return
	}
	// 部分的に区間と重なる場合は子に伝搬
	mid := (nl + nr) / 2
	seg.updateRec(l, r, val, 2*node, nl, mid)
	seg.updateRec(l, r, val, 2*node+1, mid, nr)
	// 子の値から親の値を再計算
	seg.data[node] = seg.data[2*node] + seg.data[2*node+1]
}

// queryRec は区間 [l, r) の和を取得する再帰処理です。
func (seg *LazySegmentTree) queryRec(l, r, node, nl, nr int) int {
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

// push は遅延情報を子ノードに伝搬し、現在のノードの値を更新します。
func (seg *LazySegmentTree) push(node, nl, nr int) {
	if seg.lazy[node] != 0 {
		// 現在の区間の総和に対して、遅延値を反映
		seg.data[node] += seg.lazy[node] * (nr - nl)
		// 子ノードが存在する場合は、遅延情報を子へ伝搬
		if node < seg.size {
			seg.lazy[2*node] += seg.lazy[node]
			seg.lazy[2*node+1] += seg.lazy[node]
		}
		seg.lazy[node] = 0
	}
}
