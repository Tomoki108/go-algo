package bs

import "sort"

// O(log (high-low))
// low, low+1, ..., highの範囲で条件を満たす最小の値を二分探索する
// low~highは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はlow-1を返す
func AscIntSearch(low, high int, f func(num int) bool) int {
	initialLow := low

	for low < high {
		// オーバーフローを防ぐための立式
		// 中央値はlow側に寄る
		mid := low + (high-low)/2
		if f(mid) {
			high = mid // 条件を満たす場合、よりlow側の範囲を探索
		} else {
			low = mid + 1 // 条件を満たさない場合、よりhigh側の範囲を探索
		}
	}

	// 最後に low(=high) が条件を満たしているかを確認
	if f(low) {
		return low
	}

	return initialLow - 1 // 条件を満たす値が見つからない場合
}

// O(log (high-low))
// high, high-1, ..., lowの範囲で条件を満たす最大の値を二分探索する
// high~lowは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はhigh+1を返す
func DescIntSearch(high, low int, f func(num int) bool) int {
	for low < high {
		// オーバーフローを防ぐための式.
		// 中央値はhigh側に寄る（+1しているため）
		mid := low + (high-low+1)/2
		if f(mid) {
			low = mid // 条件を満たす場合、よりhigh側の範囲を探索
		} else {
			high = mid - 1 // 条件を満たさない場合、よりlow側の範囲を探索
		}
	}

	// 最後に high(=low) が条件を満たしているかを確認
	if f(high) {
		return high
	}

	return high + 1 // 条件を満たす値が見つからない場合
}

// O(log |sl|)
// 単調増加性を満たすintスライスから、値がl以上、r以下の範囲を返す
func RangeIntSearch(sl []int, l, r int) []int {
	idx1 := sort.Search(len(sl), func(i int) bool { return sl[i] >= l })
	if idx1 == len(sl) {
		return nil
	}
	idx2 := sort.Search(len(sl), func(i int) bool { return sl[i] > r })
	if idx2 == 0 {
		return nil
	}
	return sl[idx1:idx2]
}

// O(log |sl|)
// 単調増加性を満た任意型のスライスから、任意型が持つ単調増加性を満たす値がl以上、r以下の範囲を返す
func RangeSearch[U any](sl []U, valuer func(item U) int, l, r int) []U {
	idx1 := sort.Search(len(sl), func(i int) bool { return valuer(sl[i]) >= l })
	if idx1 == len(sl) {
		return nil
	}
	idx2 := sort.Search(len(sl), func(i int) bool { val := valuer(sl[i]); return val > r })
	if idx2 == 0 {
		return nil
	}
	return sl[idx1:idx2]
}
