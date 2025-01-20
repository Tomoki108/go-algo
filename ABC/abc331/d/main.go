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

	N, Q := read2Ints(r)

	colors := readGrid(r, N)
	countGrid := createGrid(N, N, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if colors[i][j] == "B" {
				countGrid[i][j] = 1
			}
		}
	}
	countSumGrid := PrefixSum2D(countGrid)

	// h, wは右下隅の座標
	countInSquare := func(h, w int) int {
		h_q := (h + 1) / N
		h_rem := (h + 1) % N
		w_q := (w + 1) / N
		w_rem := (w + 1) % N
		fmt.Printf("h: %d, w: %d, h_q: %d, h_rem: %d, w_q: %d, w_rem: %d\n\n", h, w, h_q, h_rem, w_q, w_rem)

		ret := 0

		// add whole block count
		ret += h_q * w_q * countSumGrid[N][N]
		fmt.Printf("whole, h_q: %d, w_q: %d, count: %d\n", h_q, w_q, h_q*w_q*countSumGrid[N][N])
		// fmt.Printf("ret: %d\n\n", ret)

		// add vertical stic-out count
		{
			var h = h_rem
			var w int
			var count int
			if w_q < 1 {
				w = w_rem
				count = countSumGrid[h][w]
			} else {
				w = N
				count = countSumGrid[h][w]
				count *= w_q
			}

			fmt.Printf("vertical stick-out, h: %d, w: %d, count: %d\n", h, w, count)

			ret += count
		}

		// add horizontal stic-out count
		{
			var h int
			var w = w_rem
			var count int
			if h_q < 1 {
				h = h_rem
				count = countSumGrid[h][w]
			} else {
				h = N
				count = countSumGrid[h][w]
				count *= h_q
			}

			fmt.Printf("horizontal stick-out, h: %d, w: %d, count: %d\n", h, w, count)

			ret += count
		}

		// sub corner count
		{
			var h int
			var w int
			if h_rem > 0 && w_rem > 0 {
				h = h_rem
				w = w_rem
			} else {
				h = 0
				w = 0
			}

			fmt.Printf("corner, h: %d, w: %d, count: %d\n", h, w, countSumGrid[h][w])

			ret -= countSumGrid[h][w]
		}

		// fmt.Printf("ret: %d\n\n", ret)

		return ret
	}

	for i := 0; i < Q; i++ {
		iarr := readIntArr(r)
		A, B, C, D := iarr[0], iarr[1], iarr[2], iarr[3]
		A--
		B--

		ans := countInSquare(C, D) -
			countInSquare(C, B) -
			countInSquare(A, D) +
			countInSquare(A, B)

		fmt.Fprintln(w, ans)
	}
}

//////////////
// Libs    //
/////////////

// O(grid_size)
// 二次元累積和を返す（各次元のindex0には0を入れる。）
func PrefixSum2D(grid [][]int) [][]int {
	H := len(grid) + 1
	W := len(grid[0]) + 1

	sumGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		sumGrid[i] = make([]int, W)

		if i == 0 {
			continue
		}
		copy(sumGrid[i][1:], grid[i-1])
	}

	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			sumGrid[i][j] += sumGrid[i][j-1]
		}
	}
	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			sumGrid[i][j] += sumGrid[i-1][j]
		}
	}
	return sumGrid
}

// 二次元累積和から、任意の範囲の和を求める
// sumGridには、x, y, z方向に番兵（余分な空の一行）が含まれているものとする
// Lx, Rxは、その軸における範囲指定 => x方向には、Rxの累積和からLx-1の累積和を引く
func SumFrom2DPrefixSum(sumGrid [][]int, Lx, Rx, Ly, Ry int) int {
	Lx--
	Ly--

	// 包除原理
	result := sumGrid[Rx][Ry]

	result -= sumGrid[Lx][Ry]
	result -= sumGrid[Rx][Ly]

	result += sumGrid[Lx][Ly]

	return result
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
