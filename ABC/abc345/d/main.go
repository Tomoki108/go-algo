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

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	N, H, W := iarr[0], iarr[1], iarr[2]

	abs := make([]AB, 0, N)
	for i := 0; i < N; i++ {
		ab := AB{}
		ab.A, ab.B = read2Ints(r)
		abs = append(abs, ab)
	}

	bits := make([]int, 0, pow(2, N))
	for bit := 0; bit <= 1<<N-1; bit++ {
		bits = append(bits, bit)
	}

	abIndexes := make([]int, 0, N)
	for i := 0; i < N; i++ {
		abIndexes = append(abIndexes, i)
	}

	for _, bit := range bits {
		usedGrid := createGrid[bool](H, W)
		minH := 0
		minW := 0
	Middle1:
		for i, abIdx := range abIndexes {
			ab := abs[abIdx]
			if IsBitPop(uint64(bit), i) {
				ab = AB{A: ab.B, B: ab.A}
			}

			for h := minH; h < minH+ab.A; h++ {
				for w := minW; w < minW+ab.B; w++ {
					c := Coordinate{h, w}
					if !c.IsValid(H, W) {
						break Middle1
					}

					if usedGrid[h][w] {
						break Middle1
					}
					usedGrid[h][w] = true
				}
			}

			for h := 0; h < H; h++ {
				for w := 0; w < W; w++ {
					if !usedGrid[h][w] {
						minH = h
						minW = w
						continue Middle1
					}
				}
			}

			fmt.Println("Yes")
			return
		}
	}

	for NextPermutation(abIndexes) {
		for _, bit := range bits {
			usedGrid := createGrid[bool](H, W)
			minH := 0
			minW := 0
		Middle2:
			for i, abIdx := range abIndexes {
				ab := abs[abIdx]
				if IsBitPop(uint64(bit), i) {
					ab = AB{A: ab.B, B: ab.A}
				}

				for h := minH; h < minH+ab.A; h++ {
					for w := minW; w < minW+ab.B; w++ {
						c := Coordinate{h, w}
						if !c.IsValid(H, W) {
							break Middle2
						}

						if usedGrid[h][w] {
							break Middle2
						}
						usedGrid[h][w] = true
					}
				}

				for h := 0; h < H; h++ {
					for w := 0; w < W; w++ {
						if !usedGrid[h][w] {
							minH = h
							minW = w
							continue Middle2
						}
					}
				}

				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}

type AB struct {
	A, B int // A: 縦, B: 横
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

// NOTE: 全パターンに何らかの処理を適用したいとき、オリジナルのslに対しては別途処理を記述する
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
	//  例: [1, 2, 4, 6, 5, 3] -> [1, 2, 4, 6, 5, 3] -> [1, 2, 4, 3, 5, 6]。
	//  Goの1.21以上ならslices(sl[i+1:])でいい。
	reverse(sl[i+1:])
	return true
}

func reverse[T ~int | ~string](sl []T) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

type Coordinate struct {
	h, w int // 0-indexed
}

func (c Coordinate) IsValid(H, W int) bool {
	return 0 <= c.h && c.h < H && 0 <= c.w && c.w < W
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

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
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
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
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
