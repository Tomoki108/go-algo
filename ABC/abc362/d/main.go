package main

import (
	"bufio"
	"container/heap"
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

	N, M := read2Ints(r)

	As := readIntArr(r)

	nodeWeights := make(map[int]int, N)
	for i := 0; i < N; i++ {
		nodeWeights[i+1] = As[i]
	}

	graph := make(map[int][][2]int, N) // from => [to, weight]
	for i := 0; i < M; i++ {
		iarr := readIntArr(r)
		U, V, B := iarr[0], iarr[1], iarr[2]

		graph[U] = append(graph[U], [2]int{V, B})
		graph[V] = append(graph[V], [2]int{U, B})
	}

	ans := make(map[int]int, N-1)
	ans[1] = nodeWeights[1]

	pq := &Heap[pqItem]{}
	pq.Push(pqItem{nodeWeights[1], 1})

	isFixed := make(map[int]bool, N)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(pqItem)
		node, weightSum := item.node, item.weightSum

		if isFixed[node] {
			continue
		}
		isFixed[node] = true

		for _, adj := range graph[node] {
			nextNode := adj[0]
			edgeWeight := adj[1]

			if isFixed[nextNode] {
				continue
			}

			ws := weightSum + edgeWeight + nodeWeights[nextNode]

			if ans[nextNode] != 0 {
				ans[nextNode] = min(ans[nextNode], ws)
			} else {
				ans[nextNode] = ws
			}

			heap.Push(pq, pqItem{ws, nextNode})
		}
	}

	for i := 2; i <= N; i++ {
		fmt.Fprint(w, ans[i])

		if i != N {
			fmt.Fprint(w, " ")
		} else {
			fmt.Fprint(w, "\n")
		}
	}
}

type pqItem struct {
	weightSum int
	node      int
}

func (i pqItem) Priority() int {
	return i.weightSum
}

//////////////
// Libs    //
/////////////

type HeapItem interface {
	Priority() int
}

type Heap[T HeapItem] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
