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

	N := readInt(r)
	S := readStr(r)
	T := readStr(r)

	Ss := strings.Split(S, "")
	Ss = append(Ss, ".")
	Ss = append(Ss, ".")
	Ts := strings.Split(T, "")
	Ts = append(Ts, ".")
	Ts = append(Ts, ".")

	visited := make(map[string]bool)

	q := NewQueue[status]()
	q.Enqueue(status{N, N + 1, Ss, 0})

	for !q.IsEmpty() {
		st, _ := q.Dequeue()

		if visited[st.str()] {
			continue
		}
		visited[st.str()] = true

		sis := st.swappableIndexs()
		for _, si := range sis {
			sl := st.swap(si[0], si[1])

			if strings.Join(sl, "") == strings.Join(Ts, "") {
				fmt.Fprintln(w, st.gen+1)
				return
			}

			q.Enqueue(status{si[0], si[1], sl, st.gen + 1})
		}
	}

	fmt.Fprintln(w, -1)
}

type status struct {
	i, j int      // 空マスのインデックス
	sl   []string // 盤面（空マス含む）
	gen  int      // 何手目か
}

func (s status) str() string {
	return strings.Join(s.sl, "")
}

func (s status) swappableIndexs() [][2]int {
	result := [][2]int{}
	for i := 0; i < len(s.sl)-1; i++ {
		if i != s.i && i != s.j && i+1 != s.i {
			result = append(result, [2]int{i, i + 1})
		}
	}

	return result
}

func (s status) swap(ni, nj int) []string {
	copySl := make([]string, len(s.sl))
	copy(copySl, s.sl)

	copySl[s.i], copySl[s.j] = copySl[ni], copySl[nj]
	copySl[ni], copySl[nj] = ".", "."

	return copySl
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

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
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
