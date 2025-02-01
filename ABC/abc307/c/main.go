package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	H_A, W_A := read2Ints(r)
	grid_A := readGrid(r, H_A)

	minH_A, minW_A := INT_MAX, INT_MAX
	// maxH_A, maxW_A := 0, 0
	for h := 0; h < H_A; h++ {
		for w := 0; w < W_A; w++ {
			if grid_A[h][w] == "#" {
				minH_A, minW_A = min(h, minH_A), min(w, minW_A)
				// maxH_A, maxW_A = max(h, maxH_A), max(w, maxW_A)
			}
		}
	}

	H_B, W_B := read2Ints(r)
	grid_B := readGrid(r, H_B)
	minH_B, minW_B := INT_MAX, INT_MAX
	// maxH_B, maxW_B := 0, 0
	for h := 0; h < H_B; h++ {
		for w := 0; w < W_B; w++ {
			if grid_B[h][w] == "#" {
				minH_B, minW_B = min(h, minH_B), min(w, minW_B)
				// maxH_B, maxW_B = max(h, maxH_B), max(w, maxW_B)
			}
		}
	}

	dump("minH_A: %d, minW_A: %d\n", minH_A, minW_A)
	dump("minH_B: %d, minW_B: %d\n", minH_B, minW_B)

	H_X, W_X := read2Ints(r)
	grid_X := readGrid(r, H_X)

	for h1 := 0; h1 < H_X; h1++ {
		for w1 := 0; w1 < W_X; w1++ {
			for h2 := 0; h2 < H_X; h2++ {
			Outer:
				for w2 := 0; w2 < W_X; w2++ {

					// 「h:0」の参照を「h:minH_A」の参照としたいなら、delta_H_A = +minH_A
					// 「h:h1」の参照を「h:minH_A」の参照としたいなら、delta_H_A = +minH_A-h1
					delta_H_A := minH_A - h1
					delta_W_A := minW_A - w1
					delta_H_B := minH_B - h2
					delta_W_B := minW_B - w2

					// 正解との比較
					for hx := 0; hx < H_X; hx++ {
						for wx := 0; wx < W_X; wx++ {
							should := grid_X[hx][wx]

							cA := Coordinate{hx + delta_H_A, wx + delta_W_A}
							cB := Coordinate{hx + delta_H_B, wx + delta_W_B}

							if should == "#" {
								if !((cA.IsValid(H_A, W_A) && grid_A[cA.h][cA.w] == "#") || (cB.IsValid(H_B, W_B) && grid_B[cB.h][cB.w] == "#")) {
									continue Outer
								}
							} else {
								if cA.IsValid(H_A, W_A) && grid_A[cA.h][cA.w] == "#" {
									continue Outer
								}
								if cB.IsValid(H_B, W_B) && grid_B[cB.h][cB.w] == "#" {
									continue Outer
								}
							}
						}
					}

					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
}

//////////////
// Libs    //
/////////////

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

//////////////
// Helpers //
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

// height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArr(r)
	}

	return grid
}

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int, val T) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
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

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sort2Ints(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func sort2IntsDesc(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

func updateToMin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func updateToMax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// O(log(exp))
// 繰り返し二乗法で x^y を計算する関数
func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}

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

//////////////
// Debug   //
/////////////

var dumpFlag bool

func init() {
	args := os.Args
	dumpFlag = len(args) > 1 && args[1] == "-dump"
}

// NOTE: ループの中で使うとわずかに遅くなることに注意
func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
