package main

import (
	"bufio"
	"container/heap"
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

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N := readInt(r)

	diceNumCnts := make([]map[int]int, 0, N)
	diceKs := make([]int, 0, N)
	for i := 0; i < N; i++ {
		iarr := readIntArr(r)
		K := iarr[0]
		diceKs = append(diceKs, K)

		diceNumCnt := make(map[int]int)
		for j := 1; j <= K; j++ {
			diceNumCnt[iarr[j]]++
		}

		diceNumCnts = append(diceNumCnts, diceNumCnt)
	}

	dump("diceNumCnts: %v\n", diceNumCnts)

	ans := float64(0)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {

			cnt1 := diceNumCnts[i]
			cnt2 := diceNumCnts[j]
			K1 := diceKs[i]
			K2 := diceKs[j]

			if len(cnt1) > len(cnt2) {
				cnt1, cnt2 = cnt2, cnt1
			}

			patterns := 0
			for num, count := range cnt1 {
				patterns += count * cnt2[num]
			}

			dump("i: %d, j: %d, patterns: %d\n", i, j, patterns)
			dump("K1*K2: %d\n\n", K1*K2)

			expect := float64(patterns) / float64(K1*K2)
			ans = math.Max(ans, expect)
		}
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// 最小ヒープ（小さい値が優先）
type HeapItem interface {
	Priority() float64
}

type Heap[T HeapItem] []T

func (h *Heap[T]) PushItem(item T) {
	heap.Push(h, item)
}

func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

// to implement sort.Interface
func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].Priority() > h[j].Priority() }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *Heap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.

	*h = append(*h, x.(T))
}

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeap struct {
	iarr        []int // iarrは昇順/降順になっているとは限らないため、インデックスアクセスしないこと。
	IntHeapType IntHeapType
}

func NewIntHeap(t IntHeapType) *IntHeap {
	return &IntHeap{
		iarr:        make([]int, 0),
		IntHeapType: t,
	}
}

type IntHeapType int

const (
	MinIntHeap IntHeapType = iota // 小さい方が優先して取り出される
	MaxIntHeap                    // 大きい方が優先して取り出される
)

// O(logN)
func (h *IntHeap) PushI(i int) {
	heap.Push(h, i)
}

// O(logN)
func (h *IntHeap) PopI() int {
	return heap.Pop(h).(int)
}

// to implement sort.Interface
func (h *IntHeap) Len() int { return len(h.iarr) }
func (h *IntHeap) Less(i, j int) bool {
	if h.IntHeapType == MaxIntHeap {
		return h.iarr[i] > h.iarr[j]
	} else {
		return h.iarr[i] < h.iarr[j]
	}
}
func (h *IntHeap) Swap(i, j int) { h.iarr[i], h.iarr[j] = h.iarr[j], h.iarr[i] }

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.iarr = append(h.iarr, x.(int))
}

// DO NOT USE DIRECTLY.
// to implement heap.Interface
func (h *IntHeap) Pop() any {
	oldiarr := h.iarr
	n := len(oldiarr)
	x := oldiarr[n-1]
	h.iarr = oldiarr[0 : n-1]
	return x
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

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

func updateToMin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func updateToMax(a *int, b int) {
	if *a < b {
		*a = b
	}
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
	if exp == 0 {
		return 1
	}

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

// NOTE: ループの中で使うとわずかに遅くなることに注意
func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
