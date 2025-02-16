package segtree

// 区間和のセグメント木
// セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94#%E3%81%9D%E3%82%82%E3%81%9D%E3%82%82%E3%82%BB%E3%82%B0%E6%9C%A8%E3%81%A8%E3%81%AF
type SegTreeSum struct {
	originSize int
	leafSize   int
	data       []int
}

// O(N)
func NewSegTreeSum(n int) *SegTreeSum {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}
	// 葉のサイズはN以上の最小の2のべき乗
	// 葉のサイズの計算はO(log N)であるが、計算量ではdataの初期化のO(N)の方が支配的

	// 総ノード数は leafSize + leafSize/2 + leafSize/4 ... = 2*leafSize-1
	// 内部処理では1-indexedで扱うため、2*leafSizeを確保
	data := make([]int, 2*leafSize)
	return &SegTreeSum{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
	}
}

// O(N)
func (st *SegTreeSum) Build(arr []int) {
	// 葉ノードに値を設定
	for i := 0; i < len(arr); i++ {
		st.data[i+st.leafSize] = arr[i]
	}
	// 上位ノードの値を下から構築
	for i := st.leafSize - 1; i > 0; i-- { // leafSizeは2Nなので、計算量には影響しない
		st.data[i] = st.data[i*2] + st.data[i*2+1]
	}
}

// O(log N)
// idx番目の値をvalueに更新
func (st *SegTreeSum) Update(originIdx, value int) {
	idx := originIdx + st.leafSize
	st.data[idx] = value
	for idx > 0 {
		idx /= 2
		st.data[idx] = st.data[idx*2] + st.data[idx*2+1]
	}
}

// O(log N)
// [originL, originR) の範囲の和を取得
func (st *SegTreeSum) Query(originL, originR int) int {
	return st.queryRec(originL, originR, 1, 0, st.leafSize)
}

func (st *SegTreeSum) queryRec(originL, originR, currentNode, nl, nr int) int {
	// [originL, originR) と [nl, nr) が交差しない場合
	if originR <= nl || nr <= originL {
		return 0
	}
	// [originL, originR) が [nl, nr) を完全に覆っている場合
	if originL <= nl && nr <= originR {
		return st.data[currentNode]
	}
	// 一部だけ範囲がかぶる場合
	mid := (nl + nr) / 2
	vl := st.queryRec(originL, originR, currentNode*2, nl, mid)
	vr := st.queryRec(originL, originR, currentNode*2+1, mid, nr)
	return vl + vr
}
