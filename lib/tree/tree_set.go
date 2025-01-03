package tree

import (
	"github.com/emirpasic/gods/sets/treeset"
)

// O(1)
func First[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.Begin()
	it.Next()

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

// O(1)
func Last[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.End()
	it.Prev()

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

// ちなみに、c++のmutisetはredblacktreeを使い、keyに値、valueに回数を持てば実現出来る
