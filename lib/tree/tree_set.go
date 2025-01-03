package tree

import (
	"github.com/emirpasic/gods/sets/treeset"
)

// O(logN)
func First[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.Begin()
	it.Next() // ここでO(logN)かかる。ルートノードまで遡ってから、もっとも左の最小ノードまで降りていく

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

// O(logN)
func Last[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.End()
	it.Prev() // ここでO(logN)かかる。ルートノードまで遡ってから、もっとも右の最大ノードまで降りていく

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

// ちなみに、c++のmutisetはredblacktreeを使い、keyに値、valueに回数を持てば実現出来る
