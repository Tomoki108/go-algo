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

	S := readStr(r)
	Ss := strings.Split(S, "")

	css := RunLength(Ss)

	ans := 0
	for _, cs := range css {
		num, c := SplitRLStr(cs)

		if c == "0" {
			quotient := num / 2
			rem := num % 2

			ans += quotient + rem
		} else {
			ans += num
		}
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

var Delimiter = "_"

// ランレングス圧縮を行う。[]"数+delimiter+文字種"を返す。
func RunLength(sl []string) []string {
	comp := make([]string, 0, len(sl))
	if len(sl) == 0 {
		return comp
	}

	lastChar := sl[0]
	currentLen := 0
	for i := 0; i < len(sl); i++ {
		s := sl[i]
		if s == lastChar {
			currentLen++
		} else {
			comp = append(comp, strconv.Itoa(currentLen)+Delimiter+lastChar)
			lastChar = s
			currentLen = 1
		}
	}
	comp = append(comp, strconv.Itoa(currentLen)+Delimiter+lastChar) // 最後の一文字

	return comp
}

// "数+delimiter+文字種"を分割して数と文字種を返す
func SplitRLStr(s string) (int, string) {
	strs := strings.Split(s, Delimiter)
	num, _ := strconv.Atoi(strs[0])

	return num, strs[1]
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
