package slice

// O(n)
// 順序を無視して要素が一致しているかどうかを返す
func ElementsMatch[T comparable](sl1, sl2 []T) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	m := make(map[T]int)
	for _, v := range sl1 {
		m[v]++
	}
	for _, v := range sl2 {
		m[v]--
	}

	return len(m) == 0
}

// O(n)
// 二つのスライスが、一つ以下のインデックスを除いて要素が一致しているかを判定する
func EqualsWithAtMostOneDiff[T comparable](sl1, sl2 []T) bool {
	if len(sl1) != len(sl2) {
		panic("len(sl1) != len(sl2)")
	}

	diff := 0
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return true
}

// O(n)
// longer が shorter の何処かに一つの要素を挿入して得られるかを判定する
func EqualsWithOneInsertion[T comparable](longer, shorter []T) bool {
	if len(longer) != len(shorter)+1 {
		panic("len(longer) != len(shorter)+1")
	}

	diff := 0
	for i := 0; i < len(longer); i++ {
		if i == len(shorter) && diff == 0 {
			return true
		}

		if longer[i] != shorter[i-diff] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return true
}
