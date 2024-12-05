package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// 「Rまでの人口の累積和」 - 「L-1までの人口の累積和」を求める
func main() {
	defer w.Flush()

	N := readInt(r)

	Xs := readIntArr(r)
	Ps := readIntArr(r)

	populationCSum := make(map[int]int, N) // 座標に対する人口の累積和（cumulative sum）
	csum := 0
	for i := 0; i < N; i++ {
		x := Xs[i]
		p := Ps[i]

		csum += p
		populationCSum[x] = csum
	}

	// fmt.Printf("populationCSum: %v\n", populationCSum)hb

	Q := readInt(r)
	for i := 0; i < Q; i++ {
		// [L, R]の区間の人口を求めたい
		L, R := read2Ints(r)

		var lCsum, rCsum int

		lXsIndex := sort.Search(N, func(i int) bool { return Xs[i] >= L })
		if lXsIndex == N { // L以上の座標にある村がなければ、該当範囲の人口は0
			fmt.Fprintln(w, 0)
			continue
		}

		if lXsIndex == 0 {
			lCsum = 0
		} else {
			lCsum = populationCSum[Xs[lXsIndex-1]]
		}

		rXsIndex := sort.Search(N, func(i int) bool { return Xs[i] >= R })
		if rXsIndex == N { // R以上の座標にある村がなければ、全村の人口の累積和を使う
			rCsum = csum
		} else if Xs[rXsIndex] == R || rXsIndex == 0 {
			rCsum = populationCSum[Xs[rXsIndex]]
		} else {
			rCsum = populationCSum[Xs[rXsIndex-1]]
		}

		// fmt.Printf("lCsum: %d, rCsum: %d\n", lCsum, rCsum)

		// [lx, rx]の区間の人口を求める
		sum := rCsum - lCsum
		fmt.Fprintln(w, sum)
	}
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
