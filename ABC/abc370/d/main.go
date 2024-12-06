package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intMax = 1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	H := iarr[0]
	W := iarr[1]
	Q := iarr[2]

	type coodinate struct {
		h, w int
	}
	type memo struct {
		td, bd, ld, rd int // （あるcoodinateから）top, bottom, left, right 方向に最低この距離分、空マスが連続していることを保証する
	}
	memos := make(map[coodinate]*memo, H*W)

	grid := make([][]bool, H)
	for i := 0; i < H; i++ {
		grid[i] = make([]bool, W)
		for j := 0; j < W; j++ {
			grid[i][j] = true
			memos[coodinate{i, j}] = &memo{0, 0, 0, 0}
		}
	}

	for i := 0; i < Q; i++ {
		R, C := read2Ints(r)
		ri := R - 1
		ci := C - 1

		memo := memos[coodinate{ri, ci}]

		if grid[ri][ci] {
			grid[ri][ci] = false
			continue
		}

		// 上方向の探索
		for hd := memo.td + 1; ri-hd >= 0; hd++ {
			if grid[ri-hd][ci] {
				grid[ri-hd][ci] = false
				memo.td = hd
				break
			}
		}

		// 下方向の探索
		for hd := memo.bd + 1; ri+hd < H; hd++ {
			if grid[ri+hd][ci] {
				grid[ri+hd][ci] = false
				memo.bd = hd
				break
			}
		}

		// 左方向の探索
		for wd := memo.ld + 1; ci-wd >= 0; wd++ {
			if grid[ri][ci-wd] {
				grid[ri][ci-wd] = false
				memo.ld = wd
				break
			}
		}

		// 右方向の探索
		for wd := memo.rd + 1; ci+wd < W; wd++ {
			if grid[ri][ci+wd] {
				grid[ri][ci+wd] = false
				memo.rd = wd
				break
			}
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] {
				ans++
			}
		}
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Hepers  //
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
		grid[i] = readStrArr(r)
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
