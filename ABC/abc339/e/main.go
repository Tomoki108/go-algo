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

	N, D := read2Ints(r)
	As := readIntArr(r)

	maxNum := 5*pow(10, 5) + 1
	length := make([]int, maxNum) // length[i]: iという数字を最後に採用している場合の、最大の部分列の長さ

	// dp[j]: jを採用する場合の最大の部分列の長さ.
	// 何idx目まで処理したかは陽に持たず、0~N-1までのループで同じ配列を使い回す.
	dp := NewSegTreeMax(length)
	for i := 0; i < N; i++ {
		left := max(As[i]-D, 0)         // As[i]を採用する場合に、そこに遷移できる最小の数
		right := min(As[i]+D+1, maxNum) // As[i]を採用する場合に、そこに遷移できる最大の数+1

		dp.Update(As[i], dp.Query(left, right)+1)
	}

	ans := dp.Query(0, maxNum)
	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// 区間最大値のセグメント木
// セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type SegTreeMax struct {
	n    int
	size int
	tree []int
}

func NewSegTreeMax(arr []int) *SegTreeMax {
	n := len(arr)
	size := 1
	for size < n {
		size *= 2
	}
	tree := make([]int, 2*size)
	minVal := -1 << 63
	for i := 0; i < 2*size; i++ {
		tree[i] = minVal
	}
	for i, v := range arr {
		tree[size+i] = v
	}
	for i := size - 1; i > 0; i-- {
		tree[i] = max(tree[2*i], tree[2*i+1])
	}
	return &SegTreeMax{
		n:    n,
		size: size,
		tree: tree,
	}
}

// O(log N) N: 元々の配列の要素数
// idx番目の値をvalueに更新
func (st *SegTreeMax) Update(i, val int) {
	i += st.size
	st.tree[i] = val
	for i > 1 {
		i /= 2
		st.tree[i] = max(st.tree[2*i], st.tree[2*i+1])
	}
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) の範囲の最大値を取得
func (st *SegTreeMax) Query(l, r int) int {
	minVal := -1 << 63
	res := minVal
	l += st.size
	r += st.size
	for l < r {
		if l%2 == 1 {
			res = max(res, st.tree[l])
			l++
		}
		if r%2 == 1 {
			r--
			res = max(res, st.tree[r])
		}
		l /= 2
		r /= 2
	}
	return res
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
