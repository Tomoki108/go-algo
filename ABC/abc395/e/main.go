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

// 9223372036854775807, 19 digits, 2^63 - 1
const INT_MAX = math.MaxInt64

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt64

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	N, M, X := iarr[0], iarr[1], iarr[2]

	graph := make([][][2]int, N) // node, flipped
	for i := 0; i < M; i++ {
		u, v := read2Ints(r)
		u--
		v--
		graph[u] = append(graph[u], [2]int{v, 0})
		graph[v] = append(graph[v], [2]int{u, 1})
	}

	nodeCostNonFlipped := make([]int, N)
	nodeCostsFlipped := make([]int, N)
	for i := 0; i < N; i++ {
		nodeCostNonFlipped[i] = INT_MAX
		nodeCostsFlipped[i] = INT_MAX
	}

	nodeCostsFlipped[0] = X
	nodesCosts := [2][]int{nodeCostNonFlipped, nodeCostsFlipped}

	pq := &Heap[pqItem]{}
	pq.PushItem(pqItem{node: 0, weightSum: 0, flipped: 0})
	pq.PushItem(pqItem{node: 0, weightSum: X, flipped: 1})

	isFixedNonFlipped := make([]bool, N)
	isFixedFlipped := make([]bool, N)
	isFixedNonFlipped[0] = true
	isFixedFlipped[0] = true

	isFixed := [2][]bool{isFixedNonFlipped, isFixedFlipped}

	for pq.Len() > 0 {
		item := pq.PopItem()
		dump("item: %v\n", item)

		if isFixed[item.flipped][item.node] {
			continue
		}
		isFixed[item.flipped][item.node] = true

		nodesCosts[item.flipped][item.node] = item.weightSum

		adjacents := graph[item.node]
		for _, adj := range adjacents {
			newWeightSum := item.weightSum + 1
			newFlipped := item.flipped
			if adj[1] != item.flipped {
				newWeightSum += X
				newFlipped = adj[1]
			}

			if nodesCosts[newFlipped][adj[0]] != INT_MAX {
				nodesCosts[newFlipped][adj[0]] = min(nodesCosts[newFlipped][adj[0]], newWeightSum)
			} else {
				nodesCosts[newFlipped][adj[0]] = newWeightSum
			}
			pq.PushItem(pqItem{
				weightSum: newWeightSum,
				node:      adj[0],
				flipped:   newFlipped,
			})
		}
	}

	dump("nodesCosts: %v\n", nodesCosts)

	costA := nodesCosts[0][N-1]
	costB := nodesCosts[1][N-1]

	fmt.Fprintln(w, min(costA, costB))
}

type pqItem struct {
	weightSum int
	node      int
	flipped   int
}

func (p pqItem) Priority() int {
	return p.weightSum
}

//////////////
// Libs    //
/////////////

// 最小ヒープ（小さい値が優先）
type HeapItem interface {
	Priority() int
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
func (h Heap[T]) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
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
func readIntGrid(r *bufio.Reader, height int, withSpace bool) [][]int {
	if withSpace {
		grid := make([][]int, height)
		for i := 0; i < height; i++ {
			grid[i] = readIntArr(r)
		}
		return grid
	}

	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		strs := strings.Split(str, "")

		grid[i] = make([]int, len(strs))
		for j, s := range strs {
			grid[i][j], _ = strconv.Atoi(s)
		}
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

func btoi(b string) int {
	num, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
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
