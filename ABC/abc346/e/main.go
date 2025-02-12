package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	H, W, M := iarr[0], iarr[1], iarr[2]

	type Change struct {
		queryIdx int
		color    int
	}

	rowChanges := make([]Change, H)
	colChanges := make([]Change, W)
	for i := 1; i <= M; i++ {
		iarr = readIntArr(r)
		T, A, X := iarr[0], iarr[1], iarr[2]
		A--

		switch T {
		case 1:
			rowChanges[A] = Change{queryIdx: i, color: X}
		case 2:
			colChanges[A] = Change{queryIdx: i, color: X}
		}
	}

	sort.Slice(rowChanges, func(i, j int) bool {
		return rowChanges[i].queryIdx < rowChanges[j].queryIdx
	})
	sort.Slice(colChanges, func(i, j int) bool {
		return colChanges[i].queryIdx < colChanges[j].queryIdx
	})

	totalCount := 0
	colorCount := make(map[int]int)
	for _, rowChange := range rowChanges {
		if rowChange.queryIdx == 0 {
			continue
		}

		count := sort.Search(len(colChanges), func(j int) bool {
			return colChanges[j].queryIdx > rowChange.queryIdx
		})

		if count > 0 {
			colorCount[rowChange.color] += count
			totalCount += count
		}
	}
	for _, colChange := range colChanges {
		if colChange.queryIdx == 0 {
			continue
		}

		count := sort.Search(len(rowChanges), func(j int) bool {
			return rowChanges[j].queryIdx > colChange.queryIdx
		})

		if count > 0 {
			colorCount[colChange.color] += count
			totalCount += count
		}
	}

	if H*W-totalCount > 0 {
		colorCount[0] += H*W - totalCount
	}

	fmt.Fprintln(w, len(colorCount))
	for i := 0; i <= 2*pow(10, 5); i++ {
		if count, ok := colorCount[i]; ok && count > 0 {
			fmt.Fprintln(w, i, count)
		}
	}
}

//////////////
// Libs    //
/////////////

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

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
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
