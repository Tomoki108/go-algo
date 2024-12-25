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

const MOD = 998244353

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, K := read2Ints(r)
	S := readStr(r)
	Ss := strings.Split(S, "")

	firstKs := eunumeratePatterns(Ss, []string{}, K-1)

	// N * 512(=2^9)のテーブル
	// 列：index(i)まで埋めている
	// 行：index(i-K+1) ~ index(i) 文字目の文字列
	// セル：(K文字の部分列回文を含まずに）その状況に到達するパターン数
	table := make([]map[string]int, N)

	table[K-2] = make(map[string]int)
	for _, fk := range firstKs {
		str := strings.Join(fk, "")
		table[K-2][str] = 1
	}

	for i := K - 2; i < N-1; i++ {
		table[i+1] = make(map[string]int) // 次の列を初期化
		char := Ss[i+1]

		for str, count := range table[i] { // i列を元にi+1列を埋める
			if char == "A" || char == "B" {
				str += char
				if !IsPallindromeStr(str) {
					str = str[1:]
					table[i+1][str] += count % MOD
				}
			} else {
				str1 := str + "A"
				if !IsPallindromeStr(str1) {
					str1 = str1[1:]
					table[i+1][str1] += count % MOD
				}

				str2 := str + "B"
				if !IsPallindromeStr(str2) {
					str2 = str2[1:]
					table[i+1][str2] += count % MOD
				}
			}
		}
	}

	ans := 0
	for _, count := range table[N-1] {
		ans += count % MOD
	}

	fmt.Fprintln(w, ans%MOD)
}

func eunumeratePatterns(Ss, current []string, k int) [][]string {
	if len(current) == k {
		return [][]string{current}
	}

	nextIdx := len(current)
	var result [][]string
	if Ss[nextIdx] == "?" {
		copyCurrent1 := make([]string, len(current))
		copy(copyCurrent1, current)

		copyCurrent2 := make([]string, len(current))
		copy(copyCurrent2, current)

		result = append(eunumeratePatterns(Ss, append(copyCurrent1, "A"), k), eunumeratePatterns(Ss, append(copyCurrent2, "B"), k)...)
	} else {
		copyCurrent := make([]string, len(current))
		copy(copyCurrent, current)

		result = eunumeratePatterns(Ss, append(copyCurrent, Ss[nextIdx]), k)
	}

	return result
}

//////////////
// Libs    //
/////////////

// O(n/2)
func IsPallindromeStr(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
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
