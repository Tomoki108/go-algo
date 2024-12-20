package slice

// O(n)
// slices.Reverce() と同じ（Goのバージョンが1.21以前だと使えないため）
func SlRev[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// O(n)
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

// O(n)
func Deduplicate[T comparable](sl []T) []T {
	m := map[T]bool{}
	for _, v := range sl {
		m[v] = true
	}

	var deduped []T
	for k := range m {
		deduped = append(deduped, k)
	}

	return deduped
}

// O(n/size)
func SplitByChunks[T any](sl []T, chunkSize int) [][]T {
	if len(sl) == 0 {
		return [][]T{}
	}

	chunks := make([][]T, 0, (len(sl)+chunkSize-1)/chunkSize) // 余りを考慮したlengthの計算
	for chunkSize < len(sl) {
		chunks = append(chunks, sl[0:chunkSize])
		sl = sl[chunkSize:]
	}

	return append(chunks, sl)
}
