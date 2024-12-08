package main

import (
	"bufio"
	"container/list"
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
	H := iarr[0]
	W := iarr[1]
	D := iarr[2]

	grid := readGrid(r, H)
	// 重複カウントしないように、加湿した所をgridで管理しなければならない
	wateredGrid := make([][]bool, H)
	for i := 0; i < H; i++ {
		wateredGrid[i] = make([]bool, W)
	}

	type queueItem struct {
		c   Coordinate
		dep int
	}

	queue := NewQueue[queueItem]()
	var visited map[Coordinate]bool

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == "H" {
				ans++
				wateredGrid[i][j] = true
				visited = make(map[Coordinate]bool, H*W)

				c := Coordinate{i, j}
				visited[c] = true

				item := queueItem{c, 0}
				queue.Enqueue(item)

				for !queue.IsEmpty() {
					item, _ := queue.Dequeue()
					if item.dep == D {
						continue
					}

					for _, adj := range item.c.Adjacents() {
						if !adj.IsValid(H, W) || visited[adj] || grid[adj.h][adj.w] == "#" || grid[adj.h][adj.w] == "H" {
							continue
						}

						if !wateredGrid[adj.h][adj.w] {
							ans++
							wateredGrid[adj.h][adj.w] = true
						}
						visited[adj] = true

						queue.Enqueue(queueItem{adj, item.dep + 1})
					}
				}
			}
		}
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

type Coordinate struct {
	h, w int
}

func (c Coordinate) Adjacents() [4]Coordinate {
	return [4]Coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
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

type coordinate struct {
	h, w int
}

func (c coordinate) adjacents() [4]coordinate {
	return [4]coordinate{
		{c.h - 1, c.w}, // 上
		{c.h + 1, c.w}, // 下
		{c.h, c.w - 1}, // 左
		{c.h, c.w + 1}, // 右
	}
}

// height行の文字列グリッドを読み込む
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {

		grid[i] = strings.Split(readStr(r), "")
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
