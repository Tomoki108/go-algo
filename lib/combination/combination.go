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

// O(nCr) n: len(options), r: n
// optionsから N個選ぶ組み合わせを全列挙する
// optionsにはソート済みかつ要素に重複のないスライスを渡すこと（戻り値が辞書順になり、重複組み合わせも排除される）
func PickN[T comparable](current, options []T, n int) [][]T {
	var results [][]T

	if n == 0 {
		return [][]T{current}
	}

	for i, o := range options {
		newCurrent := make([]T, len(current), len(current)+1)
		copy(newCurrent, current)
		newCurrent = append(newCurrent, o)

		newOptions := make([]T, len(options[i+1:]))
		copy(newOptions, options[i+1:])

		results = append(results, PickN(newCurrent, newOptions, n-1)...)
	}

	return results
}

// スライスを、任意の数の区別のないグループに分けるパターンを全列挙する
// O(B(n, n))
//   - n: len(sl)
//   - B(n, n): ベル数
func Grouping[T any](sl []T) [][][]T {
	var results [][][]T
	groups := make([][]T, 0, len(sl))

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(sl) {
			results = append(results, groups)
			return
		}

		for _, g := range groups {
			g = append(g, sl[idx])
			dfs(idx + 1)
			g = g[:len(g)-1]
		}

		groups = append(groups, []T{sl[idx]})
		dfs(idx + 1)
		groups = groups[:len(groups)-1]
	}

	dfs(0)
	return results
}

// スライスを、区別のあるsize個のグループに分けるパターンを全列挙する
// O(n! * B(n, size))
//   - n: len(sl)
//   - B(n, n): ベル数
func GroupingDistintBySize[T any](sl []T, size int) [][][]T {
	groups := make([][]T, size)

	var results [][][]T

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(sl) {
			for _, g := range groups {
				if len(g) == 0 {
					return
				}
			}

			results = append(results, groups)
			return
		}

		for gi := range groups {
			groups[gi] = append(groups[gi], sl[idx])
			dfs(idx + 1)
			groups[gi] = groups[gi][:len(groups[gi])-1]
		}
	}

	dfs(0)
	return results
}
