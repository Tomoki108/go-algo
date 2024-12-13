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

	if N <= 10 {
		fmt.Fprintln(w, N-1)
		return
	}

	count := 0
	lastCount := 0
	numOfDigits := 1
	for count < N {
		half := (numOfDigits + 1) / 2 // ex, 4 -> 2, 5 -> 3

		patterns := 1
		for i := 0; i < half; i++ {
			n := 10
			if numOfDigits != 1 && i == 0 {
				n = 9
			}
			patterns *= n
		}

		count += patterns
		lastCount = patterns

		numOfDigits++
	}
	numOfDigits--

	remainder := count - N
	ansNo := lastCount - remainder // その桁数の回文において、何番目の回文か

	minPalindromeLeft := 1 // ex, 3桁の回文 -> 10, 4桁の回文 -> 10, 5桁の回文 -> 100
	for i := 0; i < ((numOfDigits+1)/2)-1; i++ {
		minPalindromeLeft *= 10
	}
	ansLeft := minPalindromeLeft + ansNo - 1
	ansLeftStr := strconv.Itoa(ansLeft)
	ansSl := strings.Split(ansLeftStr, "")

	if numOfDigits%2 == 0 {
		length := len(ansSl)
		for i := 0; i < length; i++ {
			ansSl = append(ansSl, ansSl[length-1-i])
		}
	} else {
		length := len(ansSl)
		for i := 0; i < length-1; i++ {
			ansSl = append(ansSl, ansSl[length-2-i])
		}
	}

	fmt.Fprintln(w, strings.Join(ansSl, ""))
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
