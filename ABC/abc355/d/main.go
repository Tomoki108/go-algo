package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
const intMax = 1 << 62

//lint:ignore U1000 unused
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// 全組み合わせから、重ならない区間の数を引く方法。
// Aの右端 < Bの左端　となる区間A,Bは重ならない。
// 右端、左端を数直線上にならべ、ある右端の前に何個左端の座標があるかを順に数えていくと、重ならない区間の組み合わせの数が分かる。
func main() {
	defer w.Flush()

	N := readInt(r)

	type xWithType struct {
		x int // x座標
		t int // 0: start, 1: end
	}

	xs := make([]xWithType, 0, N)
	for i := 0; i < N; i++ {
		l, r := read2Ints(r)
		xs = append(xs, xWithType{l, 0})
		xs = append(xs, xWithType{r, 1})
	}

	sort.Slice(xs, func(i, j int) bool {
		if xs[i].x == xs[j].x {
			return xs[i].t < xs[j].t
		}
		return xs[i].x < xs[j].x
	})

	ans := N * (N - 1) / 2
	numOfEndsPassed := 0
	for i := 0; i < len(xs); i++ {
		if xs[i].t == 1 {
			numOfEndsPassed++
		} else {
			ans -= numOfEndsPassed
		}
	}

	fmt.Fprintln(w, ans)
}

// 区間スライスと始点スライスを二つ作り二分探索する方法。
func alt() {
	defer w.Flush()

	N := readInt(r)

	type segment struct {
		l, r int
	}

	segments := make([]segment, 0, N)
	starts := make([]int, 0, N)

	for i := 0; i < N; i++ {
		l, r := read2Ints(r)
		segments = append(segments, segment{l, r})
		starts = append(starts, l)
	}

	sort.Slice(segments, func(i, j int) bool {
		if segments[i].l == segments[j].l {
			return segments[i].r < segments[j].r
		}
		return segments[i].l < segments[j].l
	})
	sort.Slice(starts, func(i, j int) bool {
		return starts[i] > starts[j]
	})

	ans := 0
	for _, seg := range segments {
		starts = starts[:len(starts)-1]

		idx := sort.Search(len(starts), func(j int) bool {
			return starts[j] <= seg.r
		})

		ans += len(starts) - idx
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

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
