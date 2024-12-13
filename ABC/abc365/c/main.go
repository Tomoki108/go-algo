package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)

	As := readIntArr(r)
	sort.Slice(As, func(i, j int) bool {
		return As[i] < As[j]
	})
	prefSum := PrefixSum(As)
	totalSum := prefSum[len(prefSum)-1]

	if totalSum <= M {
		fmt.Println("infinite")
		return
	}

	minX := M / N
	maxX := M

	f := func(x int) bool {
		xIdx := sort.Search(len(As), func(i int) bool {
			return As[i] >= x
		})

		cost := prefSum[xIdx]
		cost += x * (len(As) - xIdx)

		return cost <= M
	}
	x := DescIntSearch(maxX, minX, f)

	fmt.Fprintln(w, x)
}

//////////////
// Libs    //
/////////////

// 一次元配列の累積和を返す
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
}

// high, high-1, ..., lowの範囲で条件を満たす最小の値を二分探索する
// high~lowは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はhigh+1を返す
func DescIntSearch(high, low int, f func(num int) bool) int {
	for low < high {
		// オーバーフローを防ぐための式.
		// 中央値はhigh側に寄る（+1しているため）
		mid := low + (high-low+1)/2
		if f(mid) {
			low = mid // 条件を満たす場合、よりhigh側の範囲を探索
		} else {
			high = mid - 1 // 条件を満たさない場合、よりlow側の範囲を探索
		}
	}

	// 最後に high(=low) が条件を満たしているかを確認
	if f(high) {
		return high
	}

	return high + 1 // 条件を満たす値が見つからない場合
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
