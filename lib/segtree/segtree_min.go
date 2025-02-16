package segtree

// 区間最小値のセグメント木
// セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type SegTreeMin struct {
	originSize int
	leafSize   int
	data       []int
}

func NewSegTreeMin(n int) *SegTreeMin {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}

	data := make([]int, 2*leafSize)
	return &SegTreeMin{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
	}
}

// O(N) N: 元々の配列の要素数
func (st *SegTreeMin) Build(arr []int) {
	for i := 0; i < len(arr); i++ {
		st.data[i+st.leafSize] = arr[i]
	}
	// 余った葉を単位元で埋める（この場合、minを取る際に影響の出ない十分に大きな数）
	for i := len(arr); i < st.leafSize; i++ {
		st.data[i+st.leafSize] = int(1e18)
	}
	for i := st.leafSize - 1; i > 0; i-- {
		st.data[i] = min(st.data[i*2], st.data[i*2+1])
	}
}

// O(log N) N: 元々の配列の要素数
// idx番目の値をvalueに更新
func (st *SegTreeMin) Update(originIdx, value int) {
	idx := originIdx + st.leafSize
	st.data[idx] = value
	for idx > 0 {
		idx /= 2
		st.data[idx] = min(st.data[idx*2], st.data[idx*2+1])
	}
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) の範囲の最小値を取得
func (st *SegTreeMin) Query(originL, originR int) int {
	return st.queryRec(originL, originR, 1, 0, st.leafSize)
}

func (st *SegTreeMin) queryRec(originL, originR, currentNode, nl, nr int) int {
	if originR <= nl || nr <= originL {
		return int(1e18)
	}
	if originL <= nl && nr <= originR {
		return st.data[currentNode]
	}

	mid := (nl + nr) / 2
	left := st.queryRec(originL, originR, currentNode*2, nl, mid)
	right := st.queryRec(originL, originR, currentNode*2+1, mid, nr)
	return min(left, right)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
