package slwindow

// O(|sl|)
// 尺取法で、連続する部分列の和が目標値以上になる最小の長さを求める
func SlWindowSum(sl []int, targetSum int) int {
	minLen := 1<<63 - 1

	// [left, right)
	left := 0
	right := 0
	currentSum := 0
	for right < len(sl) {
		currentSum += sl[right]
		right++

		for currentSum >= targetSum {
			minLen = min(minLen, right-left)

			currentSum -= sl[left]
			left++
		}
	}

	if minLen != 1<<63-1 {
		return minLen
	}
	return -1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
