package diffarr

// O(n)
// intスライスの差分配列を返す
func DiffArray(sl []int) []int {
	res := make([]int, 0, len(sl))
	res = append(res, sl[0])
	for i := 1; i < len(sl); i++ {
		res = append(res, sl[i]-sl[i-1])
	}
	return res
}

// O(1)
// 差分配列への区間更新を行う。[l, r)にxを加算する。
// 更新後に累積和をとっていくと、各インデックスの値が求まる。所謂imos法
func RangeUpdateDiffArray(sl []int, l, r, x int) {
	if l < len(sl) {
		sl[l] += x
	}
	if r < len(sl) {
		sl[r] -= x
	}
}
