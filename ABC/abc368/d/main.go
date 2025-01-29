// steiner tree, シュタイナー木の問題
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

var tree map[int][]int
var pickedVertexes map[int]struct{}
var numOfPicked map[int]int // そのノード以下の階層に、指定された頂点がいくつ含まれているか

func main() {
	defer w.Flush()

	N, K := read2Ints(r)

	tree = make(map[int][]int, N)
	for i := 0; i < N-1; i++ {
		A, B := read2Ints(r)
		tree[A] = append(tree[A], B)
		tree[B] = append(tree[B], A)
	}

	Vs := readIntArr(r)
	pickedVertexes = make(map[int]struct{}, K)
	for i := 0; i < K; i++ {
		pickedVertexes[Vs[i]] = struct{}{}
	}

	numOfPicked = make(map[int]int, N)
	dfs(Vs[0], -1, 0) // 任意の指定された頂点からDFSを開始する

	ans := 0
	for _, v := range numOfPicked {
		if v != 0 {
			ans++
		}
	}

	fmt.Fprintln(w, ans)
}

func dfs(vertex, parent, numOfFoundPicked int) (nfp int) {
	found := 0
	for _, v := range tree[vertex] {
		if v == parent {
			continue
		}

		found += dfs(v, vertex, numOfFoundPicked)
	}

	_, picked := pickedVertexes[vertex]
	if picked {
		numOfPicked[vertex] = numOfFoundPicked + found + 1
		return numOfFoundPicked + found + 1
	}

	numOfPicked[vertex] = numOfFoundPicked + found
	return numOfFoundPicked + found
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
