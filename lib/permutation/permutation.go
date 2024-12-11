package permutation

// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
// optionsには全ての要素が異なるものを渡すこと
func Permute[T any](current []T, options []T) [][]T {
	var results [][]T

	cc := append([]T{}, current...)
	co := append([]T{}, options...)

	if len(co) == 0 {
		return [][]T{cc}
	}

	for i, o := range options {
		newcc := append([]T{}, cc...)
		newcc = append(newcc, o)
		newco := append([]T{}, co[:i]...)
		newco = append(newco, co[i+1:]...)

		subResults := Permute(newcc, newco)
		results = append(results, subResults...)
	}

	return results
}

// 要素数 len(options) で、i番目の要素が options[i] であるような順列のパターンを全列挙する
func Permute2[T any](current []T, options [][]T) [][]T {
	var results [][]T

	if len(current) == len(options) {
		results = append(results, current)
		return results
	}

	nextVals := options[len(current)]
	for _, v := range nextVals {
		copyCurrent := append([]T{}, current...)
		copyCurrent = append(copyCurrent, v)
		subResults := Permute2(copyCurrent, options)
		results = append(results, subResults...)
	}

	return results
}
