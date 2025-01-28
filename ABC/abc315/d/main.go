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

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	H, W := read2Ints(r)

	grid := readGrid(r, H)

	rowColorColSetMaps := make([]map[string]map[int]struct{}, H) // color -> [col1, col2, ...]
	colColorRowSetMaps := make([]map[string]map[int]struct{}, W) // color -> [row1, row2, ...]
	for row := 0; row < H; row++ {
		rowColorColSetMaps[row] = make(map[string]map[int]struct{})
		for col := 0; col < W; col++ {
			color := grid[row][col]
			if _, ok := rowColorColSetMaps[row][color]; !ok {
				rowColorColSetMaps[row][color] = make(map[int]struct{})
			}
			rowColorColSetMaps[row][color][col] = struct{}{}
		}
	}
	for col := 0; col < W; col++ {
		colColorRowSetMaps[col] = make(map[string]map[int]struct{})
		for row := 0; row < H; row++ {
			color := grid[row][col]
			if _, ok := colColorRowSetMaps[col][color]; !ok {
				colColorRowSetMaps[col][color] = make(map[int]struct{})
			}
			colColorRowSetMaps[col][color][row] = struct{}{}
		}
	}

	dump("rowColorSetMaps: %v\n", rowColorColSetMaps)
	dump("colColorSetMaps: %v\n", colColorRowSetMaps)

	q1 := NewQueue[qItem1]()
	q2 := NewQueue[qItem2]()

	changed := true
	for changed {
		changed = false

	Outer1:
		for row, rowColorSetMap := range rowColorColSetMaps {
			if len(rowColorSetMap) == 1 {
				for color, colSet := range rowColorSetMap { // only 1 iterate
					if len(colSet) < 2 {
						continue Outer1
					}

					qi := qItem1{
						color: color,
						cols:  mapKeys(colSet),
						row:   row,
					}
					q1.Enqueue(qi)

					delete(rowColorSetMap, color)
				}

				changed = true
			}
		}

	Outer2:
		for col, colColorSetMap := range colColorRowSetMaps {
			if len(colColorSetMap) == 1 {
				for color, rowSet := range colColorSetMap { // only 1 iterate
					if len(rowSet) < 2 {
						continue Outer2
					}

					qi := qItem2{
						color: color,
						rows:  mapKeys(rowSet),
						col:   col,
					}
					q2.Enqueue(qi)

					delete(colColorSetMap, color)
				}

				changed = true
			}
		}

		for !q1.IsEmpty() {
			qi, _ := q1.Dequeue()
			color := qi.color
			cols := qi.cols
			row := qi.row

			for _, col := range cols {
				if len(colColorRowSetMaps[col][color]) == 1 {
					delete(colColorRowSetMaps[col], color)
				} else {
					delete(colColorRowSetMaps[col][color], row)
				}
			}
		}
		for !q2.IsEmpty() {
			qi, _ := q2.Dequeue()
			color := qi.color
			rows := qi.rows
			col := qi.col

			for _, row := range rows {
				if len(rowColorColSetMaps[row][color]) == 1 {
					delete(rowColorColSetMaps[row], color)
				} else {
					delete(rowColorColSetMaps[row][color], col)
				}
			}
		}

		dump("\n")
		dump("rowColorSetMaps: %v\n", rowColorColSetMaps)
		dump("colColorSetMaps: %v\n", colColorRowSetMaps)
		dump("\n")
	}

	ans := 0
	for _, rowColorSetMap := range rowColorColSetMaps {
		for _, colSet := range rowColorSetMap {
			ans += len(colSet)
		}
	}

	dump("rowColorSetMaps: %v\n", rowColorColSetMaps)
	dump("colColorSetMaps: %v\n", colColorRowSetMaps)

	fmt.Println(ans)
}

type qItem1 struct {
	color string
	cols  []int
	row   int
}

type qItem2 struct {
	color string
	rows  []int
	col   int
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

func (q *Queue[T]) Clear() {
	q.list.Init()
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
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArr(r)
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

func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
