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
// スライスを指定したサイズで分割する
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

// O(n/caliculated_chunkSize)
// スライスを指定した数のチャンクに分割する
func SplitToChunks[T any](sl []T, numOfChunks int) [][]T {
	chunkSize := len(sl) / numOfChunks
	if len(sl)%numOfChunks != 0 {
		chunkSize++
	}

	return SplitByChunks(sl, chunkSize)
}

func Verticalize[T any](sl [][]T) [][]T {
	maxLen := 0
	for _, s := range sl {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	verticalized := make([][]T, maxLen)
	for i := 0; i < maxLen; i++ {
		verticalized[i] = make([]T, maxLen)
	}

	for rowI := 0; rowI < len(sl); rowI++ {
		for columnI := 0; columnI < maxLen; columnI++ {
			if columnI >= len(sl[rowI]) {
				break
			}

			verticalized[columnI][rowI] = sl[rowI][columnI]
		}
	}

	return verticalized
}
