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

func main() {
	defer w.Flush()

	N := readInt(r)
	S := readStr(r)
	ss := strings.Split(S, "")

	// 横：何回目まで終わったか
	// 縦：この回に [0]勝つ || [1]引き分ける
	// セル：通算勝利数
	var table [2][]int

	table[0] = make([]int, N)
	table[1] = make([]int, N)

	table[0][0] = 1
	table[1][0] = 0
	winLastHand := winHand(ss[0])
	tieLastHand := ss[0]

	for i := 1; i < N; i++ {
		oponentH := ss[i]
		winH := winHand(oponentH)
		tieH := oponentH

		// この回に勝利する場合の最大勝利数
		if winH != winLastHand && winH != tieLastHand {
			table[0][i] = max(table[0][i-1], table[1][i-1]) + 1
		} else if winH == winLastHand {
			table[0][i] = table[1][i-1] + 1
		} else {
			table[0][i] = table[0][i-1] + 1
		}

		// この回に引き分ける場合の最大勝利数
		if tieH != winLastHand && tieH != tieLastHand {
			table[1][i] = max(table[0][i-1], table[1][i-1])
		} else if tieH == winLastHand {
			table[1][i] = table[1][i-1]
		} else {
			table[1][i] = table[0][i-1]
		}

		winLastHand = winH
		tieLastHand = tieH
	}

	ans := max(table[0][N-1], table[1][N-1])

	fmt.Fprintln(w, ans)
}

func winHand(oppenentHand string) string {
	switch oppenentHand {
	case "R":
		return "P"
	case "P":
		return "S"
	case "S":
		return "R"
	}

	panic("invalid hand")
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
