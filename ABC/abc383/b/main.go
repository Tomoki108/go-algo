package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var H, W, D int

var grid [][]string

var effectMap map[twoCoordinate]int

type coordinate struct {
	h, w int
}

type twoCoordinate [2]coordinate

var visited map[coordinate]bool

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	H = iarr[0]
	W = iarr[1]
	D = iarr[2]

	grid = readGrid(r, H)

	floors := make([]coordinate, 0, H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == "." {
				floors = append(floors, coordinate{i, j})
			}
		}
	}

	effectMap = make(map[twoCoordinate]int, H*W)

	for i := 0; i < len(floors); i++ {
		for j := i + 1; j < len(floors); j++ {
			effectMap[twoCoordinate{floors[i], floors[j]}] = 2

			i1 := floors[i].h
			j1 := floors[i].w

			i2 := floors[j].h
			j2 := floors[j].w

			for p := 0; p < H; p++ {
				for q := 0; q < W; q++ {
					if (p == i1 && q == j1) || (p == i2 && q == j2) {
						continue
					}

					manhattanDistanse1 := max(i1-p, -1*(i1-p)) + max(j1-q, -1*(j1-q))
					manhattanDistanse2 := max(i2-p, -1*(i2-p)) + max(j2-q, -1*(j2-q))

					if (manhattanDistanse1 <= D || manhattanDistanse2 <= D) && grid[p][q] == "." {
						effectMap[twoCoordinate{floors[i], floors[j]}] += 1
					}

				}
			}

		}
	}

	// fmt.Printf("effectMap: %+v\n", effectMap)

	maxEffect := 0
	for _, ef := range effectMap {
		maxEffect = max(maxEffect, ef)
	}

	fmt.Fprintln(w, maxEffect)
}

//////////////
// Libs    //
/////////////

// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
func Permute[T any](current []T, options []T) [][]T {
	var results [][]T

	cc := append([]T{}, current...)
	co := append([]T{}, options...)

	if len(co) == 0 {
		return [][]T{cc}
	}

	for i, o := range options {
		newcc := append([]T{}, cc...)
		newcc = append(newcc, o)
		newco := append([]T{}, co[:i]...)
		newco = append(newco, co[i+1:]...)

		subResults := Permute(newcc, newco)
		results = append(results, subResults...)
	}

	return results
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
