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

	N, T := readIntStr(r)
	Ts := strings.Split(T, "")

	Ss := make([]string, 0, N)
	for i := 0; i < N; i++ {
		Ss = append(Ss, readStr(r))
	}

	ans := 0

	// 自身を二倍したものが条件を満たす場合、ans++
	for i := 0; i < N; i++ {
		wSs := Ss[i] + Ss[i]
		sl := strings.Split(wSs, "")

		matchIdx := -1
		tIdx := 0
		for j := 0; j < len(sl); j++ {
			if sl[j] == Ts[tIdx] {
				matchIdx++
				if tIdx == len(Ts)-1 {
					break
				}
				tIdx++
			}
		}
		if matchIdx == len(Ts)-1 {
			ans++
		}
	}

	// 自身より前の要素で、順に結合して条件を満たす組み合わせの数をansに足す
	solve := func(ss []string) {
		ft := NewFenwickTree(len(Ts))
		for i := 0; i < N; i++ {
			sl := strings.Split(ss[i], "")

			revMatchIdx := len(Ts)
			revTIdx := len(Ts) - 1
			for j := len(sl) - 1; j >= 0; j-- {
				if sl[j] == Ts[revTIdx] {
					revMatchIdx--
					if revTIdx == 0 {
						break
					}
					revTIdx--
				}
			}
			if revMatchIdx != len(Ts) {
				if revMatchIdx == 0 {
					ans += i
				} else {
					pairs := ft.RangeSum(revMatchIdx-1, len(Ts)-1)
					ans += pairs
				}
			} else {
				ans += ft.At(len(Ts) - 1)
			}

			matchIdx := -1
			tIdx := 0
			for j := 0; j < len(sl); j++ {
				if sl[j] == Ts[tIdx] {
					matchIdx++
					if tIdx == len(Ts)-1 {
						break
					}
					tIdx++
				}
			}
			if matchIdx != -1 {
				ft.Update(matchIdx, 1)
			}
		}
	}

	solve(Ss)
	solve(RevSl(Ss))

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// O(n)
func RevSl[S ~[]E, E any](s S) S {
	lenS := len(s)
	revS := make(S, lenS)
	for i := 0; i < lenS; i++ {
		revS[i] = s[lenS-1-i]
	}

	return revS
}

// 数列の区間和の取得、一点更新を O(log n) で行うデータ構造
// できることはセグメント木の完全下位互換だが、定数倍が小さい
type FenwickTree struct {
	n    int
	tree []int
}

// O(n)
// n+1 の長さのフェンウィック木を作成する.
// インターフェースでは 0-indexed で、内部では 1-indexed で扱うため+1.
func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		n:    n,
		tree: make([]int, n+1),
	}
}

// O(log n)
func (ft *FenwickTree) Update(i int, delta int) {
	i++ // 内部は 1-indexed として扱うため
	for i <= ft.n {
		ft.tree[i] += delta
		i += i & -i // 次の更新対象のインデックスへ
	}
}

// O(log n)
// 区間 [0, i] (0-indexed) の和を返す
func (ft *FenwickTree) Sum(i int) int {
	s := 0
	i++ // 内部は 1-indexed として扱うため
	for i > 0 {
		s += ft.tree[i]
		i -= i & -i
	}
	return s
}

// O(log n)
// 区間 [l, r] (0-indexed) の和を返す
func (ft *FenwickTree) RangeSum(l, r int) int {
	if l == 0 {
		return ft.Sum(r)
	}
	return ft.Sum(r) - ft.Sum(l-1)
}

// O(log n)
// index i (0-indexed)の要素を取得する
func (ft *FenwickTree) At(i int) int {
	return ft.RangeSum(i, i)
}

//////////////
// Helpers //
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

// １行の「整数 文字列」のみの入力を読み込む
func readIntStr(r *bufio.Reader) (int, string) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i, _ := strconv.Atoi(strs[0])
	return i, strs[1]
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

//////////////
// Debug   //
/////////////

var dumpFlag bool

func init() {
	args := os.Args
	dumpFlag = len(args) > 1 && args[1] == "-dump"
}

func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
