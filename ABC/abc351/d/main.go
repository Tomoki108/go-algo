package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused 9223372036854775808, 19 digits, equiv 2^63
const INT_MAX = math.MaxInt

//lint:ignore U1000 unused -9223372036854775808, 19 digits, equiv -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var H, W int

func main() {
	defer w.Flush()

	H, W = read2Ints(r)

	grid := readGrid(r, H)

	ansGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		ansGrid[i] = make([]int, W)
	}

	visitedGrid := make([][]bool, H)
	for i := 0; i < H; i++ {
		visitedGrid[i] = make([]bool, W)
	}

	ans := INT_MIN
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if grid[h][w] != "#" {
				fmt.Println("\nhi")

				moveCount := dfs(grid, ansGrid, visitedGrid, Coordinate{h, w})
				ans = max(ans, moveCount)
				ansGrid[h][w] = moveCount
			}
		}
	}

	fmt.Fprintln(w, ans)

	fmt.Println("ansGrid:")
	for _, row := range ansGrid {
		fmt.Fprintln(w, row)
	}
}

func dfs(grid [][]string, ansGrid [][]int, visitedGrid [][]bool, cell Coordinate) int {
	fmt.Printf("cell: %+v\n", cell)

	visitedGrid[cell.h][cell.w] = true

	if ansGrid[cell.h][cell.w] != 0 {
		return ansGrid[cell.h][cell.w]
	}

	canMove := true
	for _, adj := range cell.Adjacents() {
		if !adj.IsValid(H, W) {
			continue
		}

		if grid[adj.h][adj.w] == "#" {
			canMove = false
			break
		}
	}
	if !canMove {
		return 1
	}

	moveCount := 1
	for _, adj := range cell.Adjacents() {
		if !adj.IsValid(H, W) || visitedGrid[adj.h][adj.w] || grid[adj.h][adj.w] == "#" {
			continue
		}

		moveCount += dfs(grid, ansGrid, visitedGrid, adj)
	}

	return moveCount
}

//////////////
// Libs    //
/////////////

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

func (c Coordinate) CalcManhattanDistance(other Coordinate) int {
	return int(math.Abs(float64(c.h-other.h)) + math.Abs(float64(c.w-other.w)))
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
//
//lint:ignore U1000 unused
func readStr(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
//
//lint:ignore U1000 unused
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
//
//lint:ignore U1000 unused
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
//
//lint:ignore U1000 unused
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
//
//lint:ignore U1000 unused
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
//
//lint:ignore U1000 unused
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身をスペース区切りなしで出力する
//
//lint:ignore U1000 unused
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
//
//lint:ignore U1000 unused
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

//lint:ignore U1000 unused
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// O(log(exp))
// 繰り返し二乗法で x^y を計算する関数
func pow(base, exp int) int {
	// 繰り返し二乗法
	// 2^8 = 4^2^2
	// 2^9 = 4^2^2 * 2
	// この性質を利用して、基数を2乗しつつ指数を1/2にしていく

	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
