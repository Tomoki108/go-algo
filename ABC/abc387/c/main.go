package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

//lint:ignore U1000 unused -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	L, R := read2Ints(r)

	// ans := countSnakeNum(L)

	// fmt.Fprintln(w, ans)

	ans := countSnakeNum(R) - countSnakeNum(L-1)

	fmt.Fprintln(w, ans)
}

func countSnakeNum(r int) int {
	digits := toDigits(r)

	// pos 現在何桁目まで埋めたか（left indexed, starts with 0）
	// firstNum その桁までに最初に登場したゼロでは無い数字 || ゼロ
	// strict digitsを超えないように次の桁の数字を選ぶ必要があるかどうか
	var digitDP func(pos, firstNum, strict int, digits []int) int

	digitDP = func(pos, firstNum, strict int, digits []int) int {
		fmt.Printf("pos: %d, firstNum: %d, strict: %d\n", pos, firstNum, strict)

		if pos == len(digits)-1 {
			return 1
		}

		var limit int
		if strict == 1 {
			limit = digits[pos+1]
		} else if firstNum != 0 {
			limit = firstNum - 1
		} else {
			limit = 9
		}
		fmt.Printf("limit: %d\n", limit)

		res := 0
		for nextDigit := 0; nextDigit <= limit; nextDigit++ {
			nextStrict := 0
			if strict == 1 && nextDigit == limit {
				nextStrict = 1
			}

			nextFirstNum := firstNum
			if firstNum == 0 && nextDigit != 0 {
				nextFirstNum = nextDigit
			}

			res += digitDP(pos+1, nextFirstNum, nextStrict, digits)
		}

		fmt.Println("return\n")
		return res
	}

	res := 0
	for firstDigit := 0; firstDigit <= digits[0]; firstDigit++ {
		strict := 0
		if firstDigit == digits[0] {
			strict = 1
		}

		res += digitDP(0, firstDigit, strict, digits)
	}

	return res
}

func toDigits(n int) []int {
	s := strconv.FormatInt(int64(n), 10)
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i] - '0')
	}
	return digits
}

//////////////
// Libs    //
/////////////

// O(n)
// slices.Reverce() と同じ（Goのバージョンが1.21以前だと使えないため）
func SlRev[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
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
//
//lint:ignore U1000 unused
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
