package segtree

// 区間最大値のセグメント木
// セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type SegTreeMax struct {
	n    int
	size int
	tree []int
}

func NewSegTreeMax(arr []int) *SegTreeMax {
	n := len(arr)
	size := 1
	for size < n {
		size *= 2
	}
	tree := make([]int, 2*size)
	minVal := -1 << 63
	for i := 0; i < 2*size; i++ {
		tree[i] = minVal
	}
	for i, v := range arr {
		tree[size+i] = v
	}
	for i := size - 1; i > 0; i-- {
		tree[i] = max(tree[2*i], tree[2*i+1])
	}
	return &SegTreeMax{
		n:    n,
		size: size,
		tree: tree,
	}
}

// O(log N) N: 元々の配列の要素数
// idx番目の値をvalueに更新
func (st *SegTreeMax) Update(i, val int) {
	i += st.size
	st.tree[i] = val
	for i > 1 {
		i /= 2
		st.tree[i] = max(st.tree[2*i], st.tree[2*i+1])
	}
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) の範囲の最大値を取得
func (st *SegTreeMax) Query(l, r int) int {
	minVal := -1 << 63
	res := minVal
	l += st.size
	r += st.size
	for l < r {
		if l%2 == 1 {
			res = max(res, st.tree[l])
			l++
		}
		if r%2 == 1 {
			r--
			res = max(res, st.tree[r])
		}
		l /= 2
		r /= 2
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
