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

// 9223372036854775808, 19 digits, 2^63
const INT_MAX = math.MaxInt

// -9223372036854775808, 19 digits, -1 * 2^63
const INT_MIN = math.MinInt

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	grid1 := readGrid(r, 4)
	grid2 := readGrid(r, 4)
	grid3 := readGrid(r, 4)

	getParts := func(grid [][]string) [][2]int {
		var pSlice [][2]int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if grid[i][j] == "#" {
					pSlice = append(pSlice, [2]int{i, j})
				}
			}
		}
		return pSlice
	}

	p1 := getParts(grid1)
	p2 := getParts(grid2)
	p3 := getParts(grid3)
	if len(p1)+len(p2)+len(p3) != 16 {
		fmt.Fprintln(w, "No")
		return
	}

	partsSl := [3][][2]int{p1, p2, p3}

	var partsMap [3][4][][2]int // partsNo -> rotateNo -> parts slice (sorted by most left up)
	for i := 0; i < 3; i++ {    // partsNo
		partsMap[i] = [4][][2]int{}
		parts := partsSl[i]

		for j := 0; j < 4; j++ { // rotateNo
			partsMap[i][j] = make([][2]int, len(parts))

			for k := 0; k < len(parts); k++ {
				ni, nj := RotateSquareGridCell(4, parts[k][0], parts[k][1], j)
				partsMap[i][j][k] = [2]int{ni, nj}
			}

			sort.Slice(partsMap[i][j], func(a, b int) bool {
				if partsMap[i][j][a][0] == partsMap[i][j][b][0] {
					return partsMap[i][j][a][1] < partsMap[i][j][b][1]
				}
				return partsMap[i][j][a][0] < partsMap[i][j][b][0]
			})
		}
	}

	var dfs func(mostLeftUp [2]int, partsNoPerm []int, permIdx int, grid [][]string) bool
	dfs = func(mostLeftUp [2]int, partsNoPerm []int, permIdx int, grid [][]string) bool {
		if permIdx >= 3 {
			panic("can't reach here")
		}

		dump("mostLeftUp: %v, partsNoPerm: %v, permIdx: %d\n", mostLeftUp, partsNoPerm, permIdx)

	Outer1:
		for i := 0; i <= 3; i++ {
			cgrid := CopyGrid(grid)

			parts := partsMap[partsNoPerm[permIdx]][i]
			delta := [2]int{mostLeftUp[0] - parts[0][0], mostLeftUp[1] - parts[0][1]}

			dump("parts: %v\n", parts)
			dump("delta: %v\n", delta)

			for _, part := range parts {
				nh, nw := part[0]+delta[0], part[1]+delta[1]
				c := Coordinate{nh, nw}
				if !c.IsValid(4, 4) || cgrid[c.h][c.w] == "#" {
					continue Outer1
				}
				cgrid[c.h][c.w] = "#"
			}

			// for h := 0; h < 4; h++ {
			// 	fmt.Println(cgrid[h])
			// }

			var newMostLeftUp *[2]int
		Outer2:
			for h := 0; h < 4; h++ {
				for w := 0; w < 4; w++ {
					if cgrid[h][w] == "." {
						newMostLeftUp = &[2]int{h, w}
						break Outer2
					}
				}
			}
			if newMostLeftUp == nil {
				return true
			}
			dump("newMostLeftUp: %v\n\n", *newMostLeftUp)

			newPartsIdx := permIdx + 1

			return dfs(*newMostLeftUp, partsNoPerm, newPartsIdx, cgrid)
		}

		return false
	}

	partsNoPerm := []int{0, 1, 2}
	next := true
	for next {
		if dfs([2]int{0, 0}, partsNoPerm, 0, createGrid(4, 4, ".")) {
			fmt.Fprintln(w, "Yes")
			return
		}
		next = NextPermutation(partsNoPerm)
	}

	fmt.Fprintln(w, "No")
}

//////////////
// Libs    //
/////////////

// O(H*W)
// T型グリッドのコピーを作成する
func CopyGrid[T any](grid [][]T) [][]T {
	H := len(grid)
	W := len(grid[0])
	res := make([][]T, H)
	for i := 0; i < H; i++ {
		res[i] = make([]T, W)
		copy(res[i], grid[i])
	}
	return res
}

// 一辺がnの正方形グリッドのマス目(hight, width)を、時計回りにtime回回転させたときの座標を返す
func RotateSquareGridCell(n, height, width, time int) (h, w int) {
	time = time % 4
	switch time {
	case 0:
		return height, width
	case 1:
		return width, n - height + 1
	case 2:
		return n - height + 1, n - width + 1
	case 3:
		return n - width + 1, height
	}

	panic("can't reach here")
}

type Coordinate struct {
	h, w int // 0-indexed
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
}

// NOTE:
// next := true; for next { some(sl); next = NextPermutation(sl); } で使う
//
// O(len(sl)*len(sl)!)
// sl の要素を並び替えて、次の辞書順の順列にする
func NextPermutation[T ~int | ~string](sl []T) bool {
	n := len(sl)
	i := n - 2

	// Step1: 右から左に探索して、「スイッチポイント」を見つける:
	// 　「スイッチポイント」とは、右から見て初めて「リストの値が減少する場所」です。
	// 　例: [1, 2, 3, 6, 5, 4] の場合、3 がスイッチポイント。
	for i >= 0 && sl[i] >= sl[i+1] {
		i--
	}

	//　スイッチポイントが見つからない場合、最後の順列に到達しています。
	if i < 0 {
		return false
	}

	// Step2: スイッチポイントの右側の要素から、スイッチポイントより少しだけ大きい値を見つけ、交換します。
	// 　例: 3 を右側で最小の大きい値 4 と交換。
	j := n - 1
	for sl[j] <= sl[i] {
		j--
	}
	sl[i], sl[j] = sl[j], sl[i]

	// Step3: スイッチポイントの右側を反転して、辞書順に次の順列を作ります。
	// 　例: [1, 2, 4, 6, 5, 3] → [1, 2, 4, 3, 5, 6]。
	reverse(sl[i+1:])
	return true
}

func reverse[T ~int | ~string](sl []T) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
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

func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
