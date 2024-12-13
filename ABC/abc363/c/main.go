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

	_, K := read2Ints(r)

	S := readStr(r)
	ss := strings.Split(S, "")
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	permitations := Permute([]string{}, ss)

	// fmt.Printf("permitations: %v\n", permitations)

	// done := make(map[string]bool, len(permitations))
	ans := 0

	// fmt.Printf("permitations: %v\n", permitations)

Outer:
	for _, p := range permitations {
		// if done[strings.Join(p, "")] {
		// 	continue
		// }
		// done[strings.Join(p, "")] = true

	Middle:
		for i := 0; i <= len(p)-K; i++ { // K文字ずつチェック、インデックスを一個ずつずらす
			toCheck := p[i : i+K]

			for j := 0; j < K/2; j++ {
				if toCheck[j] != toCheck[len(toCheck)-1-j] {
					continue Middle
				}
			}

			continue Outer
		}

		ans++
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
// optionsには全ての要素が異なるものを渡すこと
func Permute[T comparable](current []T, options []T) [][]T {
	var results [][]T

	cc := append([]T{}, current...)
	co := append([]T{}, options...)

	if len(co) == 0 {
		return [][]T{cc}
	}

	var lastO T
	// usedMap := make(map[T]bool, len(co))
	for i, o := range options {
		// if usedMap[o] {
		// 	continue
		// }
		// usedMap[o] = true

		if o == lastO {
			continue
		}
		lastO = o

		newcc := append([]T{}, cc...)
		newcc = append(newcc, o)
		newco := append([]T{}, co[:i]...)
		newco = append(newco, co[i+1:]...)

		subResults := Permute(newcc, newco)
		results = append(results, subResults...)
	}

	return results
}

// func RevSl[S ~[]E, E any](s S) S {
// 	lenS := len(s)
// 	revS := make(S, lenS)
// 	for i := 0; i < lenS; i++ {
// 		revS[i] = s[lenS-1-i]
// 	}

// 	return revS
// }

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
