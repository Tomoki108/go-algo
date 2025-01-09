package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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
	a, b, C := iarr[0], iarr[1], uint64(iarr[2])

	C_pc := bits.OnesCount64(C)
	pc_diff := abs(a - b)

	if C_pc%2 != pc_diff%2 {
		fmt.Fprintln(w, -1)
		return
	}
	if pc_diff > C_pc {
		fmt.Fprintln(w, -1)
		return
	}
	if a+b < C_pc {
		fmt.Fprintln(w, -1)
		return
	}

	if C_pc == 0 {
		if a != b {
			fmt.Fprintln(w, -1)
			return
		} else {
			var X uint64
			var Y uint64
			for exp := 0; exp < a; exp++ {
				X += uint64(pow(2, exp))
				Y += uint64(pow(2, exp))
			}

			fmt.Fprintln(w, X, Y)
			return
		}
	}

	var X uint64
	var Y uint64
	if a > b {
		X = C
		pc_X := C_pc

		passedBits := 0
		for exp := 0; exp < 60; exp++ {
			if IsBitPop(X, exp+1) {
				X -= uint64(pow(2, exp))
				Y += uint64(pow(2, exp))
				passedBits++
				if pc_X-passedBits*2 == pc_diff {
					break
				}
			}
		}

		pc_X = bits.OnesCount64(X)
		toAdd := a - pc_X
		if toAdd > 0 {
			for exp := 0; exp < 60; exp++ {
				if !IsBitPop(X, exp+1) && !IsBitPop(Y, exp+1) {
					X += uint64(pow(2, exp))
					Y += uint64(pow(2, exp))
					toAdd--
					if toAdd == 0 {
						break
					}
				}
			}
		}

		if toAdd != 0 {
			fmt.Fprintln(w, -1)
			return
		}
	} else {
		Y = C
		pc_Y := C_pc

		passedBits := 0
		for exp := 0; exp < 60; exp++ {
			if IsBitPop(Y, exp+1) {
				Y -= uint64(pow(2, exp))
				X += uint64(pow(2, exp))
				passedBits++
				if pc_Y-passedBits*2 == pc_diff {
					break
				}
			}
		}

		pc_Y = bits.OnesCount64(Y)
		toAdd := b - pc_Y
		if toAdd != 0 {
			for exp := 0; exp < 60; exp++ {
				if !IsBitPop(X, exp+1) && !IsBitPop(Y, exp+1) {
					X += uint64(pow(2, exp))
					Y += uint64(pow(2, exp))
					toAdd--
					if toAdd == 0 {
						break
					}
				}
			}
		}

		if toAdd > 0 {
			fmt.Fprintln(w, -1)
			return
		}
	}

	fmt.Fprintln(w, X, Y)
}

//////////////
// Libs    //
/////////////

// k桁目のビットが1かどうかを判定（一番右を１桁目とする）
func IsBitPop(num uint64, k int) bool {
	// 1 << (k - 1)はビットマスク。1をk - 1桁左にシフトすることで、k桁目のみが1で他の桁が0の二進数を作る。
	// numとビットマスクの論理積（各桁について、numとビットマスクが両方trueならtrue）を作り、その結果が0でないかどうかで判定できる
	return (num & (1 << (k - 1))) != 0
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
