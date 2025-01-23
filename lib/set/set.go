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
