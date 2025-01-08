// xy平面のパッケージ
package xy

import "math"

// ユークリッド距離を求める
func CalcDist(fromX, fromY, toX, toY int) float64 {
	return math.Sqrt(float64((toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)))
}

// ユークリッド距離の2乗を求める
func CalcDistSquare(fromX, fromY, toX, toY int) int {
	return (toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)
}
