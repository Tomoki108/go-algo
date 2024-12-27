package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
const intMax = 1 << 62

//lint:ignore U1000 unused
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	N, M, K := iarr[0], iarr[1], iarr[2]

	keys := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		keys = append(keys, i)
	}
	patterns := PickN([]int{}, keys, K)

	caseOpened := make([]map[int]struct{}, 0, M)
	caseNotOpened := make([]map[int]struct{}, 0, M)
	for i := 0; i < M; i++ {
		sarr := readStrArr(r)

		CS := sarr[0]
		C, _ := strconv.Atoi(CS)

		ASs := sarr[1 : C+1]
		As := make(map[int]struct{}, C)
		for _, AS := range ASs {
			A, _ := strconv.Atoi(AS)
			As[A] = struct{}{}
		}

		R := sarr[C+1]

		if R == "o" {
			caseOpened = append(caseOpened, As)
		} else {
			caseNotOpened = append(caseNotOpened, As)
		}
	}

	ans := len(patterns)

Outer:
	for _, p := range patterns {
		m := make(map[int]struct{}, len(p))
		for _, key := range p {
			m[key] = struct{}{}
		}

		for _, co := range caseOpened {
			for key := range m {
				_, exist := co[key]
				if !exist {
					ans--
					continue Outer
				}
			}
		}

	Middle:
		for _, cno := range caseNotOpened {
			for key := range m {
				_, exist := cno[key]
				if !exist {
					continue Middle
				}
			}

			ans--
			continue Outer
		}
	}

	fmt.Fprintln(w, ans)
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
		newCurrent := make([]T, len(current), len(current)+1)
		copy(newCurrent, current)
		newCurrent = append(newCurrent, o)

		newOptions := make([]T, len(options[i+1:]))
		copy(newOptions, options[i+1:])

		results = append(results, PickN(newCurrent, newOptions, n-1)...)
	}

	return results
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
