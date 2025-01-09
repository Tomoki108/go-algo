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

	N := readInt(r)
	S := readStr(r)
	Ss := strings.Split(S, "")

	Cs := readIntArr(r)

	zeroOneSeqCosts := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var expected string
		if i%2 == 0 {
			expected = "0"
		} else {
			expected = "1"
		}

		cost := 0
		if Ss[i] != expected {
			cost = Cs[i]
		}
		zeroOneSeqCosts = append(zeroOneSeqCosts, cost)
	}
	zeroOneSeqCostPrefSum := PrefixSum(zeroOneSeqCosts)

	oneZeroSeqCosts := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var expected string
		if i%2 == 0 {
			expected = "1"
		} else {
			expected = "0"
		}

		cost := 0
		if Ss[i] != expected {
			cost = Cs[i]
		}
		oneZeroSeqCosts = append(oneZeroSeqCosts, cost)
	}
	oneZeroSeqCostPrefSum := PrefixSum(oneZeroSeqCosts)

	ans := INT_MAX
	// iは、01と10を反転させる仕切りの位置。仕切りは、index iの直後にある。
	// 仕切りの左が01、右が10になるパターン１と、仕切りの左が10、右が01になるパターン２を考える。
	for i := 0; i <= N-2; i++ {
		cost1 := zeroOneSeqCostPrefSum[i+1] + oneZeroSeqCostPrefSum[N] - oneZeroSeqCostPrefSum[i+1]
		ans = min(ans, cost1)

		cost2 := oneZeroSeqCostPrefSum[i+1] + zeroOneSeqCostPrefSum[N] - zeroOneSeqCostPrefSum[i+1]
		ans = min(ans, cost2)
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// O(n)
// 一次元配列の累積和を返す（index0には0を入れる。）
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
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
