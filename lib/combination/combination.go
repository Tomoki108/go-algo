package combination

// O(r)
// nCrの計算
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

// NOTE: スライスのcopyが多く、n = 10 程度で致命的に遅い. DFSで組み合わせに対して逐次処理をすることを推奨.
//
// O(nCr) n: len(options), r: n
// optionsから N個選ぶ組み合わせを全列挙する
// optionsにはソート済みのスライスを渡すこと（戻り値が辞書順になる）
// 同じ値を区別しない場合は、要素に重複のないスライスを渡すこと（重複組み合わせが排除される）
func PickN[T comparable](current, options []T, n int) [][]T {
	var results [][]T

	if n == 0 {
		ccurrent := make([]T, len(current))
		copy(ccurrent, current)
		return [][]T{ccurrent}
	}

	for i, o := range options {
		current = append(current, o)
		newOptions := make([]T, len(options[i+1:]))
		copy(newOptions, options[i+1:])

		results = append(results, PickN(current, newOptions, n-1)...)
		current = current[:len(current)-1]
	}

	return results
}
