package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	type dot struct {
		x, y int
	}

	dots := make([]dot, 0, 3)
	for i := 0; i < 3; i++ {
		x, y := read2Ints(r)
		dots = append(dots, dot{x, y})
	}

	edge1Square := CalcDistanceSquare(dots[0].x, dots[0].y, dots[1].x, dots[1].y)
	edge2Square := CalcDistanceSquare(dots[1].x, dots[1].y, dots[2].x, dots[2].y)
	edge3Square := CalcDistanceSquare(dots[2].x, dots[2].y, dots[0].x, dots[0].y)

	if edge1Square == edge2Square+edge3Square {
		fmt.Fprintln(w, "Yes")
		return
	}

	if edge2Square == edge1Square+edge3Square {
		fmt.Fprintln(w, "Yes")
		return
	}

	if edge3Square == edge1Square+edge2Square {
		fmt.Fprintln(w, "Yes")
		return
	}

	fmt.Fprintln(w, "No")
}

//////////////
// Libs    //
/////////////

func CalcDistance(fromX, fromY, toX, toY int) float64 {
	return math.Sqrt(float64((toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)))
}

func CalcDistanceSquare(fromX, fromY, toX, toY int) int {
	return (toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readStr(r *bufio.Reader) string {
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
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
