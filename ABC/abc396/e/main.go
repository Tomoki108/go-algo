package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 9223372036854775807, 19 digits, 2^63 - 1
const INT_MAX = math.MaxInt64

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt64

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)

	graph := make([][][2]int, N)
	edges := make([][3]int, 0, M)
	for i := 0; i < M; i++ {
		iarr := readIntArr(r)
		X, Y, Z := iarr[0]-1, iarr[1]-1, iarr[2]
		graph[X] = append(graph[X], [2]int{Y, Z})
		graph[Y] = append(graph[Y], [2]int{X, Z})
		edges = append(edges, [3]int{X, Y, Z})
	}

	nodeVals := make([]int, N)
	for i := 0; i < N; i++ {
		nodeVals[i] = -1
	}

	for _, e := range edges {
		X, Y, Z := e[0], e[1], e[2]

		if nodeVals[X] == -1 && nodeVals[Y] == -1 {
			nodeVals[X] = 0
			nodeVals[Y] = Z
		} else if nodeVals[Y] == -1 {
			nodeVals[Y] = nodeVals[X] ^ Z
		} else if nodeVals[X] == -1 {
			nodeVals[X] = nodeVals[Y] ^ Z
		} else {
			if nodeVals[X]^nodeVals[Y] != Z {
				fmt.Fprintln(w, -1)
				return
			}
		}
	}

	for i := 0; i < N; i++ {
		if nodeVals[i] == -1 {
			nodeVals[i] = 0
		}
	}

	visited := make([]bool, N)
	var divComponents func(node int, component []int) []int
	divComponents = func(node int, component []int) []int {
		component = append(component, node)

		for _, next := range graph[node] {
			nextNode := next[0]
			if visited[nextNode] {
				continue
			}
			visited[nextNode] = true

			component = divComponents(nextNode, component)
		}
		return component
	}

	components := make([][]int, 0)
	for i := 0; i < N; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true

		component := make([]int, 0)
		components = append(components, divComponents(i, component))
	}

	for _, component := range components {
		for i := 0; i <= 30; i++ {
			zeroCnt := 0
			oneCnt := 0

			for _, node := range component {
				if IsBitPop(uint64(nodeVals[node]), i) {
					oneCnt++
				} else {
					zeroCnt++
				}
			}

			if oneCnt > zeroCnt {
				for _, node := range component {
					nodeVals[node] = int(BitFlip(uint64(nodeVals[node]), i))
				}
			}
		}
	}

	writeSlice(w, nodeVals)
}

//////////////
// Libs    //
/////////////

// k桁目のビットが1かどうかを判定（一番右を0桁目とする）
func IsBitPop(num uint64, k int) bool {
	// 1 << k はビットマスク。1をk桁左にシフトすることで、k桁目のみが1で他の桁が0の二進数を作る。
	// numとビットマスクの論理積（各桁について、numとビットマスクが両方trueならtrue）を作り、その結果が0でないかどうかで判定できる
	return (num & (1 << k)) != 0
}

// k桁目のビットが立っていれば0に、立っていなければ1にする（一番右を0桁目とする）
func BitFlip(num uint64, k int) uint64 {
	if IsBitPop(num, k) {
		return num & ^(1 << k) // &^ はビットクリア演算子。A &^ Bは、AからBのビットが立っている桁を0にしたものを返す。
	} else {
		return num | (1 << k) // | は論理和演算子。A | Bは、少なくともどちらか一方のビットが立っている桁を1にしたものを返す。
	}
}

//////////////
// Helpers //
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

// １行の「整数 文字列」のみの入力を読み込む
func readIntStr(r *bufio.Reader) (int, string) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i, _ := strconv.Atoi(strs[0])
	return i, strs[1]
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

// height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int, withSpace bool) [][]int {
	if withSpace {
		grid := make([][]int, height)
		for i := 0; i < height; i++ {
			grid[i] = readIntArr(r)
		}
		return grid
	}

	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		strs := strings.Split(str, "")

		grid[i] = make([]int, len(strs))
		for j, s := range strs {
			grid[i][j], _ = strconv.Atoi(s)
		}
	}
	return grid
}

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int, val T) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
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

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func btoi(b string) int {
	num, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sort2Ints(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func sort2IntsDesc(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

func updateToMin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func updateToMax(a *int, b int) {
	if *a < b {
		*a = b
	}
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
	if exp == 0 {
		return 1
	}

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

//////////////
// Debug   //
/////////////

var dumpFlag bool

func init() {
	args := os.Args
	dumpFlag = len(args) > 1 && args[1] == "-dump"
}

// NOTE: ループの中で使うとわずかに遅くなることに注意
func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
