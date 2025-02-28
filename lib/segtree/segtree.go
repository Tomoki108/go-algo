package segtree

const IDENTITY_MIN_TREE = 1<<63 - 1
const IDENTITY_MAX_TREE = -1 << 63
const IDENTITY_SUM_TREE = 0

// 区間和や区間の最小値などの取得、一点更新、区間更新を高速に行うデータ構造
// セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type SegTree struct {
	n        int                // 元々の配列の要素数
	size     int                // 葉のサイズ
	tree     []int              // 内部で保持する木構造
	identity int                // 単位元
	combine  func(int, int) int // 結合関数（ex, 加算、最小値取得、最大値取得）
}

func NewSegTree(arr []int, combine func(int, int) int, identity int) *SegTree {
	n := len(arr)
	size := 1
	for size < n {
		size *= 2
	}
	tree := make([]int, 2*size)
	// 全体を単位元で初期化
	for i := 0; i < 2*size; i++ {
		tree[i] = identity
	}
	// 葉ノードの初期化
	for i, v := range arr {
		tree[size+i] = v
	}
	// 内部ノードの構築
	for i := size - 1; i > 0; i-- {
		tree[i] = combine(tree[2*i], tree[2*i+1])
	}
	return &SegTree{
		n:        n,
		size:     size,
		tree:     tree,
		identity: identity,
		combine:  combine,
	}
}

// O(log N)
// idx の値を value に更新
func (st *SegTree) Update(idx, value int) {
	i := idx + st.size
	st.tree[i] = value
	for i > 1 {
		i /= 2
		st.tree[i] = st.combine(st.tree[2*i], st.tree[2*i+1])
	}
}

// O(log N)
// [l, r) の範囲の結合結果を取得
func (st *SegTree) Query(l, r int) int {
	res := st.identity
	l += st.size
	r += st.size
	for l < r {
		if l%2 == 1 {
			res = st.combine(res, st.tree[l])
			l++
		}
		if r%2 == 1 {
			r--
			res = st.combine(res, st.tree[r])
		}
		l /= 2
		r /= 2
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
