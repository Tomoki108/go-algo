package slwindow

// O(len(sl))
// 尺取法の実装例 (連続する部分列の和が目標値以上になる最小の長さを求める)
func SlWindowExample(sl []int, targetSum int) string {
	// (ex,
	// sl := []int{1, 2, 29, 4, 11, 6, 2, 9, 9}
	// targetSum := 17

	/// (lefgt, right] の範囲の和を考える
	left := 0
	right := 0
	currentSum := 0
	for right < len(sl) {
		currentSum += sl[right]
		right++

		if currentSum < targetSum {
			continue
		}

		if currentSum == targetSum {
			return "found"
		}

		for currentSum > targetSum {
			currentSum -= sl[left]
			left++

			if left == right {
				currentSum = 0
				break
			}

			if currentSum == targetSum {
				return "found"
			}
		}
	}

	return "not found"
}
