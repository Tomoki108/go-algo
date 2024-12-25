package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
const intMax = 1 << 62

//lint:ignore U1000 unused
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N := readInt(r)

	size := Pow(3, N)

	grid := createGrid(size)

	paintCenterWithWhiteRecursively(grid, size)

	writeGrid2(w, grid)
}

func createGrid(size int) [][]*string {
	grid := make([][]*string, size)
	for i := 0; i < size; i++ {
		str := strings.Repeat("#", size)
		strs := strings.Split(str, "")
		strPtrs := make([]*string, 0, size)
		for _, s := range strs {
			strPtrs = append(strPtrs, &s)
		}

		grid[i] = strPtrs
	}

	return grid
}

func writeGrid2(w *bufio.Writer, grid [][]*string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Fprint(w, *grid[i][j])
		}
		fmt.Fprintln(w)
	}
}

func paintCenterWithWhiteRecursively(grid [][]*string, size int) {
	blockSize := size / 3

	rowIdx := 0
	for no := 1; no <= 9; no++ {
		blockGrid := make([][]*string, 0, blockSize)
		wIdx := (no % 3) * blockSize

		// fmt.Fprintf(w, "no: %d, blockSize: %d, wIdx: %d, rowIdx: %d\n", no, blockSize, wIdx, rowIdx)

		for i := 0; i < blockSize; i++ {
			blockGrid = append(blockGrid, grid[rowIdx+i][wIdx:wIdx+blockSize])
		}
		// blockGrid = append(blockGrid, grid[rowIdx+1][wIdx:wIdx+blockSize])
		// blockGrid = append(blockGrid, grid[rowIdx+2][wIdx:wIdx+blockSize])

		if no == 5 {
			// fmt.Fprintf(w, "hey! blockSize: %d, no: %d\n", blockSize, no)
			// writeGrid2(w, blockGrid)
			// fmt.Fprintln(w)

			for _, row := range blockGrid {
				for _, cell := range row {
					*cell = "."
				}
			}
		} else if blockSize > 1 {
			paintCenterWithWhiteRecursively(blockGrid, blockSize)
		}

		//		fmt.Fprintf(w, "blockSize: %d, no: %d\n", blockSize, no)
		// fmt.Fprintf(w, "len(blockGrid[0]): %d\n", len(blockGrid[0]))
		// writeGrid2(w, blockGrid)
		// fmt.Fprintln(w)

		if no%3 == 0 {
			rowIdx += blockSize
		}
	}

}

// func isCenterBlock(size, h, w int) bool {
// 	blockSize := size / 3

// 	if h > blockSize && h <= blockSize*2 && w > blockSize && w <= blockSize*2 {
// 		return true
// 	}

// 	return false
// }

//////////////
// Libs    //
/////////////

// O(log(exp))
// 繰り返し二乗法で x^y を計算する関数
func Pow(base, exp int) int {
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
