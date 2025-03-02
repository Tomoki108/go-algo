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

	iarr := readIntArr(r)
	N, K, P := iarr[0], iarr[1], iarr[2]

	projects := make([][]int, N)
	for i := 0; i < N; i++ {
		projects[i] = readIntArr(r) // cost, A1, A2, ..., AK
	}

	genKey := func(params []int) string {
		return strings.Join(strings.Fields(fmt.Sprint(params)), "_")
	}

	ans := AscIntSearch(0, pow(10, 9), func(costLimt int) bool {
		memos := make(map[string]int, pow(P, K)) // key => cost

		var dp func(projectIdx int, current []int, currentCost int) bool
		dp = func(projectIdx int, current []int, currentCost int) bool {
			if projectIdx == N {
				return false
			}

			key := genKey(current)
			memoCost, ok := memos[key]
			if ok && memoCost <= currentCost {
				return false
			}

			p := projects[projectIdx]
			cost := p[0]

			// projectを実行しない場合
			ret := dp(projectIdx+1, current, currentCost)
			if ret {
				return true
			}

			// projectを実行する場合
			if newCost := currentCost + cost; newCost <= costLimt {
				ccurrent := make([]int, K)
				copy(ccurrent, current)

				achieved := true
				for i := 0; i < K; i++ {
					ccurrent[i] = min(ccurrent[i]+p[i+1], P)
					if ccurrent[i] != P {
						achieved = false
					}
				}

				if achieved {
					return true
				} else {
					memos[key] = newCost
					return dp(projectIdx+1, ccurrent, newCost)
				}
			}

			return false
		}

		return dp(0, make([]int, K), 0)
	})

	fmt.Println(ans) // 達成できない場合は元々-1が入っている
}

//////////////
// Libs    //
/////////////

// O(log (high-low))
// low, low+1, ..., highの範囲で条件を満たす最小の値を二分探索する
// low~highは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はlow-1を返す
func AscIntSearch(low, high int, f func(num int) bool) int {
	initialLow := low

	for low < high {
		// オーバーフローを防ぐための立式
		// 中央値はlow側に寄る
		mid := low + (high-low)/2
		if f(mid) {
			high = mid // 条件を満たす場合、よりlow側の範囲を探索
		} else {
			low = mid + 1 // 条件を満たさない場合、よりhigh側の範囲を探索
		}
	}

	// 最後に low(=high) が条件を満たしているかを確認
	if f(low) {
		return low
	}

	return initialLow - 1 // 条件を満たす値が見つからない場合
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

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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
	if exp == 0 {
		return 1
	}

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
