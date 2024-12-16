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

	type queueItem struct {
		node, weightSum int
	}
	queue := NewQueue[queueItem]()

	ansMap := make(map[int]int, N-1)

	for goal := 2; goal <= N; goal++ {
		ansMap[goal] = intMax
		queue.Enqueue(queueItem{node: 1, weightSum: nodeWeights[1]})

		for !queue.IsEmpty() {
			item, _ := queue.Dequeue()
			node, weightSum := item.node, item.weightSum

			if weightSum >= ansMap[goal] {
				continue
			}

			for _, next := range graph[node] {
				nextNode, nextWeight := next[0], next[1]

				ws := weightSum + nextWeight + nodeWeights[nextNode]
				if ws >= ansMap[goal] {
					continue
				}

				if nextNode == goal {
					ansMap[goal] = ws
					continue
				}

				queue.Enqueue(queueItem{node: nextNode, weightSum: ws})
			}
		}
	}

	for i := 2; i <= N; i++ {
		if i != 2 {
			fmt.Fprint(w, " ")
		}

		fmt.Fprint(w, ansMap[i])

		if i == N {
			fmt.Fprintln(w)
		}
	}
}

//////////////
// Libs    //
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
