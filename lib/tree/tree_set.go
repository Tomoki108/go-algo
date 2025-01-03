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
// https://qiita.com/hima398/items/26e86d8dc3ab0b7c5e12#%E6%AE%8B%E3%81%A3%E3%81%A6%E3%81%84%E3%82%8B%E8%AA%B2%E9%A1%8C
