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
	for i := 0; i < N; i++ {
		cube[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			cube[i][j] = readIntArr(r)
		}
	}
	sumCube := PrefixSum3D(cube)

	Q := readInt(r)
	for i := 0; i < Q; i++ {
		iarr := readIntArr(r)
		Lx, Rx, Ly, Ry, Lz, Rz := iarr[0], iarr[1], iarr[2], iarr[3], iarr[4], iarr[5]
		sum := SumFrom3DPrefixSum(sumCube, Lx, Rx, Ly, Ry, Lz, Rz)
		fmt.Fprintln(w, sum)
	}
}

//////////////
// Libs    //
/////////////

// O(cube_size)
// 三次元累積和を返す（各次元のindex0には0を入れる。）
func PrefixSum3D(cube [][][]int) [][][]int {
	X := len(cube) + 1
	Y := len(cube[0]) + 1
	Z := len(cube[0][0]) + 1

	sumCube := make([][][]int, X)
	for i := 0; i < X; i++ {
		sumCube[i] = make([][]int, Y)
		for j := 0; j < Y; j++ {
			sumCube[i][j] = make([]int, Z)
			if i == 0 || j == 0 {
				continue
			}

			copy(sumCube[i][j][1:], cube[i-1][j-1])
		}
	}

	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j][k-1]
			}
		}
	}
	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j-1][k]
			}
		}
	}
	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i-1][j][k]
			}
		}
	}
	return sumCube
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
