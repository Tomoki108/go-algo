package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, Q := read2Ints(r)
	graph := make([]map[int]struct{}, N)
	for i := 0; i < N; i++ {
		graph[i] = make(map[int]struct{})
	}

	ans := N
	for i := 0; i < Q; i++ {
		iarr := readIntArr(r)
		q := iarr[0]

		if q == 1 {
			u, v := iarr[1]-1, iarr[2]-1

			if len(graph[u]) == 0 {
				ans--
			}
			if len(graph[v]) == 0 {
				ans--
			}
			graph[u][v] = struct{}{}
			graph[v][u] = struct{}{}
			fmt.Fprintln(w, ans)
		} else {
			u := iarr[1] - 1
			if len(graph[u]) != 0 {
				ans++
			}
			for v := range graph[u] {
				if len(graph[v]) == 1 {
					ans++
				}
				delete(graph[v], u)
			}
			graph[u] = make(map[int]struct{})
			fmt.Fprintln(w, ans)
		}
	}
}

//////////////
// Libs    //
/////////////

type UnionFind struct {
	parent []int // len(parent)分のノードを考え、各ノードの親を記録している
	size   []int // そのノードを頂点とする部分木の頂点数
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	s := make([]int, size)
	for i := range parent {
		parent[i] = i
		s[i] = 1
	}
	return &UnionFind{parent, s}
}

// O(α(N))　※定数時間。α(N)はアッカーマン関数の逆関数
// xの親を見つける
func (uf *UnionFind) Find(xIdx int) int {
	if uf.parent[xIdx] != xIdx {
		uf.parent[xIdx] = uf.Find(uf.parent[xIdx]) // 経路圧縮
	}
	return uf.parent[xIdx]
}

// O(α(N))
// xとyを同じグループに統合する（サイズが大きい方に統合）
func (uf *UnionFind) Union(xIdx, yIdx int) {
	rootX := uf.Find(xIdx)
	rootY := uf.Find(yIdx)

	if rootX != rootY {
		if uf.size[rootX] < uf.size[rootY] {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		} else if uf.size[rootX] > uf.size[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		} else {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		}
	}
}

// O(1)
func (uf *UnionFind) IsRoot(xIdx int) bool {
	return uf.parent[xIdx] == xIdx
}

// O(α(N))
func (uf *UnionFind) IsSameRoot(xIdx, yIdx int) bool {
	return uf.Find(xIdx) == uf.Find(yIdx)
}

// O(N * α(N))
func (uf *UnionFind) CountRoots() int {
	count := 0
	for i := range uf.parent {
		if uf.Find(i) == i {
			count++
		}
	}
	return count
}

// O(N * α(N))
func (uf *UnionFind) Roots() []int {
	roots := make([]int, 0)
	for i := range uf.parent {
		if uf.Find(i) == i {
			roots = append(roots, i)
		}
	}
	return roots
}

// O(α(N))
func (uf *UnionFind) GroupSize(xIdx int) int {
	return uf.size[uf.Find(xIdx)]
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
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArr(r)
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
