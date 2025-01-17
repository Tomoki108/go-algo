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
	N--

	s := strconv.FormatInt(int64(N), 5)
	ss := strings.Split(s, "")

	m := map[string]string{
		"0": "0",
		"1": "2",
		"2": "4",
		"3": "6",
		"4": "8",
	}

	for _, v := range ss {
		fmt.Fprint(w, m[v])
	}
	fmt.Fprintln(w)
}

func alt() {
	defer w.Flush()

	N := readInt(r)

	genKey := func(pos int, strict bool) string {
		return fmt.Sprintf("%d-%v", pos, strict)
	}

	// var digits []int
	// var memos map[string]int

	// pos: pos桁目まで現在埋まっている。（左始まりの、0-indexed）
	// strict: 上限を気にして列挙する必要があるかどうか。（posより前の桁が全て上限の同一の桁と一致している。）
	// return: 現在の状態に合致する数のパターン数。
	var digitDP func(pos int, strict bool, memos map[string]int, digits []int) int
	digitDP = func(pos int, strict bool, memos map[string]int, digits []int) int {
		memo, exist := memos[genKey(pos, strict)]
		if exist {
			return memo
		}

		if pos == len(digits)-1 {
			return 1
		}

		limit := 8
		if strict {
			limit = digits[pos+1]
		}

		count := 0
		for i := 0; i <= limit; i++ {
			if i%2 != 0 {
				continue
			}

			newStrict := strict && i == limit

			count += digitDP(pos+1, newStrict, memos, digits)
		}

		memos[genKey(pos, strict)] = count
		return count
	}

	minNum := 0
	maxNum := INT_MAX
	ans := AscIntSearch(minNum, maxNum, func(num int) bool {
		memos := make(map[string]int)
		digits := ToDigits(num)

		count := 0
		limit := digits[0]

		for i := 0; i <= limit; i++ {
			if i%2 != 0 {
				continue
			}
			count += digitDP(0, digits[0] == i, memos, digits)
		}

		return count >= N
	})

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// O(log(high-low))
// low, low+1, ..., highの範囲で条件を満たす最小の値を二分探索する
// low~highは条件に対して単調増加性を満たす必要がある
// 条件を満たす値が見つからない場合はlow-1を返す
func AscIntSearch(low, high int, f func(num int) bool) int {
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

	return low - 1 // 条件を満たす値が見つからない場合
}

// O(n) n: numの桁数
// numの各桁の数字を返す
func ToDigits(n int) []int {
	s := strconv.FormatInt(int64(n), 10)
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i] - '0') // （'x'に対応する数字 - '0'に対応する数字）のruneの数字 = x as int
	}
	return digits
}

// O(n) n: numの桁数
// numの桁数を返す
func GetDigists(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	return digits
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
func createGrid[T any](height, width int) [][]T {
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
