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

	N, M := read2Ints(r)

	mostRightBlackWs := make([]int, N) // 数字が大きいほど右
	mostLeftWhiteWs := make([]int, N)  // 数字が小さいほど左

	mostBottomBlackHs := make([]int, N) // 数字が大きいほど下
	mostTopWhiteHs := make([]int, N)    //  数字が小さいほど上

	for i := 0; i < M; i++ {
		sarr := readStrArr(r)
		XS, YS, C := sarr[0], sarr[1], sarr[2]

		X, _ := strconv.Atoi(XS) // hight
		Y, _ := strconv.Atoi(YS) // width

		if C == "B" {
			// 横の矛盾チェック
			mostLeftWhiteW := mostLeftWhiteWs[X-1]
			if mostLeftWhiteW != 0 && mostLeftWhiteW < Y { // 最も左にある白より右にある（数字が大きい）黒があるとだめ
				fmt.Fprintln(w, "No")
				return
			}
			mostRightBlackW := mostRightBlackWs[X-1]
			mostRightBlackWs[X-1] = max(mostRightBlackW, Y)

			// 縦の矛盾チェック
			mostTopWhiteH := mostTopWhiteHs[Y-1]
			if mostTopWhiteH != 0 && mostTopWhiteH < X { // 最も上にある白より下にある（数字が大きい）黒があるとだめ
				fmt.Fprintln(w, "No")
				return
			}
			mostBottomBlackH := mostBottomBlackHs[Y-1]
			mostBottomBlackHs[Y-1] = max(mostBottomBlackH, X)
		} else {
			// 横の矛盾チェック
			mostRightBlackW := mostRightBlackWs[X-1]
			if mostRightBlackW != 0 && mostRightBlackW > Y { // 最も右にある黒より左にある（数字が小さい）白があるとだめ
				fmt.Fprintln(w, "No")
				return
			}
			mostLeftWhiteW := mostLeftWhiteWs[X-1]
			if mostLeftWhiteW == 0 {
				mostLeftWhiteWs[X-1] = Y
			} else {
				mostLeftWhiteWs[X-1] = min(mostLeftWhiteW, Y)
			}

			// 縦の矛盾チェック
			mostBottomBlackH := mostBottomBlackHs[Y-1]
			if mostBottomBlackH != 0 && mostBottomBlackH > X { // 最も下にある黒より上にある（数字が小さい）白があるとだめ
				fmt.Fprintln(w, "No")
				return
			}
			mostTopWhiteH := mostTopWhiteHs[Y-1]
			if mostTopWhiteH == 0 {
				mostTopWhiteHs[Y-1] = X
			} else {
				mostTopWhiteHs[Y-1] = min(mostTopWhiteH, X)
			}
		}

		// fmt.Printf("mostRightBlackWs: %v\n", mostRightBlackWs)
		// fmt.Printf("mostLeftWhiteWs: %v\n", mostLeftWhiteWs)
		// fmt.Printf("mostBottomBlackHs: %v\n", mostBottomBlackHs)
		// fmt.Printf("mostTopWhiteHs: %v\n", mostTopWhiteHs)
		// fmt.Println()
	}

	fmt.Fprintln(w, "Yes")
}

//////////////
// Libs    //
/////////////

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
