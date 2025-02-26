package main

import (
	"bufio"
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

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, Q := read2Ints(r)
	S := readStr(r)
	Ss := strings.Split(S, "")

	ones := make([]int, N)
	twos := make([]int, N)
	slashIndexes := make([]int, 0, N)
	for i := 0; i < N; i++ {
		if Ss[i] == "1" {
			ones[i] = 1
		} else if Ss[i] == "2" {
			twos[i] = 1
		} else {
			slashIndexes = append(slashIndexes, i)
		}
	}
	psumOne := PrefixSum(ones)
	psumTwo := PrefixSum(twos)

	type SlashInfo struct {
		idx       int
		leftOnes  int
		rightTwos int
	}
	slashInfos := make([]SlashInfo, 0, len(slashIndexes))
	for _, slIdx := range slashIndexes {
		leftOnes := psumOne[slIdx+1]
		rightTwos := psumTwo[N] - psumTwo[slIdx+1]
		slashInfos = append(slashInfos, SlashInfo{slIdx, leftOnes, rightTwos})
	}

	for i := 0; i < Q; i++ {
		L, R := read2Ints(r)
		L--
		R--

		toSearch := RangeSearch(slashInfos, func(info SlashInfo) int { return info.idx }, L, R)
		if toSearch == nil {
			fmt.Fprintln(w, 0)
			continue
		}

		leftBuff := psumOne[L]
		rightBuff := psumTwo[N] - psumTwo[R+1]
		maxHalfCnt := toSearch[len(toSearch)-1].leftOnes - leftBuff

		halfCnt := DescIntSearch(maxHalfCnt, 0, func(halfCnt int) bool {
			toSearchIdx := sort.Search(len(toSearch), func(idx int) bool {
				return toSearch[idx].leftOnes-leftBuff >= halfCnt
			})

			rightTwos := toSearch[toSearchIdx].rightTwos - rightBuff
			return rightTwos >= halfCnt
		})
		dump("halfCnt: %d\n\n", halfCnt)

		fmt.Fprintln(w, halfCnt*2+1)
	}
}

//////////////
// Libs    //
/////////////

// O(log |sl|)
// 単調増加性を満た任意型のスライスから、任意型が持つ単調増加性を満たす値がl以上、r以下の範囲を返す
func RangeSearch[U any](sl []U, valuer func(item U) int, l, r int) []U {
	idx1 := sort.Search(len(sl), func(i int) bool { return valuer(sl[i]) >= l })
	if idx1 == len(sl) {
		return nil
	}
	idx2 := sort.Search(len(sl), func(i int) bool { val := valuer(sl[i]); return val > r })
	if idx2 == idx1 {
		return nil
	}
	return sl[idx1:idx2]
}

// O(log(high-low))
// high, high-1, ..., lowの範囲で条件を満たす最大の値を二分探索する
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

// O(n)
// 一次元累積和を返す（index0には0を入れる。）
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
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
func readIntGrid(r *bufio.Reader, height int, withSpace bool) [][]int {
	if withSpace {
		grid := make([][]int, height)
		for i := 0; i < height; i++ {
			grid[i] = readIntArr(r)
		}
		return grid
	}

	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		strs := strings.Split(str, "")

		grid[i] = make([]int, len(strs))
		for j, s := range strs {
			grid[i][j], _ = strconv.Atoi(s)
		}
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

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

func updateToMin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func updateToMax(a *int, b int) {
	if *a < b {
		*a = b
	}
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

// NOTE: ループの中で使うとわずかに遅くなることに注意
func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
