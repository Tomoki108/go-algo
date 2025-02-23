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

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N := readInt(r)
	grid := createGrid[string](N, N)

	var fill4edges func(start [2]int, startNo, edgeLen int)

	fill4edges = func(start [2]int, startNo, edgeLen int) {
		grid[start[0]][start[1]] = strconv.Itoa(startNo)

		height := start[0]
		width := start[1]
		no := startNo

		// 上の辺
		for width < start[1]+edgeLen {
			grid[height][width] = strconv.Itoa(no)
			no++
			width++
		}
		width--

		// 右の辺
		height++
		for height < start[0]+edgeLen {
			grid[height][width] = strconv.Itoa(no)
			no++
			height++
		}
		height--

		// 下の辺
		width--
		for width >= start[1] {
			grid[height][width] = strconv.Itoa(no)
			no++
			width--
		}
		width++

		// 左の辺
		height--
		for height > start[0] {
			grid[height][width] = strconv.Itoa(no)
			no++
			height--
		}
		height++

		edgeLen -= 2
		if edgeLen <= 0 {
			return
		}

		// for i := 0; i < N; i++ {
		// 	writeSlice(w, grid[i])
		// }
		// panic("end")

		fill4edges([2]int{start[0] + 1, start[0] + 1}, no, edgeLen)
	}

	fill4edges([2]int{0, 0}, 1, N)

	grid[(N+1)/2-1][(N+1)/2-1] = "T"

	for i := 0; i < N; i++ {
		writeSlice(w, grid[i])
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

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
	}

	return grid
}

// 文字列グリッドを出力する
func writeGrid(width *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(width, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
func writeSlice[T any](width *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(width, vs...)
}

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](width *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(width)
		return
	}

	for idx, v := range sl {
		fmt.Fprint(width, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(width)
		}
	}
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](width *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(width)
		return
	}

	for _, v := range sl {
		fmt.Fprintln(width, v)
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
