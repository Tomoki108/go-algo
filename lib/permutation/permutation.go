package permutation

import "fmt"

// O(len(sl)*len(sl)!)
// sl の要素を並び替えて、次の辞書順の順列にする
func NextPermutation[T ~int | ~string](sl []T) bool {
	n := len(sl)
	i := n - 2

	// Step1: 右から左に探索して、「スイッチポイント」を見つける:
	// 　「スイッチポイント」とは、右から見て初めて「リストの値が減少する場所」です。
	// 　例: [1, 2, 3, 6, 5, 4] の場合、3 がスイッチポイント。
	for i >= 0 && sl[i] >= sl[i+1] {
		i--
	}

	//　スイッチポイントが見つからない場合、最後の順列に到達しています。
	if i < 0 {
		return false
	}

	// Step2: スイッチポイントの右側の要素から、スイッチポイントより少しだけ大きい値を見つけ、交換します。
	// 　例: 3 を右側で最小の大きい値 4 と交換。
	j := n - 1
	for sl[j] <= sl[i] {
		j--
	}
	sl[i], sl[j] = sl[j], sl[i]

	// Step3: スイッチポイントの右側を反転して、辞書順に次の順列を作ります。
	// 　例: [1, 2, 4, 6, 5, 3] → [1, 2, 4, 3, 5, 6]。
	reverse(sl[i+1:])
	return true
}

func reverse[T ~int | ~string](sl []T) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

// NOTE: スライスのcopyが多く、n = 10 程度で致命的に遅い。NetxPermutationを推奨。
//
// O(n!) n: len(options)
// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
// options[i]に重複した要素が含まれていても、あらかじめソートしておけば重複パターンは除かれる
func Permute[T comparable](current []T, options []T) [][]T {
	fmt.Printf("current: %v, options: %v\n", current, options)

	var results [][]T

	if len(options) == 0 {
		return [][]T{current}
	}

	var lastO T
	for i, o := range options {
		if o == lastO {
			continue
		}
		lastO = o

		newCurrent := make([]T, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, o)

		newOptions := make([]T, 0, len(options)-1)
		newOptions = append(newOptions, options[:i]...)
		newOptions = append(newOptions, options[i+1:]...)

		subResults := Permute(newCurrent, newOptions)
		results = append(results, subResults...)
	}

	return results
}

// NOTE: スライスのcopyが多く、m*n = 10 程度で致命的に遅い。
//
// O(m^n * n) m: len(options), n: 各サブスライスの平均長
// 要素数 len(options) で、i番目の要素が options[i] であるような順列のパターンを全列挙する
// options[i]に重複した要素が含まれていても、あらかじめソートしておけば重複パターンは除かれる
func Permute2[T comparable](current []T, options [][]T) [][]T {
	var results [][]T

	if len(current) == len(options) {
		results = append(results, current)
		return results
	}

	nextVals := options[len(current)]
	var lastV T
	for _, v := range nextVals {
		if v == lastV {
			continue
		}
		lastV = v

		copyCurrent := append([]T{}, current...)
		copyCurrent = append(copyCurrent, v)
		subResults := Permute2(copyCurrent, options)
		results = append(results, subResults...)
	}

	return results
}
