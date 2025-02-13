package main

import (
	"bufio"
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

	readInt(r)
	As := readIntArr(r)

	_, nodeMap := CreateBDList(As)

	Q := readInt(r)
	for i := 0; i < Q; i++ {
		iarr := readIntArr(r)
		q := iarr[0]
		if q == 1 {
			x, y := iarr[1], iarr[2]
			xNode := nodeMap[x]
			xNode.InsertAfter(y)
			nodeMap[y] = xNode.next
		} else {
			x := iarr[1]
			xNode := nodeMap[x]
			xNode.Remove()
			delete(nodeMap, x)
		}
	}

	var head *NodeBD[int]
	for _, node := range nodeMap {
		head = node.Head()
		break
	}

	current := head
	for current != nil {
		if current == head {
			fmt.Fprint(w, current.val)
		} else {
			fmt.Fprint(w, " ", current.val)
		}
		current = current.next
	}
	fmt.Fprintln(w)
}

//////////////
// Libs    //
/////////////

type NodeBD[T comparable] struct {
	val  T
	prev *NodeBD[T]
	next *NodeBD[T]
}

func (n *NodeBD[T]) Head() *NodeBD[T] {
	for n.prev != nil {
		n = n.prev
	}
	return n
}

func (n *NodeBD[T]) Tail() *NodeBD[T] {
	for n.next != nil {
		n = n.next
	}
	return n
}

func (n *NodeBD[T]) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

func (n *NodeBD[T]) InsertAfter(val T) {
	node := &NodeBD[T]{val: val, prev: n, next: n.next}
	if n.next != nil {
		n.next.prev = node
	}
	n.next = node
}

func CreateBDList[T comparable](sl []T) (head *NodeBD[T], nodeMap map[T]*NodeBD[T]) {
	if len(sl) == 0 {
		return nil, nil
	}

	head = &NodeBD[T]{val: sl[0]}
	nodeMap = make(map[T]*NodeBD[T])
	nodeMap[sl[0]] = head

	prev := head
	for i := 1; i < len(sl); i++ {
		node := &NodeBD[T]{val: sl[i], prev: prev}
		prev.next = node
		prev = node
		nodeMap[sl[i]] = node
	}
	return head, nodeMap
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
func createGrid[T any](r *bufio.Reader, height, width int) [][]T {
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
