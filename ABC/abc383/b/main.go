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

var effectMap map[coordinate]int

type coordinate struct {
	h, w int
}

var visited map[coordinate]bool

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	H = iarr[0]
	W = iarr[1]
	D = iarr[2]

	effectMap = make(map[coordinate]int, H*W)

	grid = readGrid(r, H)

	// fmt.Println(grid)
	// fmt.Println(len(grid[0]))
	// return

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == "." {
				effectMap[coordinate{i, j}] = 1
			}
		}
	}

	visited = make(map[coordinate]bool, H*W)

	for c := range effectMap {
		visited[c] = true
		dfs(c.h, c.w, c, 0)
	}

	fmt.Printf("%+v\n", effectMap)

	maxEffect := 0
	for _, ef := range effectMap {
		maxEffect = max(maxEffect, ef)
	}

	secondEffect := 0
	for _, ef := range effectMap {
		if ef == maxEffect {
			continue
		}
		secondEffect = max(maxEffect, ef)
	}

	fmt.Fprintln(w, maxEffect+secondEffect)
}

func dfs(i, j int, originalCoordinate coordinate, distanse int) {

	fmt.Printf("i: %d, j: %d, originalCoordinate: %+v, distanse: %d\n", i, j, originalCoordinate, distanse)

	if distanse == D {
		visited[coordinate{i, j}] = false
		return
	}

	adjacents := []coordinate{
		{i - 1, j}, // 上
		{i + 1, j}, // 下
		{i, j - 1}, // 右
		{i, j + 1}, // 左
	}

	for _, adj := range adjacents {
		if visited[adj] {
			continue
		}

		if adj.h < 0 || adj.h >= H || adj.w < 0 || adj.w >= W {
			continue
		}

		isFloor := grid[adj.h][adj.w] == "."
		_, isSet := effectMap[adj]

		if isFloor && !isSet {
			effectMap[originalCoordinate] += 1
		}

		visited[adj] = true

		dfs(adj.h, adj.w, originalCoordinate, distanse+1)
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
