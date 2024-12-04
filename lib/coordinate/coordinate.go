package coordinate

import "math"

func CalcDistance(fromX, fromY, toX, toY int) float64 {
	return math.Sqrt(float64((toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)))
}
