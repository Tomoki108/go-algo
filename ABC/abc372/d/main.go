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

	N := readInt(r)
	Hs := readIntArr(r)

	ans := make([]int, N)

	// 現在のビルから見えるビルの高さ（必然的に単調増加（末尾=>先頭））
	visibleHights := NewStack[int]()
	for i := N - 1; 0 <= i; i-- {
		if i == N-1 {
			ans[i] = 0
			continue
		}

		// 現在のビルの一つ後ろのビルをスタックに追加する。
		// その前にスタックから一つづつ取り出し、一つ後ろのビルより低いビルを取り除く
		for visibleHights.Len() > 0 {
			last, _ := visibleHights.Peek()
			if last > Hs[i+1] {
				break
			}

			visibleHights.Pop()
		}

		visibleHights.Push(Hs[i+1])
		ans[i] = visibleHights.Len()
	}

	writeSlice(w, ans)
}

type Stack[T any] struct {
	list *list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: list.New(),
	}
}

func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	back := s.list.Back()
	if back == nil {
		var zero T
		return zero, false
	}
	s.list.Remove(back)
	return back.Value.(T), true
}

// Peek returns the back element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	back := s.list.Back()
	if back == nil {
		var zero T
		return zero, false
	}
	return back.Value.(T), true
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

//////////////
// Hepers  //
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

// slices.Reverce() （Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func slReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
