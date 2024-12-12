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

	cube := make([][][]int, N+1)
	sumCube := make([][][]int, N+1)
	for i := 0; i < N+1; i++ {
		cube[i] = make([][]int, N+1)
		sumCube[i] = make([][]int, N+1)
		for j := 0; j < N+1; j++ {
			cube[i][j] = make([]int, N+1)
			sumCube[i][j] = make([]int, N+1)
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			iarr := readIntArr(r)
			for k := 0; k < N; k++ {
				cube[i+1][j+1][k+1] = iarr[k]
				sumCube[i+1][j+1][k+1] = iarr[k]
			}
		}
	}

	for x := 1; x < N+1; x++ {
		for y := 1; y < N+1; y++ {
			for z := 1; z < N+1; z++ {
				sumCube[x][y][z] += sumCube[x-1][y][z]
			}
		}
	}
	for x := 1; x < N+1; x++ {
		for y := 1; y < N+1; y++ {
			for z := 1; z < N+1; z++ {
				sumCube[x][y][z] += sumCube[x][y-1][z]
			}
		}
	}
	for x := 1; x < N+1; x++ {
		for y := 1; y < N+1; y++ {
			for z := 1; z < N+1; z++ {
				sumCube[x][y][z] += sumCube[x][y][z-1]
			}
		}
	}

	Q := readInt(r)
	for i := 0; i < Q; i++ {
		iarr := readIntArr(r)
		Lx, Rx, Ly, Ry, Lz, Rz := iarr[0], iarr[1], iarr[2], iarr[3], iarr[4], iarr[5]
		sum := SumFrom3DPrefixSum(sumCube, Lx, Rx, Ly, Ry, Lz, Rz)
		fmt.Fprintln(w, sum)
	}
}

// 三次元累積和から、任意の範囲の和を求める
// sumCubには、x, y, z方向に番兵（余分な空の一行）が含まれているものとする
// Lx, Rxは、その軸における範囲指定 => x方向には、Rxの累積和からLx-1の累積和を引く
func SumFrom3DPrefixSum(sumCube [][][]int, Lx, Rx, Ly, Ry, Lz, Rz int) int {
	Lx--
	Ly--
	Lz--

	// 包除原理
	result := sumCube[Rx][Ry][Rz]
	result -= sumCube[Lx][Ry][Rz]
	result -= sumCube[Rx][Ly][Rz]
	result -= sumCube[Rx][Ry][Lz]

	result += sumCube[Lx][Ly][Rz]
	result += sumCube[Lx][Ry][Lz]
	result += sumCube[Rx][Ly][Lz]

	result -= sumCube[Lx][Ly][Lz]

	return result
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
