package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
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
	grid := readGrid(r, N)

	obstacles := 0
	ps := make([]Coordinate, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == "P" {
				ps = append(ps, Coordinate{i, j})
			} else if grid[i][j] == "#" {
				obstacles++
			}
		}
	}

	memoSize := CombinationNum(N*N-obstacles, 2)
	// fmt.Fprintln(w, CombinationNum(3600, 2)) // 6468600

	occurredMap := make(map[string]bool, memoSize)
	genKey := func(p1, p2 Coordinate) string {
		p1, p2 = sort(p1, p2)
		return fmt.Sprintf("%d %d %d %d", p1.h, p1.w, p2.h, p2.w)
	}

	q := NewQueue[qItem]()
	q.Enqueue(qItem{ps[0], ps[1], 0})
	occurredMap[genKey(ps[0], ps[1])] = true

	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		p1, p2 := item.p1, item.p2

		if p1 == p2 {
			fmt.Fprintln(w, item.depth)
			return
		}

		// 上下左右
		for _, delta := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			np1 := Coordinate{p1.h + delta[0], p1.w + delta[1]}
			np2 := Coordinate{p2.h + delta[0], p2.w + delta[1]}

			if !(np1.IsValid(N, N) && grid[np1.h][np1.w] != "#") {
				np1 = p1
			}
			if !(np2.IsValid(N, N) && grid[np2.h][np2.w] != "#") {
				np2 = p2
			}

			if !(p1 == np1 && p2 == np2) {
				key := genKey(np1, np2)
				if !occurredMap[key] {
					q.Enqueue(qItem{np1, np2, item.depth + 1})
					occurredMap[key] = true
				}
			}
		}
	}

	fmt.Fprintln(w, -1)
}

type qItem struct {
	p1, p2 Coordinate
	depth  int
}

//////////////
// Libs    //
/////////////

type Coordinate struct {
	h, w int // 0-indexed
}

func sort(one, another Coordinate) (Coordinate, Coordinate) {
	if one.h > another.h {
		return another, one
	} else if one.h == another.h {
		if one.w > another.w {
			return another, one
		}
		return one, another
	}
	return one, another
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

// O(r)
// nCrの計算
// (n * (n-1) ... * (n-r+1)) / r!
func CombinationNum(n, r int) int {
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

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
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
