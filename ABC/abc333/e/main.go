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

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N := readInt(r)

	txs := make([][2]int, 0, N)
	for i := 0; i < N; i++ {
		t, x := read2Ints(r)
		txs = append(txs, [2]int{t, x})
	}

	ansMap := make(map[int]int, N)           // potion found idx => take or not
	stackMap := make(map[int]*Stack[int], N) // type => stack of potion found idx
	for i, tx := range txs {
		t, x := tx[0], tx[1]

		if t == 1 {
			if _, ok := stackMap[x]; !ok {
				stackMap[x] = NewStack[int]()
			}
			stackMap[x].Push(i)
			ansMap[i] = 0
		} else {
			stack, ok := stackMap[x]
			if !ok {
				fmt.Fprintln(w, -1)
				return
			}

			idx, ok := stack.Pop()
			if !ok {
				fmt.Fprintln(w, -1)
				return
			}

			ansMap[idx] = 1
		}
	}

	currentK := 0
	maxK := 0
	for i, tx := range txs {
		t, _ := tx[0], tx[1]
		if t == 1 {
			if ansMap[i] == 1 {
				currentK++
				maxK = max(maxK, currentK)
			}
		} else {
			currentK--
		}
	}
	fmt.Fprintln(w, maxK)

	mKeys := mapKeys(ansMap)
	sort.Ints(mKeys)
	for i, mKey := range mKeys {
		fmt.Fprint(w, ansMap[mKey])

		if i == len(mKeys)-1 {
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, " ")
		}
	}
}

//////////////
// Libs    //
/////////////

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

func (s *Stack[T]) Clear() {
	s.list.Init()
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

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

//////////////
// Debug   //
/////////////

var dumpFlag bool

func init() {
	args := os.Args
	dumpFlag = len(args) > 1 && args[1] == "-dump"
}

// NOTE: ループの中で使うとわずかに遅くなることに注意
func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
