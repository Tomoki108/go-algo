package grid

import "math"

type Coordinate struct {
	h, w int
}

func (c Coordinate) Adjacents() [4]Coordinate {
	return [4]Coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
		{c.h, c.w - 1}, // 左
		{c.h, c.w + 1}, // 右
	}
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
}

func CalcManhattanDistance(c1, c2 Coordinate) int {
	return int(math.Abs(float64(c1.h-c2.h)) + math.Abs(float64(c1.w-c2.w)))
}

// 一辺がnの正方形グリッドのマス目(hight, width)を、時計回りにtime回回転させたときの座標を返す
func RotateGridCell(n, height, width, time int) (h, w int) {
	time = time % 4
	switch time {
	case 0:
		return height, width
	case 1:
		return width, n - height + 1
	case 2:
		return n - height + 1, n - width + 1
	case 3:
		return n - width + 1, height
	}

	panic("can't reach here")
}

// 一辺がnの正方形グリッドのマス目(hight, width)が、最も外側のマス目達を1周目としたときに何周目にあるかを返す
func GetGridCellLayer(n, h, w int) int {
	return int(math.Min(math.Min(float64(h), float64(w)), math.Min(float64(n-h+1), float64(n-w+1))))
}
