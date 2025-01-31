package grid

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	h, w int // 0-indexed
}

func (c Coordinate) Adjacents() [4]Coordinate {
	return [4]Coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
		{c.h, c.w - 1}, // 左
		{c.h, c.w + 1}, // 右
	}
}

func (c Coordinate) AdjacentsWithDiagonals() [8]Coordinate {
	return [8]Coordinate{
		{c.h - 1, c.w},     // 上
		{c.h + 1, c.w},     // 下
		{c.h, c.w - 1},     // 左
		{c.h, c.w + 1},     // 右
		{c.h - 1, c.w - 1}, // 左上
		{c.h - 1, c.w + 1}, // 右上
		{c.h + 1, c.w - 1}, // 左下
		{c.h + 1, c.w + 1}, // 右下
	}
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
}

func (c Coordinate) CalcManhattanDistance(other Coordinate) int {
	return int(math.Abs(float64(c.h-other.h)) + math.Abs(float64(c.w-other.w)))
}

// H行W列の文字列グリッドを文字列に変換（マップのキー用など）
func GridToString(H, W int, grid [][]string) string {
	str := ""
	for i := 0; i < H; i++ {
		str += strings.Join(grid[i], "_")
	}
	return str
}

// H行W列の整数グリッドを文字列に変換（マップのキー用など）
func IntGridToString(H, W int, grid [][]int) string {
	str := ""
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 && j == 0 {
				str += strconv.Itoa(grid[i][j])
			} else {
				str += "_" + strconv.Itoa(grid[i][j])
			}
		}
	}
	return str
}

// O(H*W)
// T型グリッドのコピーを作成する
func CopyGrid[T any](grid [][]T) [][]T {
	H := len(grid)
	W := len(grid[0])
	res := make([][]T, H)
	for i := 0; i < H; i++ {
		res[i] = make([]T, W)
		copy(res[i], grid[i])
	}
	return res
}

// 一辺がnの正方形グリッドのマス目(hight, width)を、時計回りにtime回回転させたときの座標を返す
func RotateSquareGridCell(n, height, width, time int) (h, w int) {
	time = time % 4
	switch time {
	case 0:
		return height, width
	case 1:
		return width, n - height - 1
	case 2:
		return n - height - 1, n - width - 1
	case 3:
		return n - width - 1, height
	}

	panic("can't reach here")
}

// 一辺がnの正方形グリッドのマス目(hight, width)が、最も外側のマス目達を1周目としたときに何周目にあるかを返す
func GetSquareGridCellLayer(n, h, w int) int {
	return int(math.Min(math.Min(float64(h), float64(w)), math.Min(float64(n-h+1), float64(n-w+1))))
}
