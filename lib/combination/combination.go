package combination

// nCrの計算 O(r)
// (n * (n-1) ... * (n-r+1)) / r!
func CombinationNum(n, r int) int {
	if r > n {
		return 0
	}
	if r > n/2 {
		r = n - r // Use smaller r for efficiency
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// optionsから N個選ぶ組み合わせを全列挙する
// optionsにはソート済みかつ要素に重複のないスライスを渡すこと（戻り値が辞書順になり、重複組み合わせも排除される）
// O(nCr)
func PickN[T comparable](current, options []T, n int) [][]T {
	var results [][]T

	if n == 0 {
		return [][]T{current}
	}

	for i, o := range options {
		newCurrent := append([]T{}, current...)
		newCurrent = append(newCurrent, o)
		newOptions := append([]T{}, options[i+1:]...)

		results = append(results, PickN(newCurrent, newOptions, n-1)...)
	}

	return results
}
