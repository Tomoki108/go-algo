package dict

import (
	"sort"
)

func AscSortSlicesByDict[T ~int | ~string](sl [][]T) {
	sort.Slice(sl, func(i, j int) bool {
		one := sl[i]
		another := sl[j]

		length := len(one)
		if len(another) < length {
			length = len(another)
		}

		for k := 0; k < length; k++ {
			if one[k] == another[k] {
				continue
			}

			return one[k] < another[k]
		}

		return len(one) < len(another)
	})
}

func DescSortSlicesByDict[T ~int | ~string](sl [][]T) {
	sort.Slice(sl, func(i, j int) bool {
		one := sl[i]
		another := sl[j]

		length := len(one)
		if len(another) < length {
			length = len(another)
		}

		for k := 0; k < length; k++ {
			if one[k] == another[k] {
				continue
			}

			return one[k] > another[k]
		}

		return len(one) > len(another)
	})
}
