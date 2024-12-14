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

	// fmt.Println("hey")
	// fmt.Println(m)

	comb := map[int][][]string{
		5: {[]string{"A", "B", "C", "D", "E"}},
		4: {[]string{"A", "B", "C", "D"}, []string{"A", "B", "C", "E"}, []string{"A", "B", "D", "E"}, []string{"A", "C", "D", "E"}, []string{"B", "C", "D", "E"}},
		3: {[]string{"A", "B", "C"}, []string{"A", "B", "D"}, []string{"A", "B", "E"}, []string{"A", "C", "D"}, []string{"A", "C", "E"}, []string{"A", "D", "E"}, []string{"B", "C", "D"}, []string{"B", "C", "E"}, []string{"B", "D", "E"}, []string{"C", "D", "E"}},
		2: {[]string{"A", "B"}, []string{"A", "C"}, []string{"A", "D"}, []string{"A", "E"}, []string{"B", "C"}, []string{"B", "D"}, []string{"B", "E"}, []string{"C", "D"}, []string{"C", "E"}, []string{"D", "E"}},
		1: {[]string{"A"}, []string{"B"}, []string{"C"}, []string{"D"}, []string{"E"}},
	}

	var ps [][]string
	for i := 1; i <= 5; i++ {
		ps = append(ps, comb[i]...)
	}

	// var ps [][]string
	// for i := 5; i >= 1; i-- {
	// 	optionsSl := comb[i]
	// 	for _, options := range optionsSl {

	// 		// fmt.Println("hey")

	// 		result := Permute([]string{}, options)

	// 		// fmt.Println(result)

	// 		ps = append(ps, result...)
	// 	}

	// }

	sort.Slice(ps, func(i, j int) bool {
		p1 := ps[i]
		p2 := ps[j]

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

	for _, p := range ps {
		fmt.Fprintln(w, strings.Join(p, ""))
	}

	// // fmt.Println(pss[0])
	// // return

	// for _, ps := range pss {
	// 	sort.Slice(ps, func(i, j int) bool {
	// 		p1 := ps[i]
	// 		p2 := ps[j]

	// 		score1 := 0
	// 		score2 := 0

	// 		length := len(p1)

	// 		for i := 0; i < length; i++ {
	// 			char := p1[i]
	// 			score1 += m[char]

	// 			char2 := p2[i]
	// 			score2 += m[char2]
	// 		}

	// 		return score1 > score2
	// 	})

	// 	for _, p := range ps {
	// 		fmt.Println(w, strings.Join(p, ""), "\n")
	// 	}
	// }

}

//////////////
// Libs    //
/////////////

func Permute[T comparable](current []T, options []T) [][]T {
	var results [][]T

	cc := append([]T{}, current...)
	co := append([]T{}, options...)

	if len(co) == 0 {
		return [][]T{cc}
	}

	var lastO T
	for i, o := range options {
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

// 要素数 len(options) で、i番目の要素が options[i] であるような順列のパターンを全列挙する
// options[i]に重複した要素が含まれていても、あらかじめソートしておけば重複パターンは除かれる
func Permute2[T comparable](current []T, options [][]T) [][]T {
	var results [][]T

	if len(current) == len(options) {
		results = append(results, current)
		return results
	}

	nextVals := options[len(current)]
	var lastV T
	for _, v := range nextVals {
		if v == lastV {
			continue
		}
		lastV = v

		copyCurrent := append([]T{}, current...)
		copyCurrent = append(copyCurrent, v)
		subResults := Permute2(copyCurrent, options)
		results = append(results, subResults...)
	}

	return results
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
