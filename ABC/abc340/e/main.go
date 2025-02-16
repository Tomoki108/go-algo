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

	segtree := NewLazySegTreeSum(N)
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
// 区間和の遅延セグメント木
// 遅延セグメント木とは：https://qiita.com/Kept1994/items/d156a1ac1fe28553bf94
type LazySegTreeSum struct {
	originSize int
	leafSize   int
	data       []int
	lazy       []int // 遅延伝搬用配列
}

// O(N) N: 元々の配列の要素数
func NewLazySegTreeSum(n int) *LazySegTreeSum {
	leafSize := 1
	for leafSize < n {
		leafSize *= 2
	}
	data := make([]int, 2*leafSize)
	lazy := make([]int, 2*leafSize)
	return &LazySegTreeSum{
		originSize: n,
		leafSize:   leafSize,
		data:       data,
		lazy:       lazy,
	}
}

// O(N) N: 元々の配列の要素数
func (seg *LazySegTreeSum) Build(arr []int) {
	for i := 0; i < len(arr); i++ {
		seg.data[i+seg.leafSize] = arr[i]
	}
	for i := seg.leafSize - 1; i > 0; i-- {
		seg.data[i] = seg.data[2*i] + seg.data[2*i+1]
	}
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) に対して値 val を加算
func (seg *LazySegTreeSum) Update(originL, originR, val int) {
	seg.updateRec(originL, originR, val, 1, 0, seg.leafSize)
}

// O(log N) N: 元々の配列の要素数
// [originL, originR) の和を取得
func (seg *LazySegTreeSum) Query(l, r int) int {
	return seg.queryRec(l, r, 1, 0, seg.leafSize)
}

func (seg *LazySegTreeSum) updateRec(originL, originR, val, currentNode, nl, nr int) {
	// 遅延情報を先に処理
	seg.push(currentNode, nl, nr)
	// 完全に区間外の場合
	if originR <= nl || nr <= originL {
		return
	}
	// 完全に区間内の場合
	if originL <= nl && nr <= originR {
		seg.lazy[currentNode] += val
		seg.push(currentNode, nl, nr)
		return
	}
	// 部分的に区間と重なる場合は子に伝搬
	mid := (nl + nr) / 2
	seg.updateRec(originL, originR, val, 2*currentNode, nl, mid)
	seg.updateRec(originL, originR, val, 2*currentNode+1, mid, nr)
	// 子の値から親の値を再計算
	seg.data[currentNode] = seg.data[2*currentNode] + seg.data[2*currentNode+1]
}

func (seg *LazySegTreeSum) queryRec(l, r, node, nl, nr int) int {
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

// 遅延情報を子ノードに伝搬し、現在のノードの値を更新
func (seg *LazySegTreeSum) push(node, nl, nr int) {
	if seg.lazy[node] != 0 {
		// 現在の区間の総和に対して、遅延値を反映
		seg.data[node] += seg.lazy[node] * (nr - nl)
		// 子ノードが存在する場合は、遅延情報を子へ伝搬
		if node < seg.leafSize {
			seg.lazy[2*node] += seg.lazy[node]
			seg.lazy[2*node+1] += seg.lazy[node]
		}
		seg.lazy[node] = 0
	}
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
