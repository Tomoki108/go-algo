package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// 何個のABC文字列が含まれているかをカウントしておき、クエリのたびに周辺のABC文字列の個数の変動をカウントする方法
func main() {
	defer w.Flush()

	N, Q := read2Ints(r) // len(S), len(Queries)
	S := readStr(r)      // 文字列

	ss := strings.Split(S, "")

	count := 0
	for i := 0; i < N-2; i++ {
		if ss[i] == "A" && ss[i+1] == "B" && ss[i+2] == "C" {
			count++
		}
	}

	for i := 0; i < Q; i++ {
		sarr := readStrArr(r)
		XS := sarr[0]
		X, _ := strconv.Atoi(XS)
		Xi := X - 1
		C := sarr[1]

		p1Exists := Xi <= N-3 && ss[Xi] == "A" && ss[Xi+1] == "B" && ss[Xi+2] == "C"
		p2Exists := Xi <= N-2 && Xi >= 1 && ss[Xi-1] == "A" && ss[Xi] == "B" && ss[Xi+1] == "C"
		p3Exists := Xi >= 2 && ss[Xi-2] == "A" && ss[Xi-1] == "B" && ss[Xi] == "C"

		if p1Exists {
			if C == "A" {
				fmt.Fprintln(w, count)
				continue
			} else {
				count--
			}
		}

		if p2Exists {
			if C == "B" {
				fmt.Fprintln(w, count)
				continue
			} else {
				count--
			}
		}

		if p3Exists {
			if C == "C" {
				fmt.Fprintln(w, count)
				continue
			} else {
				count--
			}
		}

		ss[Xi] = C
		p1Exists = Xi <= N-3 && ss[Xi] == "A" && ss[Xi+1] == "B" && ss[Xi+2] == "C"
		p2Exists = Xi <= N-2 && Xi >= 1 && ss[Xi-1] == "A" && ss[Xi] == "B" && ss[Xi+1] == "C"
		p3Exists = Xi >= 2 && ss[Xi-2] == "A" && ss[Xi-1] == "B" && ss[Xi] == "C"

		if p1Exists || p2Exists || p3Exists {
			count++
		}

		fmt.Fprintln(w, count)
	}
}

// どこにABC文字列があるかをマップに保持しておく解法
func alt() {
	defer w.Flush()

	N, Q := read2Ints(r) // len(S), len(Queries)
	S := readStr(r)      // 文字列

	ss := strings.Split(S, "")

	type abcIndex [3]int // a, b, cのインデックスを順に保持
	abcIndexes := make(map[abcIndex]struct{}, len(ss)/3)

	for i := 0; i < N-2; i++ {
		if ss[i] == "A" && ss[i+1] == "B" && ss[i+2] == "C" {
			abcIndexes[abcIndex{i, i + 1, i + 2}] = struct{}{}
		}
	}

	for i := 0; i < Q; i++ {
		sarr := readStrArr(r)
		XS := sarr[0]
		X, _ := strconv.Atoi(XS)
		C := sarr[1]

		xi := X - 1
		ss[xi] = C

		abcIndex1 := abcIndex{xi, xi + 1, xi + 2}
		abcIndex2 := abcIndex{xi - 1, xi, xi + 1}
		abcIndex3 := abcIndex{xi - 2, xi - 1, xi}

		if _, ok := abcIndexes[abcIndex1]; ok {
			if C == "A" {
				fmt.Fprintln(w, len(abcIndexes))
				continue
			} else {
				delete(abcIndexes, abcIndex1)
			}
		}

		if _, ok := abcIndexes[abcIndex2]; ok {
			if C == "B" {
				fmt.Fprintln(w, len(abcIndexes))
				continue
			} else {
				delete(abcIndexes, abcIndex2)
			}
		}

		if _, ok := abcIndexes[abcIndex3]; ok {
			if C == "C" {
				fmt.Fprintln(w, len(abcIndexes))
				continue
			} else {
				delete(abcIndexes, abcIndex3)
			}
		}

		ss[xi] = C
		if xi+2 <= len(ss)-1 && ss[xi] == "A" && ss[xi+1] == "B" && ss[xi+2] == "C" {
			abcIndexes[abcIndex1] = struct{}{}
		} else if xi+1 <= len(ss)-1 && xi-1 >= 0 && ss[xi-1] == "A" && ss[xi] == "B" && ss[xi+1] == "C" {
			abcIndexes[abcIndex2] = struct{}{}
		} else if xi-2 >= 0 && ss[xi-2] == "A" && ss[xi-1] == "B" && ss[xi] == "C" {
			abcIndexes[abcIndex3] = struct{}{}
		}

		fmt.Fprintln(w, len(abcIndexes))
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
