package slice

// slices.Reverce() と同じ（Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func SlRev[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func RevSl[S ~[]E, E any](s S) S {
	lenS := len(s)
	revS := make(S, lenS)
	for i := 0; i < lenS; i++ {
		revS[i] = s[lenS-1-i]
	}

	return revS
}

// どちらか一方のスライスにのみ含まれる要素で構成されたスライスを返す
func SlDiff[T comparable](slice1, slice2 []T) []T {
	// 要素の出現回数を記録するマップ
	countMap := make(map[T]int)

	// slice1 の要素をマップに記録
	for _, v := range slice1 {
		countMap[v]++
	}

	// slice2 の要素をマップに記録
	for _, v := range slice2 {
		countMap[v]++
	}

	// 片方にのみ含まれる要素を収集
	var result []T
	for k, v := range countMap {
		if v == 1 { // 1度だけ出現した要素を追加
			result = append(result, k)
		}
	}

	return result
}
