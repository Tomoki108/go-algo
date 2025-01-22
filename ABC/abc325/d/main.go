package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
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

	sections := make([][2]int, 0, N) // []{start, end}
	for i := 0; i < N; i++ {
		T, D := read2Ints(r)
		sections = append(sections, [2]int{T, T + D})
	}
	sort.Slice(sections, func(i, j int) bool {
		return sections[i][0] < sections[j][0]
	})

	ans := 0

	t := sections[0][0] // 現在時刻
	sectionIdx := 0     // 次にこのindexの区間から、endをpriority queueに追加しようとする。

	pq := NewIntHeap(MinIntHeap) // すでに開始している区間の終了時間を格納。最も早く終了するものが先頭に来る。
	for sectionIdx < N || pq.Len() != 0 {
		dump(fmt.Sprintf("[at start] t: %d, sectionIdx: %d, pq: %v, ans: %d", t, sectionIdx, pq.iarr, ans))

		// push end-time of start-time t
		for sectionIdx < N && sections[sectionIdx][0] == t {
			pq.PushI(sections[sectionIdx][1])
			dump(fmt.Sprintf("[pushed] %d", sections[sectionIdx][1]))
			sectionIdx++
		}

		// waste expired end-time
		for pq.Len() > 0 {
			last := pq.PopI()
			if last >= t {
				dump(fmt.Sprintf("[wasted] %d", last))
				pq.PushI(last)
				break
			}
		}

		// if pq.Len() == 0 {
		// 	if sectionIdx < N {
		// 		t = sections[sectionIdx][0]
		// 		continue
		// 	}
		// 	break
		// }

		// pop closest end-time
		if pq.Len() > 0 {
			poped := pq.PopI()
			dump(fmt.Sprintf("[poped] %d", poped))
			ans++
		}

		// update t
		if pq.Len() == 0 && sectionIdx < N {
			t = sections[sectionIdx][0]
		} else {
			t++
		}

		dump(fmt.Sprintf("[at end] t: %d, sectionIdx: %d, pq: %v, ans %d\n", t, sectionIdx, pq.iarr, ans))
	}

	fmt.Println(ans)
}

//////////////
// Libs    //
/////////////

type IntHeap struct {
	iarr        []int
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
	MinIntHeap IntHeapType = iota // 大きい方が優先
	MaxIntHeap                    // 小さい方が優先
)

func (h *IntHeap) PushI(i int) {
	heap.Push(h, i)
}

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
