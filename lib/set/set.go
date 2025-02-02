package set

import (
	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

func NewIntSetAsc() *set.Set[int] {
	return set.New(comparator.IntComparator)
}

func NewIntSetDesc() *set.Set[int] {
	return set.New(comparator.Reverse(comparator.IntComparator))
}

func GetValues[T any](s *set.Set[T]) []T {
	it := s.First()

	values := make([]T, 0, s.Size())
	for it.IsValid() {
		values = append(values, it.Value())
		it.Next()
	}

	return values
}
