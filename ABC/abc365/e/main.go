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
	As := readIntArr(r)

	pXor := PrefixXOR(As)

	N++
	ans := 0
	// 10^8 = 101111101011110000100000000 = 27
	for digit := 0; digit < 27; digit++ {
		popped := 0
		for i := 0; i < N; i++ {
			if IsBitPop(uint64(pXor[i]), digit) {
				popped++
			}
		}
		ans += (1 << digit) * (N - popped) * popped
	}
	for i := 0; i < N-1; i++ {
		ans -= As[i]
	}

	fmt.Fprintln(w, ans)
}

//////////////
// Libs    //
/////////////

// O(n)
// 一次元累積XORを返す（index0には0を入れる。）
func PrefixXOR(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] ^ sl[i]
	}
	return res
}

// k桁目のビットが1かどうかを判定（一番右を0桁目とする）
func IsBitPop(num uint64, k int) bool {
	// 1 << k はビットマスク。1をk桁左にシフトすることで、k桁目のみが1で他の桁が0の二進数を作る。
	// numとビットマスクの論理積（各桁について、numとビットマスクが両方trueならtrue）を作り、その結果が0でないかどうかで判定できる
	return (num & (1 << k)) != 0
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
