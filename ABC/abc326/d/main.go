package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	N := readInt(r)
	R := readStr(r)
	C := readStr(r)
	Rs := strings.Split(R, "")
	Cs := strings.Split(C, "")

	var set []string
	if N == 3 {
		set = []string{"A", "B", "C"}
	} else if N == 4 {
		set = []string{"A", "B", "C", "."}
	} else if N == 5 {
		set = []string{"A", "B", "C", ".", "."}
	}

	ps := Permute([]string{}, set)

	sets := make([][]string, 0, N)
	for i := 0; i < N; i++ {

	}

	set1 := make([]string, N)
	set2 := make([]string, N)
	set

	for NextPermutation(set) {

	}
}

//////////////
// Libs    //
/////////////

// NOTE: スライスのcopyが多く、n = 10 程度で致命的に遅い。NetxPermutationを推奨。
//
// O(n!) n: len(options)
// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
// options[i]に重複した要素が含まれていても、あらかじめソートしておけば重複パターンは除かれる
func Permute[T comparable](current []T, options []T) [][]T {
	var results [][]T

	if len(options) == 0 {
		ccrrent := make([]T, len(current))
		copy(ccrrent, current)

		return [][]T{ccrrent}
	}

	var lastO T
	for i, o := range options {
		if o == lastO {
			continue
		}
		lastO = o

		current = append(current, o)
		newOptions := make([]T, 0, len(options)-1)
		newOptions = append(newOptions, options[:i]...)
		newOptions = append(newOptions, options[i+1:]...)

		subResults := Permute(current, newOptions)
		results = append(results, subResults...)

		current = current[:len(current)-1]
	}

	return results
}

// NOTE: 全パターンに何らかの処理を適用したいとき、オリジナルのslに対しては別途処理を記述する
//
// O(len(sl)*len(sl)!)
// sl の要素を並び替えて、次の辞書順の順列にする
func NextPermutation[T ~int | ~string](sl []T) bool {
	n := len(sl)
	i := n - 2

	// Step1: 右から左に探索して、「スイッチポイント」を見つける:
	// 　「スイッチポイント」とは、右から見て初めて「リストの値が減少する場所」です。
	// 　例: [1, 2, 3, 6, 5, 4] の場合、3 がスイッチポイント。
	for i >= 0 && sl[i] >= sl[i+1] {
		i--
	}

	//　スイッチポイントが見つからない場合、最後の順列に到達しています。
	if i < 0 {
		return false
	}

	// Step2: スイッチポイントの右側の要素から、スイッチポイントより少しだけ大きい値を見つけ、交換します。
	// 　例: 3 を右側で最小の大きい値 4 と交換。
	j := n - 1
	for sl[j] <= sl[i] {
		j--
	}
	sl[i], sl[j] = sl[j], sl[i]

	// Step3: スイッチポイントの右側を反転して、辞書順に次の順列を作ります。
	// 　例: [1, 2, 4, 6, 5, 3] → [1, 2, 4, 3, 5, 6]。
	reverse(sl[i+1:])
	return true
}

//////////////
// Helpers //
/////////////

func dump(msg string) {
	dumpFlag := strings.Contains(strings.Join(os.Args, " "), "-dump")
	if dumpFlag {
		fmt.Println(msg)
	}
}

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

// height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArr(r)
	}

	return grid
}

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int, val T) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
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

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func sort2Ints(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func sort2IntsDesc(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
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
