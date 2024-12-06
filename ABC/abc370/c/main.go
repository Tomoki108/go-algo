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

	S := readStr(r)
	ss := strings.Split(S, "")

	T := readStr(r)
	ts := strings.Split(T, "")

	type change struct {
		index int    // S上のインデックス
		char  string // 変更後の文字
	}
	priorChanges := make([]change, 0, len(S))
	inferiorChanges := make([]change, 0, len(S))

	for i := 0; i < len(S); i++ {
		if ss[i] != ts[i] {
			if ts[i] < ss[i] {
				priorChanges = append(priorChanges, change{i, ts[i]})
			} else {
				inferiorChanges = append(inferiorChanges, change{i, ts[i]})
			}
		}
	}

	X := make([]string, 0, len(S))
	for _, c := range priorChanges {
		ss[c.index] = c.char
		newS := strings.Join(ss, "")
		X = append(X, newS)
	}

	SlRev(inferiorChanges)
	for _, c := range inferiorChanges {
		ss[c.index] = c.char
		newS := strings.Join(ss, "")
		X = append(X, newS)
	}

	fmt.Fprintln(w, len(priorChanges)+len(inferiorChanges))
	writeSliceByLine(w, X)
}

func SlRev[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
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
