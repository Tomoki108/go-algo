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

	N, M := read2Ints(r)
	As := readIntArr(r)
	Bs := readIntArr(r)

	segtree := NewLazySegmentTree(N)
	segtree.Build(As)

	for i := 0; i < M; i++ {
		B := Bs[i]
		A := segtree.Query(B, B+1)
		segtree.Update(B, B+1, -A)

		if N-1-B >= A {
			segtree.Update(B+1, B+1+A, 1)
		} else {
			segtree.Update(B+1, N, 1)
			rem := A - (N - 1 - B)
			if rem > N {
				segtree.Update(0, N, rem/N)
				segtree.Update(0, rem%N, 1)
			} else {
				segtree.Update(0, rem, 1)
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fprint(w, segtree.Query(i, i+1))
		if i < N-1 {
			fmt.Fprint(w, " ")
		} else {
			fmt.Fprintln(w)
		}
	}

}

// ////////////
// Libs    //
// ///////////
type LazySegmentTree struct {
	n, size int
	data    []int // 各区間の和を保持（ノードの値）
	lazy    []int // 遅延伝搬用配列
}

// NewLazySegmentTree は要素数 n の遅延セグメント木を初期化します。
func NewLazySegmentTree(n int) *LazySegmentTree {
	size := 1
	for size < n {
		size *= 2
	}
	// 2*size は完全2分木のノード数の上限
	data := make([]int, 2*size)
	lazy := make([]int, 2*size)
	return &LazySegmentTree{
		n:    n,
		size: size,
		data: data,
		lazy: lazy,
	}
}

// build は元の配列 arr からセグメント木を構築します。
func (seg *LazySegmentTree) Build(arr []int) {
	// 葉ノードに値を設定
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.size] = arr[i]
	}
	// 足りない葉は単位元（ここでは0）で初期化
	for i := len(arr); i < seg.size; i++ {
		seg.data[i+seg.size] = 0
	}
	// 内部ノードの値を構築（ここでは和）
	for i := seg.size - 1; i > 0; i-- {
		seg.data[i] = seg.data[2*i] + seg.data[2*i+1]
	}
}

// push は遅延情報を子ノードに伝搬し、現在のノードの値を更新します。
func (seg *LazySegmentTree) push(node, nl, nr int) {
	if seg.lazy[node] != 0 {
		// 現在の区間の総和に対して、遅延値を反映
		seg.data[node] += seg.lazy[node] * (nr - nl)
		// 子ノードが存在する場合は、遅延情報を子へ伝搬
		if node < seg.size {
			seg.lazy[2*node] += seg.lazy[node]
			seg.lazy[2*node+1] += seg.lazy[node]
		}
		seg.lazy[node] = 0
	}
}

// updateRec は区間 [l, r) に対して値 val を加算します。（再帰処理）
func (seg *LazySegmentTree) updateRec(l, r, val, node, nl, nr int) {
	// 遅延情報を先に処理
	seg.push(node, nl, nr)
	// 完全に区間外の場合
	if r <= nl || nr <= l {
		return
	}
	// 完全に区間内の場合
	if l <= nl && nr <= r {
		seg.lazy[node] += val
		seg.push(node, nl, nr)
		return
	}
	// 部分的に区間と重なる場合は子に伝搬
	mid := (nl + nr) / 2
	seg.updateRec(l, r, val, 2*node, nl, mid)
	seg.updateRec(l, r, val, 2*node+1, mid, nr)
	// 子の値から親の値を再計算
	seg.data[node] = seg.data[2*node] + seg.data[2*node+1]
}

// Update は区間 [l, r) に対して値 val を加算する外部インターフェースです。
func (seg *LazySegmentTree) Update(l, r, val int) {
	seg.updateRec(l, r, val, 1, 0, seg.size)
}

// queryRec は区間 [l, r) の和を取得する再帰処理です。
func (seg *LazySegmentTree) queryRec(l, r, node, nl, nr int) int {
	seg.push(node, nl, nr)
	// 完全に区間外の場合
	if r <= nl || nr <= l {
		return 0 // 単位元
	}
	// 完全に区間内の場合
	if l <= nl && nr <= r {
		return seg.data[node]
	}
	// 部分的に重なる場合は子ノードに問い合わせ
	mid := (nl + nr) / 2
	left := seg.queryRec(l, r, 2*node, nl, mid)
	right := seg.queryRec(l, r, 2*node+1, mid, nr)
	return left + right
}

// Query は区間 [l, r) の和を取得する外部インターフェースです。
func (seg *LazySegmentTree) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.size)
}

// O(n)
// 一次元累積和を返す（index0には0を入れる。）
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
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
