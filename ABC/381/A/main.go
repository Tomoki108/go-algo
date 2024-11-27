package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N := readInt(r)
	S := readString(r)
	sl := strings.Split(S, "")

	if N%2 == 0 {
		fmt.Fprint(w, "No")
		return
	}

	for i := 0; i < (N+1)/2-1; i++ {
		if sl[i] != "1" {
			fmt.Fprint(w, "No")
			return
		}
	}

	if sl[(N+1)/2-1] != "/" {
		fmt.Fprint(w, "No")
		return
	}

	for i := (N + 1) / 2; i < N; i++ {
		if sl[i] != "2" {
			fmt.Fprint(w, "No")
			return
		}
	}

	fmt.Fprint(w, "Yes")
}

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readString(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
func readIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}

// height行の文字列グリッドを読み込む
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		grid[i] = readStrArr(r)
	}

	return grid
}

// 文字列グリッドを出力する
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// 一辺がnの正方形グリッドのマス目(hight, width)を、時計回りにtime回回転させたときの座標を返す
func rotateGridCell(n, height, width, time int) (h, w int) {
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
func getGridCellLayer(n, h, w int) int {
	return int(math.Min(math.Min(float64(h), float64(w)), math.Min(float64(n-h+1), float64(n-w+1))))
}

// nCrの計算 O(r)
// (n * (n-1) ... * (n-r+1)) / r!
func combination(n, r int) int {
	if r > n {
		return 0
	}
	if r > n/2 {
		r = n - r // Use smaller r for efficiency
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// slices.Reverce() （Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func slReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}