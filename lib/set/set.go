package set

import (
	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

// set.First().Value() => min
func NewIntSetAsc() *set.Set[int] {
	return set.New(comparator.IntComparator)
}

// set.First().Value() => max
func NewIntSetDesc() *set.Set[int] {
	return set.New(comparator.Reverse(comparator.IntComparator))
}

// set.First().Value() => min
func NewMultiIntSetAsc() *set.MultiSet[int] {
	return set.NewMultiSet(comparator.IntComparator)
}

// set.First().Value() => max
func NewMultiIntSetDesc() *set.MultiSet[int] {
	return set.NewMultiSet(comparator.Reverse(comparator.IntComparator))
}
