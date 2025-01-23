package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/rbtree"
	"github.com/liyue201/gostl/utils/comparator"
)

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)
	As := readIntArr(r)

	scoreMap := make(map[int]int, N) // person index => score
	scoreSet := NewIntMultiSetAsc()

	remAsMap := make(map[int]*MultiSet[int], N) // person index => remaining As

	grid := readGrid(r, N)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] == "o" {
				scoreMap[i] += As[j]
			} else {
				if _, ok := remAsMap[i]; !ok {
					remAsMap[i] = NewIntMultiSetAsc()
				}
				remAsMap[i].Insert(As[j])
			}
		}
		scoreMap[i] += i + 1 // ボーナス得点
		scoreSet.Insert(scoreMap[i])
	}

	// 降順ソート
	sort.Slice(As, func(i, j int) bool {
		return As[i] > As[j]
	})

	for i := 0; i < N; i++ {
		score := scoreMap[i]

		scoreSet.Erase(score)
		highestScore := scoreSet.Last().Key()
		scoreSet.Insert(score)

		remAs := remAsMap[i]

		ans := 0
		for score <= highestScore {
			solve := remAs.Last().Key()
			remAs.Erase(solve)

			score += solve
			ans++
		}

		fmt.Fprintln(w, ans)
	}
}

//////////////
// Libs    //
/////////////

func NewIntMultiSetAsc() *MultiSet[int] {
	return NewMultiSet[int](comparator.IntComparator)
}

func NewIntMultiSetDesc() *MultiSet[int] {
	return NewMultiSet[int](comparator.Reverse(comparator.IntComparator))
}

// gostlのNativeのMultiSetは、Erace()が同一の値を全て削除してしまうためこちらを使う。
type MultiSet[T any] struct {
	tree *rbtree.RbTree[T, int]
}

// O(1)
func NewMultiSet[T any](compare comparator.Comparator[T]) *MultiSet[T] {
	return &MultiSet[T]{
		tree: rbtree.New[T, int](compare),
	}
}

// O(log n)
func (ms *MultiSet[T]) Insert(value T) {
	if node := ms.tree.FindNode(value); node != nil {
		node.SetValue(node.Value() + 1)
	} else {
		ms.tree.Insert(value, 1)
	}
}

// Erase は、要素を1つだけ削除します。（カウントが1の場合はキーごと削除）
// O(log n)
func (ms *MultiSet[T]) Erase(value T) {
	if node := ms.tree.FindNode(value); node != nil {
		count := node.Value()
		if count > 1 {
			node.SetValue(count - 1)
		} else {
			ms.tree.Delete(node)
			return
		}
	}
}

// Count は、指定した要素の出現回数を返します。
// O(log n)
func (ms *MultiSet[T]) Count(value T) int {
	if node := ms.tree.FindNode(value); node != nil {
		return node.Value()
	}
	return 0
}

// Values は、マルチセット内の全ての要素をソートされた順序（比較関数に準ずる）で重複も含めてスライスとして返します。
// O(n + k)
//   - n はユニークなキーの数
//   - k は要素の総数（重複含む）
func (ms *MultiSet[T]) Values() []T {
	var result []T
	it := ms.tree.IterFirst()
	for it.IsValid() {
		key, count := it.Key(), it.Value()
		for i := 0; i < count; i++ {
			result = append(result, key)
		}

		it = it.Next().(*rbtree.RbTreeIterator[T, int])
	}
	return result
}

// Size は、マルチセットに含まれる要素の総数（重複含む）を返します。
// O(n)
//   - n はユニークなキーの数（RBTree 全体を走査）
func (ms *MultiSet[T]) Size() int {
	total := 0
	it := ms.tree.IterFirst()
	for it.IsValid() {
		_, count := it.Key(), it.Value()
		total += count

		it = it.Next().(*rbtree.RbTreeIterator[T, int])
	}
	return total
}

func (ms *MultiSet[T]) First() *rbtree.RbTreeIterator[T, int] {
	return ms.tree.IterFirst()
}

func (ms *MultiSet[T]) Last() *rbtree.RbTreeIterator[T, int] {
	return ms.tree.IterLast()
}

// Clear は、マルチセットを空にします。
// O(1) （実装上、要素数に依存しないクリア処理を行うため）
func (ms *MultiSet[T]) Clear() {
	ms.tree.Clear()
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

func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
