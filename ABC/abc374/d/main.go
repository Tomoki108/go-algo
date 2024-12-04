package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type edge [4]int // fromX, fromY,toX, toY

var permutaions = make([][]edge, 0)

// (O(N! * 2^N))
// max: O(6! * 2^6)= O(720 * 64) = O(46080)
func main() {
	defer w.Flush()

	intarr := readIntArr(r)
	N := intarr[0] // num of edges
	S := intarr[1] // move speed
	T := intarr[2] // print speed

	edges := make([]edge, 0, N)
	for i := 0; i < N; i++ {
		arr := readIntArr(r)
		edges = append(edges, edge{arr[0], arr[1], arr[2], arr[3]})
	}

	options := make([]edge, 0, N)
	for i := 0; i < N; i++ {
		options = append(options, edges[i])
	}
	permutaions := Permute([]edge{}, options)

	var minCost float64
	minCost = 1 << 62
	for _, p := range permutaions {
		for i := uint64(0); i < 1<<N; i++ {
			var cost float64
			lastX := 0
			lastY := 0

			for j := 1; j <= N; j++ {
				toX := p[j-1][0]
				toY := p[j-1][1]
				fromX := p[j-1][2]
				fromY := p[j-1][3]

				flipped := IsBitPop(i, j)

				if !flipped {
					cost += CalcDistance(lastX, lastY, fromX, fromY) / float64(S)
					cost += CalcDistance(fromX, fromY, toX, toY) / float64(T)
					lastX = toX
					lastY = toY
				} else {
					cost += CalcDistance(lastX, lastY, toX, toY) / float64(S)
					cost += CalcDistance(toX, toY, fromX, fromY) / float64(T)
					lastX = fromX
					lastY = fromY
				}
			}
			minCost = math.Min(minCost, cost)
		}
	}

	fmt.Fprintln(w, minCost)
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

// k桁目のビットが1かどうかを判定（一番右を１桁目とする）
func IsBitPop(num uint64, k int) bool {
	// 1 << (k - 1)はビットマスク。1をk - 1桁左にシフトすることで、k桁目のみが1で他の桁が0の二進数を作る。
	// numとビットマスクの論理積（各桁について、numとビットマスクが両方trueならtrue）を作り、その結果が0でないかどうかで判定できる
	return (num & (1 << (k - 1))) != 0
}

func CalcDistance(fromX, fromY, toX, toY int) float64 {
	return math.Sqrt(float64((toX-fromX)*(toX-fromX) + (toY-fromY)*(toY-fromY)))
}

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readString(r *bufio.Reader) string {
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

// nCrの計算 O(r)
// (n * (n-1) ... * (n-r+1)) / r!
func combination(n, r int) int {
	if r > n {
		return 0
	}
	if r > n/2 {
		r = n - r // Use smaller r for efficiency
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// slices.Reverce() （Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func slReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
