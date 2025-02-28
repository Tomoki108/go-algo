package fenwicktree

// 数列の区間和の取得、一点更新を O(log n) で行うデータ構造
// できることはセグメント木の完全下位互換だが、定数倍が小さい
type FenwickTree struct {
	n    int
	tree []int
}

// O(n)
// n+1 の長さのフェンウィック木を作成する.
// インターフェースでは 0-indexed で、内部では 1-indexed で扱うため+1.
func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		n:    n,
		tree: make([]int, n+1),
	}
}

// O(log n)
func (ft *FenwickTree) Update(i int, delta int) {
	i++ // 内部は 1-indexed として扱うため
	for i <= ft.n {
		ft.tree[i] += delta
		i += i & -i // 次の更新対象のインデックスへ
	}
}

// O(log n)
// 区間 [0, i] (0-indexed) の和を返す
func (ft *FenwickTree) Sum(i int) int {
	s := 0
	i++ // 内部は 1-indexed として扱うため
	for i > 0 {
		s += ft.tree[i]
		i -= i & -i
	}
	return s
}

// O(log n)
// 区間 [l, r] (0-indexed) の和を返す
func (ft *FenwickTree) RangeSum(l, r int) int {
	if l == 0 {
		return ft.Sum(r)
	}
	return ft.Sum(r) - ft.Sum(l-1)
}
