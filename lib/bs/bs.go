package bs

// low, low+1, ..., highの範囲で条件を満たす最小の値を二分探索する
// low~highは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はlow-1を返す
func IntSearch(low, high int, f func(num int) bool) int {
	for low < high {
		mid := low + (high-low)/2 // オーバーフローを防ぐための式
		if f(mid) {
			high = mid // 条件を満たす可能性がある範囲を縮める
		} else {
			low = mid + 1 // 条件を満たさない場合、次の範囲を探索
		}
	}

	// 最後に low が条件を満たしているかを確認
	if f(low) {
		return low
	}

	return low - 1 // 条件を満たす値が見つからない場合
}

// high, high-1, ..., lowの範囲で条件を満たす最小の値を二分探索する
// high~lowは条件に対して単調減少性を満たす必要がある
// 条件を満たす値が見つからない場合はhigh+1を返す
func RevIntSearch(low, high int, f func(num int) bool) int {
	for low < high {
		mid := low + (high-low+1)/2 // オーバーフローを防ぐための式. 高い方を切り捨てるため +1
		if f(mid) {
			low = mid // 条件を満たす可能性がある範囲を縮める
		} else {
			high = mid - 1 // 条件を満たさない場合、次の範囲を探索
		}
	}

	// 最後に high が条件を満たしているかを確認
	if f(high) {
		return high
	}

	return high + 1 // 条件を満たす値が見つからない場合
}
