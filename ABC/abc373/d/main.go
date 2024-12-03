package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)

	weightMap := make(map[int]map[int]int, N)
	for i := 1; i <= N; i++ {
		weightMap[i] = make(map[int]int, N-1)
	}

	for i := 0; i < M; i++ {
		iarr := readIntArr(r)

		from := iarr[0]
		to := iarr[1]
		weight := iarr[2]

		weightMap[from][to] = weight
		weightMap[to][from] = -1 * weight
	}

	visted := make(map[int]bool, N) // 訪問済みの頂点

	xs := make([]int, N)
	xs[0] = 1

	queue := NewQueue[int]()

	for i := 1; i <= N; i++ {
		if visted[i] {
			continue
		}
		queue.Enqueue(i)

		for queue.Size() > 0 {
			from, _ := queue.Dequeue()
			visted[from] = true

			adjacents := weightMap[from]
			for to, weight := range adjacents {
				if visted[to] {
					continue
				}

				// toX - fromX = weight
				fromX := xs[from-1]
				toX := weight + fromX
				xs[to-1] = toX

				visted[to] = true
				queue.Enqueue(to)
			}
		}

	}

	// var builder strings.Builder
	// for i := 1; i <= N; i++ {
	// 	if i > 1 {
	// 		builder.WriteString(" ")
	// 	}
	// 	builder.WriteString(strconv.Itoa(xMap[i]))
	// }
	// builder.WriteString("\n")
	// fmt.Fprint(w, builder.String())

	for i, v := range xs {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, v)

		if i == N {
			fmt.Fprint(w, "\n")
		}
	}
}

//////////////
// Queue  //
/////////////

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

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readString(r *bufio.Reader) string {
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
		grid[i] = readStrArr(r)
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

// nCrの計算 O(r)
// (n * (n-1) ... * (n-r+1)) / r!
func combination(n, r int) int {
	if r > n {
		return 0
	}
	if r > n/2 {
		r = n - r // Use smaller r for efficiency
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// slices.Reverce() （Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func slReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
