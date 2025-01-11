package set

import (
	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

func NewIntSet() *set.Set[int] {
	return set.New(comparator.IntComparator)
}

func NewMultiIntSet() *set.MultiSet[int] {
	return set.NewMultiSet(comparator.IntComparator)
}
