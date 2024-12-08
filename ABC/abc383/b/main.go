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

	iarr := readIntArr(r)
	H := iarr[0]
	W := iarr[1]
	D := iarr[2]
	grid := readGrid(r, H)

	floors := make([]Coordinate, 0, H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == "." {
				floors = append(floors, Coordinate{i, j})
			}
		}
	}

	maxEffect := 2
	// 「二つの床の組み合わせ」を全探索し、それらに加湿器を置いた場合の加湿範囲も全探索する
	for i := 0; i < len(floors); i++ {
		for j := i + 1; j < len(floors); j++ {
			effect := 2 // 加湿器分

			for p := 0; p < H; p++ {
				for q := 0; q < W; q++ {
					current := Coordinate{p, q}

					if (current == floors[i]) || (current == floors[j]) {
						continue
					}

					manhattanDistanse1 := floors[i].CalcManhattanDistance(current)
					manhattanDistanse2 := floors[j].CalcManhattanDistance(current)

					if (manhattanDistanse1 <= D || manhattanDistanse2 <= D) && grid[p][q] == "." {
						effect++
					}
				}
			}

			maxEffect = max(maxEffect, effect)
		}
	}

	fmt.Fprintln(w, maxEffect)
}

//////////////
// Libs    //
/////////////

type Coordinate struct {
	h, w int
}

func (c Coordinate) CalcManhattanDistance(other Coordinate) int {
	return int(math.Abs(float64(c.h-other.h)) + math.Abs(float64(c.w-other.w)))
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

		grid[i] = strings.Split(readStr(r), "")
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
