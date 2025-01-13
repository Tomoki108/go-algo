package sort

// https://zenn.dev/satumahayato010/articles/58fb473c40d9af#merge-sort
func MergeSort(nums []int) []int {
	var lenNums = len(nums)
	if lenNums == 1 {
		return nums
	}

	mid := lenNums / 2
	var (
		left  = make([]int, mid)
		right = make([]int, lenNums-mid)
	)
	for i := 0; i < mid; i++ {
		left[i] = nums[i]
	}
	for i := mid; i < lenNums; i++ {
		right[i-mid] = nums[i]
	}

	return merge(MergeSort(left), MergeSort(right))
}

// O(len(left) + len(right))
// ソート済みの2つのスライスをマージしてソートする
func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// append the remaining elements to result
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
