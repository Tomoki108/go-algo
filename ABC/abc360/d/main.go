package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, T := read2Ints(r)
	S := readStr(r)
	Ss := strings.Split(S, "")
	Xs := readIntArr(r)

	forwardRanges := make([][2]int, 0, N)  // [start, end] (start == end - T)
	backwardRanges := make([][2]int, 0, N) // [end, start] (end == start - T)

	for i := 0; i < N; i++ {
		start := Xs[i]
		if Ss[i] == "1" {
			forwardRanges = append(forwardRanges, [2]int{start, start + T})
		} else {
			backwardRanges = append(backwardRanges, [2]int{start - T, start})
		}
	}

	sort.Slice(forwardRanges, func(i, j int) bool {
		return forwardRanges[i][0] < forwardRanges[j][0]
	})
	sort.Slice(backwardRanges, func(i, j int) bool {
		return backwardRanges[i][0] < backwardRanges[j][0]
	})
	revbackwardRange := RevSl(backwardRanges)

	ans := 0
	for _, fRange := range forwardRanges {
		left := fRange[0]
		right := fRange[1]

		tooLeftIdx := sort.Search(len(revbackwardRange), func(i int) bool {
			right2 := revbackwardRange[i][1]
			return right2 < left
		})
		tooLeftCount := len(revbackwardRange) - tooLeftIdx

		tooRightIdx := sort.Search(len(backwardRanges), func(i int) bool {
			left2 := backwardRanges[i][0]
			return right < left2
		})
		tooRightCount := len(backwardRanges) - tooRightIdx

		ans += len(backwardRanges) - tooLeftCount - tooRightCount
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

func RevSl[S ~[]E, E any](s S) S {
	lenS := len(s)
	revS := make(S, lenS)
	for i := 0; i < lenS; i++ {
		revS[i] = s[lenS-1-i]
	}

	return revS
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

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
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
