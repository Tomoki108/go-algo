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

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// タイルを通る：現在中（枠線を含まない）にいるタイルから、別のタイルの中に入った時に、元いたタイルを通ったとみなす。
func main() {
	defer w.Flush()

	sx, sy := read2Ints(r)
	Sx, Sy := float64(sx)+0.5, float64(sy)+0.5

	tx, ty := read2Ints(r)
	Tx, Ty := float64(tx)+0.5, float64(ty)+0.5

	var goRight bool
	if Sx < Tx {
		goRight = true
	}
	// ゴールのx座標が同じまたは前にある場合、右にはいかない。

	var goLeft bool
	if Sx > Tx {
		goLeft = true
	}
	// ゴールのx座標が同じまたは後ろにある場合、左にはいかない。

	var goUp bool
	if Sy < Ty {
		goUp = true
	}
	// ゴールのy座標が同じまたは前にある場合、上にはいかない。

	var goDown bool
	if Sy > Ty {
		goDown = true
	}
	// ゴールのy座標が同じまたは後ろにある場合、下にはいかない。

	q := NewQueue[qItem]()
	q.Enqueue(qItem{Sx, Sy, 0})

	prevX := Sx
	prevY := Sy
	cost := 0
	for !q.IsEmpty() {
		item, _ := q.Dequeue()
		x := item.x
		y := item.y
		cost := item.cost

		if x == Tx && y == Ty {
			break
		}

		adjs := getAdjacents(x, y, goRight, goLeft, goUp, goDown)
		for _, adj := range adjs {
			c := cost
			if !withinSameTile(prevX, prevY, x, y) {
				c++
			}

			q.Enqueue(qItem{adj[0], adj[1], c})
		}
	}

	fmt.Fprintln(w, cost)
}

type qItem struct {
	x, y float64
	cost int
}

// 5.5, 0.5 =>  math.Ceil(5.5), math.Ceil(0.5)の正方形に属する =>（5, 0）
// 4.5. 0.5 => (4, 0)
// xが小さい方の正方形の、xyの和が偶数かつ、xの差が1である場合に、同じタイルに属する
func withinSameTile(x1, y1, x2, y2 float64) bool {
	i1, j1 := int(math.Ceil(x1)), int(math.Ceil(float64(y1)))
	i2, j2 := int(math.Ceil(x2)), int(math.Ceil(float64(y2)))

	if j1 != j2 {
		return false
	}

	// 同じx軸（全く同じ座標）は考慮しない
	if i1 > i2 {
		return i1-i2 == 1 && i2+j2%2 == 0
	} else {
		return i2-i1 == 1 && i1+j1%2 == 0
	}
}

func getAdjacents(x, y float64, goRight, goLeft, goUp, goDown bool) [][2]float64 {
	right := [2]float64{x + 1, y}
	left := [2]float64{x - 1, y}
	up := [2]float64{x, y + 1}
	down := [2]float64{x, y - 1}

	adjacents := [][2]float64{}
	if goRight {
		adjacents = append(adjacents, right)
	}
	if goLeft {
		adjacents = append(adjacents, left)
	}
	if goUp {
		adjacents = append(adjacents, up)
	}
	if goDown {
		adjacents = append(adjacents, down)
	}

	return adjacents
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
