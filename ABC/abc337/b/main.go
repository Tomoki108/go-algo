package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	S := readStr(r)
	Ss := strings.Split(S, "")

	compressed := RunLength(Ss, "_")

	if len(compressed) > 3 {
		fmt.Fprintln(w, "No")
		return
	}

	expectedChars := [][]string{
		{"A", "B", "C"},
		{"A", "B"},
		{"A", "C"},
		{"B", "C"},
		{"A"},
		{"B"},
		{"C"},
	}

	chars := make([]string, 0, 3)
	for _, c := range compressed {
		_, char := SplitRLStr(c, "_")
		chars = append(chars, char)
	}

	matched := false
	for _, expected := range expectedChars {
		if reflect.DeepEqual(expected, chars) {
			matched = true
			break
		}
	}

	if matched {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}

//////////////
// Libs    //
/////////////

// O(n)
// ランレングス圧縮を行う。[]"数+delimiter+文字種"を返す。
func RunLength(sl []string, delimiter string) []string {
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
			comp = append(comp, strconv.Itoa(currentLen)+delimiter+lastChar)
			lastChar = s
			currentLen = 1
		}
	}
	comp = append(comp, strconv.Itoa(currentLen)+delimiter+lastChar) // 最後の一文字

	return comp
}

// O(1)
// "数+delimiter+文字種"を分割して数と文字種を返す
func SplitRLStr(s, delimiter string) (int, string) {
	strs := strings.Split(s, delimiter)
	num, _ := strconv.Atoi(strs[0])

	return num, strs[1]
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

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
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
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

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

// O(log(exp))
// 繰り返し二乗法で x^y を計算する関数
func pow(base, exp int) int {
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
