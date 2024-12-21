package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// タイルを通る：現在中（枠線を含まない）にいるタイルから、別のタイルの中に入った時に、元いたタイルを通ったとみなす。
//
// 5.5, 0.5 =>  math.Ceil(5.5), math.Ceil(0.5)の正方形に属する =>（5, 0）
// 4.5. 0.5 => (4, 0)
// xが小さい方の正方形の、xyの和が偶数かつ、xの差が1である場合に、同じタイルに属する
func main() {
	defer w.Flush()

	Sx, Sy := read2Ints(r)
	Tx, Ty := read2Ints(r)

	if Sx == Tx && Sy == Ty {
		fmt.Fprintln(w, 0)
		return
	}

	xDelta := abs(Tx - Sx)
	yDelta := abs(Ty - Sy)

	if yDelta >= xDelta {
		fmt.Fprintln(w, yDelta)
		return
	}

	// xが同じ場合はすでにここまででリターン済み
	var goRight bool
	if Sx < Tx {
		goRight = true
	}

	if yDelta == 0 {
		cost := 0
		isAtLeft := isLeftAtTile(Sx, Sy)

		if goRight {
			if isAtLeft {
				cost = xDelta / 2
			} else {
				cost = xDelta/2 + xDelta%2
			}
		} else {
			if !isAtLeft {
				cost = xDelta / 2
			} else {
				cost = xDelta/2 + xDelta%2
			}
		}

		fmt.Fprintln(w, cost)
		return
	} else {
		cost := yDelta
		cost += (xDelta - yDelta) / 2

		fmt.Fprintln(w, cost)
		return
	}
}

func isLeftAtTile(x, y int) bool {
	if y%2 == 0 {
		return x%2 == 0
	} else {
		return x%2 == 1
	}
}

//////////////
// Libs    //
/////////////

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
