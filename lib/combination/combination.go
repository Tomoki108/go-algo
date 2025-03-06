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

// NOTE: 呼び出し側に結果を反映するため、callback内でansポインタの値などを更新すること
//
// O(|options| C n)
// optionsからn個選ぶ組み合わせ全てに対して、callbackを呼び出す.
// idx: 現在考慮するoptionsのidx.
// options: 選択肢.
// current: 現在選ばれている要素.
// n: 選ぶ個数.
// callback: 組み合わせが揃った時に呼び出される関数
func AllCombination[T any](idx int, options []T, current []T, n int, callback func([]T)) {
	if len(current) == n {
		callback(current)
		return
	}
	if len(options)-idx < n-len(current) {
		return
	}

	// 選ぶ場合
	current = append(current, options[idx])
	AllCombination(idx+1, options, current, n, callback)
	current = current[:len(current)-1]

	// 選ばない場合
	AllCombination(idx+1, options, current, n, callback)
}

// NOTE: スライスのcopyが多く、n = 10 程度で致命的に遅い.AllCombination()を推奨.
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
