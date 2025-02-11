package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused 9223372036854775808, 19 digits, equiv 2^63
const INT_MAX = math.MaxInt

//lint:ignore U1000 unused -9223372036854775808, 19 digits, equiv -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)

	edges := make([][3]int, 0, M) // []{cost, node_a, node_b}
	for i := 0; i < M; i++ {
		K, C := read2Ints(r)
		As := readIntArr(r)
		for i := 0; i < K-1; i++ {
			edges = append(edges, [3]int{C, As[i], As[i+1]})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	ans := 0
	uf := NewUnionFind(N)
	for _, edge := range edges {
		if !uf.IsSameRoot(edge[1]-1, edge[2]-1) {
			uf.Union(edge[1]-1, edge[2]-1)
			ans += edge[0]
		}
	}

	if uf.CountRoots() > 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
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

// O(N)
func (uf *UnionFind) CountRoots() int {
	count := 0
	for i := range uf.parent {
		if uf.parent[i] == i {
			count++
		}
	}
	return count
}

// O(N)
func (uf *UnionFind) Roots() []int {
	roots := make([]int, 0)
	for i := range uf.parent {
		if uf.parent[i] == i {
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
