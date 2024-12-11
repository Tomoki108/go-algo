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

func main() {
	defer w.Flush()

	N := readInt(r)

	cube := make([][][]int, N)
	sumCube := make([][][]int, N)

	for i := 0; i < N; i++ {
		cube[i] = make([][]int, N)
		sumCube[i] = make([][]int, N)

		for j := 0; j < N; j++ {

			iarr := readIntArr(r)
			cube[i][j] = iarr
			sumCube[i][j] = PrefixSum(iarr)
		}
	}

	Q := readInt(r)

	for i := 0; i < Q; i++ {

		iarr := readIntArr(r)
		Lx, Rx, Ly, Ry, Lz, Rz := iarr[0], iarr[1], iarr[2], iarr[3], iarr[4], iarr[5]

		sum := 0
		for j := Lx - 1; j < Rx; j++ {
			for k := Ly - 1; k < Ry; k++ {
				// RzIdx := Rz - 1
				// LzIdx := Lz - 2

				// fmt.Printf("Lx: %d, Rx: %d, Ly: %d, Ry: %d, Lz: %d, Rz: %d\n", Lx, Rx, Ly, Ry, Lz, Rz)
				// fmt.Printf("sumCube[j][k][RzIdx]: %d, sumCube[j][k][LzIdx]: %d\n", sumCube[j][k][RzIdx], sumCube[j][k][LzIdx])

				sum += sumCube[j][k][Rz] - sumCube[j][k][Lz-1]

			}
		}

		fmt.Fprintln(w, sum)
	}
}

//////////////
// Libs    //
/////////////

func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
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
