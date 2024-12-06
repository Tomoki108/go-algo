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
	As := readIntArr(r)

	// DPテーブル
	// 縦：偶数体倒した || 奇数体倒した
	// 横：何体目まで処理したか
	// セルの値：X体処理時点の取得済み経験値の最大
	var table [2][]int
	table[0] = make([]int, N+1)
	table[1] = make([]int, N+1)

	table[0][0] = 0
	table[1][0] = intMin // 0体倒した時に奇数体倒した状態はありえないため、このセルの値が後で（直後の列の計算で）使われないように大きな負の値を入れておく

	for i := 0; i < N; i++ {
		A := As[i]
		evenPrev := table[0][i]
		oddPrev := table[1][i]

		// max(「偶数対倒した状態から、モンスターを逃す処理をする場合」, 「奇数体倒した状態から、モンスターを倒す処理をする場合（偶数回目なのでボーナスあり）」)
		table[0][i+1] = max(evenPrev, oddPrev+2*A)
		// max(「奇数対倒した状態から、モンスターを逃す処理をする場合」, 「偶数体倒した状態から、モンスターを倒す処理をする場合（奇数回目なのでボーナスなし」)
		table[1][i+1] = max(oddPrev, evenPrev+A)
	}

	fmt.Fprintln(w, max(table[0][N], table[1][N]))
}

//////////////
// Libs    //
/////////////

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
