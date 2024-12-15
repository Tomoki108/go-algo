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

	iarr := readIntArr(r)
	m := make(map[string]int, 5)

	idx := 0
	for i := 'A'; i <= 'E'; i++ {
		m[string(i)] = iarr[idx]
		idx++
	}

	patterns := make([][]string, 0, 31)
	for i := 5; i >= 1; i-- {
		patterns = append(patterns, PickN([]string{}, []string{"A", "B", "C", "D", "E"}, i)...)
	}

	// bit全探索でpatternsを列挙する方法
	// abcde := []string{"A", "B", "C", "D", "E"}
	// bitMax := 1<<5 - 1
	// for bit := bitMax; bit > 0; bit-- {
	// 	pattern := []string{}
	// 	for i := 1; i <= 5; i++ {
	// 		if IsBitPop(uint64(bit), i) {
	// 			pattern = append(pattern, abcde[i-1])
	// 		}
	// 	}
	// 	patterns = append(patterns, pattern)
	// }

	sort.Slice(patterns, func(i, j int) bool {
		p1 := patterns[i]
		p2 := patterns[j]

		score1 := 0
		score2 := 0
		for i := 0; i < len(p1); i++ {
			char := p1[i]
			score1 += m[char]
		}
		for i := 0; i < len(p2); i++ {
			char := p2[i]
			score2 += m[char]
		}

		length := min(len(p1), len(p2))
		if score1 == score2 {
			for i := 0; i < length; i++ {
				if i == length {
					if len(p1) < len(p2) {
						return true
					} else if len(p1) > len(p2) {
						return false
					}
				}

				char := p1[i]
				char2 := p2[i]

				if char < char2 {
					return true
				} else if char > char2 {
					return false
				}
			}
		}

		return score1 > score2
	})

	for _, p := range patterns {
		fmt.Fprintln(w, strings.Join(p, ""))
	}
}

//////////////
// Libs    //
/////////////

// O(nCr) n: len(options), r: n
// optionsから N個選ぶ組み合わせを全列挙する
// optionsにはソート済みかつ要素に重複のないスライスを渡すこと（戻り値が辞書順になり、重複組み合わせも排除される）
func PickN[T comparable](current, options []T, n int) [][]T {
	var results [][]T

	if n == 0 {
		return [][]T{current}
	}

	for i, o := range options {
		newCurrent := append([]T{}, current...)
		newCurrent = append(newCurrent, o)
		newOptions := append([]T{}, options[i+1:]...)

		results = append(results, PickN(newCurrent, newOptions, n-1)...)
	}

	return results
}

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
