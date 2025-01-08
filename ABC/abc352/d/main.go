package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"

	"github.com/emirpasic/gods/sets/treeset"
)

//lint:ignore U1000 unused 9223372036854775808, 19 digits, equiv 2^63
const INT_MAX = math.MaxInt

//lint:ignore U1000 unused -9223372036854775808, 19 digits, equiv -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// use gostl.set
func main() {
	defer w.Flush()

	N, K := read2Ints(r)
	Ps := readIntArr(r)

	if K == 1 {
		fmt.Fprintln(w, 0)
		return
	}

	type PN struct {
		P  int
		No int
	}
	PNs := make([]PN, 0, N)
	for i, P := range Ps {
		PNs = append(PNs, PN{P: P, No: i + 1})
	}
	sort.Slice(PNs, func(i, j int) bool {
		return PNs[i].P < PNs[j].P
	})

	ans := INT_MAX

	currentNos := set.New[int](comparator.IntComparator)
	for i := 0; i < K; i++ {
		currentNos.Insert(PNs[i].No)
	}

	right := K - 1
	left := 0
	for right < N {
		minNo := currentNos.First().Value()
		maxNo := currentNos.Last().Value()
		ans = min(ans, maxNo-minNo)

		right++
		left++
		if right != len(PNs) {
			currentNos.Erase(PNs[left-1].No)
			currentNos.Insert(PNs[right].No)
		}
	}

	fmt.Fprintln(w, ans)
}

// use gods.treeset
func alt() {
	defer w.Flush()

	N, K := read2Ints(r)
	Ps := readIntArr(r)

	if K == 1 {
		fmt.Fprintln(w, 0)
		return
	}

	type PN struct {
		P  int
		No int
	}
	PNs := make([]PN, 0, N)
	for i, P := range Ps {
		PNs = append(PNs, PN{P: P, No: i + 1})
	}
	sort.Slice(PNs, func(i, j int) bool {
		return PNs[i].P < PNs[j].P
	})

	ans := INT_MAX

	currentNos := treeset.NewWithIntComparator()
	for i := 0; i < K; i++ {
		currentNos.Add(PNs[i].No)
	}

	right := K - 1
	left := 0
	for right < N {
		minNo := First[int](currentNos)
		maxNo := Last[int](currentNos)
		ans = min(ans, maxNo-minNo)

		right++
		left++
		if right != len(PNs) {
			currentNos.Remove(PNs[left-1].No)
			currentNos.Add(PNs[right].No)
		}
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// O(1)
func First[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.Begin()
	it.Next()

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

// O(1)
func Last[T any](set *treeset.Set) T {
	it := set.Iterator()
	it.End()
	it.Prev()

	val, ok := it.Value().(T)
	if !ok {
		panic("Type assertion failed")
	}

	return val
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
//
//lint:ignore U1000 unused
func readStr(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
//
//lint:ignore U1000 unused
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
//
//lint:ignore U1000 unused
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
//
//lint:ignore U1000 unused
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
//
//lint:ignore U1000 unused
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
//
//lint:ignore U1000 unused
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身をスペース区切りなしで出力する
//
//lint:ignore U1000 unused
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
//
//lint:ignore U1000 unused
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

//lint:ignore U1000 unused
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
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
