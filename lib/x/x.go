// 数直線のパッケージ
package x

import "sort"

// points（昇順ソートされた数直線上の座標）の中で、baseから距離dist以内にある点の数を返す
func CountPointsInDistance(points []int, base, dist int) int {
	minX := base - dist
	maxX := base + dist

	num := len(points)

	idx1 := sort.Search(len(points), func(i int) bool {
		return points[i] >= minX
	})
	num -= idx1

	idx2 := sort.Search(len(points), func(i int) bool {
		return points[i] > maxX
	})
	num -= len(points) - idx2

	return num
}
