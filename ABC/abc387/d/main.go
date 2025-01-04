package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused 9223372036854775808, 19 digits, equiv 2^63
const INT_MAX = math.MaxInt

//lint:ignore U1000 unused -9223372036854775808, 19 digits, equiv -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var H, W int

// 0: not visited
// 1: visited by vertical
// 2: visited by horizontal
// 3: visited by both
const (
	NOT_VISITED           = 0
	VISITED_BY_VERTICAL   = 1
	VISITED_BY_HORIZONTAL = 2
	VISITED_BY_BOTH       = 3
)

func main() {
	defer w.Flush()

	H, W = read2Ints(r)
	grid := readGrid(r, H)

	var start Coordinate
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == "S" {
				start = Coordinate{i, j}
				break
			}
		}
	}

	// 0: not visited
	// 1: visited by vertical
	// 2: visited by horizontal
	// 3: visited by both
	visitedGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		visitedGrid[i] = make([]int, W)
	}
	firstAns := bfs(start, true, grid, visitedGrid)

	fmt.Println()

	visitedGrid = make([][]int, H) // 0: not visited, 1: visited by vertical, 2: visited by horizontal, 3: visited by both
	for i := 0; i < H; i++ {
		visitedGrid[i] = make([]int, W)
	}
	secondAns := bfs(start, false, grid, visitedGrid)

	var candidates []int
	if firstAns != -1 {
		candidates = append(candidates, firstAns)
	}
	if secondAns != -1 {
		candidates = append(candidates, secondAns)
	}

	if len(candidates) == 0 {
		fmt.Fprintln(w, -1)
		return
	}

	sort.Ints(candidates)
	fmt.Fprintln(w, candidates[0])
}

func bfs(start Coordinate, lastMoveVertical bool, grid [][]string, visitedGrid [][]int) int {
	q := NewQueue[qItem]()
	q.Enqueue(qItem{start, 0, lastMoveVertical})

	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		fmt.Printf("item: %v\n", item)

		if visitedGrid[item.c.h][item.c.w] == NOT_VISITED {
			if item.lastMoveVertical {
				visitedGrid[item.c.h][item.c.w] = VISITED_BY_VERTICAL
			} else {
				visitedGrid[item.c.h][item.c.w] = VISITED_BY_HORIZONTAL
			}
		} else {
			visitedGrid[item.c.h][item.c.w] = VISITED_BY_BOTH
		}

		if grid[item.c.h][item.c.w] == "G" {
			return item.depth
		}

		// 隣接探索
		var adjacents [2]Coordinate
		var ngVisitedMark int
		if item.lastMoveVertical {
			adjacents = item.c.HorizontalAdjacents()
			ngVisitedMark = VISITED_BY_HORIZONTAL
		} else {
			adjacents = item.c.VerticalAdjacents()
			ngVisitedMark = VISITED_BY_VERTICAL
		}

		for _, adj := range adjacents {
			if !adj.IsValid(H, W) || visitedGrid[adj.h][adj.w] == VISITED_BY_BOTH || visitedGrid[adj.h][adj.w] == ngVisitedMark || grid[adj.h][adj.w] == "#" {
				continue
			}

			q.Enqueue(qItem{adj, item.depth + 1, !item.lastMoveVertical})
		}
	}

	return -1
}

type qItem struct {
	c                Coordinate
	depth            int
	lastMoveVertical bool // true: vertical, false: horizontal
}

//////////////
// Libs    //
/////////////

type Coordinate struct {
	h, w int
}

func (c Coordinate) VerticalAdjacents() [2]Coordinate {
	return [2]Coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
	}
}

func (c Coordinate) HorizontalAdjacents() [2]Coordinate {
	return [2]Coordinate{
		{c.h, c.w - 1}, // 左
		{c.h, c.w + 1}, // 右
	}
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
}

type Queue[T any] struct {
	list *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.New(),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	front := q.list.Front()
	if front == nil {
		var zero T
		return zero, false
	}
	q.list.Remove(front)
	return front.Value.(T), true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue[T]) Size() int {
	return q.list.Len()
}

// Peek returns the front element without removing it
func (q *Queue[T]) Peek() (T, bool) {
	front := q.list.Front()
	if front == nil {
		var zero T
		return zero, false
	}
	return front.Value.(T), true
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
//
//lint:ignore U1000 unused
func readStr(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
//
//lint:ignore U1000 unused
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
//
//lint:ignore U1000 unused
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
//
//lint:ignore U1000 unused
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
//
//lint:ignore U1000 unused
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
//
//lint:ignore U1000 unused
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身をスペース区切りなしで出力する
//
//lint:ignore U1000 unused
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
//
//lint:ignore U1000 unused
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

//lint:ignore U1000 unused
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
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
