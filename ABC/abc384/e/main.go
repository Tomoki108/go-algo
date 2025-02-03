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

	iarr := readIntArr(r)
	H, W, X := iarr[0], iarr[1], iarr[2]
	P, Q := read2Ints(r)

	grid := readIntGrid(r, H, true)

	visited := createGrid(H, W, false)

	pq := Heap[*pqItem]{}
	start := Coordinate{P - 1, Q - 1}
	pq.PushItem(&pqItem{strength: 0, c: start})
	visited[P-1][Q-1] = true

	strength := grid[P-1][Q-1]
	for len(pq) > 0 {
		item := pq.PopItem()

		if strength > item.strength {
			strength += item.strength / X

			adjacents := item.c.Adjacents()
			for _, adj := range adjacents {
				if adj.IsValid(H, W) && !visited[adj.h][adj.w] {
					item := &pqItem{strength: grid[adj.h][adj.w] * X, c: adj}
					pq.PushItem(item)
					visited[adj.h][adj.w] = true
				}
			}
		}
	}

	fmt.Fprintln(w, strength)
}

type pqItem struct {
	strength int
	c        Coordinate
}

func (p *pqItem) Priority() int {
	return p.strength
}

//////////////
// Libs    //
/////////////

type Coordinate struct {
	h, w int // 0-indexed
}

func (c Coordinate) Adjacents() [4]Coordinate {
	return [4]Coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
		{c.h, c.w - 1}, // 左
		{c.h, c.w + 1}, // 右
	}
}

func (c Coordinate) AdjacentsWithDiagonals() [8]Coordinate {
	return [8]Coordinate{
		{c.h - 1, c.w},     // 上
		{c.h + 1, c.w},     // 下
		{c.h, c.w - 1},     // 左
		{c.h, c.w + 1},     // 右
		{c.h - 1, c.w - 1}, // 左上
		{c.h - 1, c.w + 1}, // 右上
		{c.h + 1, c.w - 1}, // 左下
		{c.h + 1, c.w + 1}, // 右下
	}
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
}

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
