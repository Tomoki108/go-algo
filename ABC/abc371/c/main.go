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

func main() {
	defer w.Flush()

	N := readInt(r)
	M_G := readInt(r)

	graphG := make(map[int][]int, N)
	for i := 0; i < M_G; i++ {
		u, v := read2Ints(r)
		graphG[u] = append(graphG[u], v)
		graphG[v] = append(graphG[v], u)
	}

	M_H := readInt(r)
	graphH := make(map[int][]int, N)
	for i := 0; i < M_H; i++ {
		a, b := read2Ints(r)
		graphH[a] = append(graphH[a], b)
		graphH[b] = append(graphH[b], a)
	}

	type edge struct {
		u, v int
	}

	As := make(map[edge]int, 0)
	for i := 1; i <= N-1; i++ {
		intarr := readIntArr(r)
		for idx, val := range intarr {
			As[edge{i, i + idx + 1}] = val
		}
	}

	options := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		options = append(options, i)
	}
	permutations := Permute([]int{}, options) // graphGの各点に対して、graphHのどの頂点を対応させるかの順列を全列挙

	minCost := 1 << 60
	for _, p := range permutations {
		HtoGconvertMap := make(map[int]int, N) // graphHの頂点 => graphGの頂点 への変換マップ
		GtoHconvertMap := make(map[int]int, N) // graphGの頂点 => graphHの頂点 への変換マップ
		for idx, val := range p {
			HtoGconvertMap[val] = idx + 1
			GtoHconvertMap[idx+1] = val
		}

		convertedGraphH := make(map[int][]int, N)
		for i, sl := range graphH {
			convertedSl := make([]int, 0, len(sl))
			for _, v := range sl {
				convertedSl = append(convertedSl, HtoGconvertMap[v])
			}
			convertedGraphH[HtoGconvertMap[i]] = convertedSl
		}

		cost := 0
		for j := 1; j <= N; j++ {
			gAdjacents := graphG[j]
			hAdjacents := convertedGraphH[j]

			diff := symmetricDifference(hAdjacents, gAdjacents)
			for _, v := range diff {

				j_h := GtoHconvertMap[j]
				v_h := GtoHconvertMap[v]

				var start, end int
				start, end = min(j_h, v_h), max(j_h, v_h)

				cost += As[edge{start, end}]
			}
		}

		minCost = min(minCost, cost)
	}

	fmt.Fprintln(w, minCost/2)
}

// 順列のパターンを全列挙する
// ex, Permute([]int{}, []int{1, 2, 3}) returns [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
func Permute[T any](current []T, options []T) [][]T {
	var results [][]T

	cc := append([]T{}, current...)
	co := append([]T{}, options...)

	if len(co) == 0 {
		return [][]T{cc}
	}

	for i, o := range options {
		newcc := append([]T{}, cc...)
		newcc = append(newcc, o)
		newco := append([]T{}, co[:i]...)
		newco = append(newco, co[i+1:]...)

		subResults := Permute(newcc, newco)
		results = append(results, subResults...)
	}

	return results
}

func symmetricDifference(slice1, slice2 []int) []int {
	// 要素の出現回数を記録するマップ
	countMap := make(map[int]int)

	// slice1 の要素をマップに記録
	for _, v := range slice1 {
		countMap[v]++
	}

	// slice2 の要素をマップに記録
	for _, v := range slice2 {
		countMap[v]++
	}

	// 片方にのみ含まれる要素を収集
	var result []int
	for k, v := range countMap {
		if v == 1 { // 1度だけ出現した要素を追加
			result = append(result, k)
		}
	}

	return result
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
