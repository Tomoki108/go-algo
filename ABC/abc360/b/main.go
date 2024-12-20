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

// 純粋な、wとcの組み合わせの全探索
func main() {
	defer w.Flush()

	sarr := readStrArr(r)
	S := sarr[0]
	T := sarr[1]
	Ss := strings.Split(S, "")
	Ts := strings.Split(T, "")

	for W := 1; W < len(Ss); W++ {
		numOfChunks := len(Ss) / W
		remainder := len(Ss) % W
		if remainder != 0 {
			numOfChunks++
		}

		if numOfChunks < len(Ts) {
			break
		}

		for c := 1; c <= W; c++ {
			chars := make([]string, 0, len(Ts))
			for i := 0; i < numOfChunks; i++ {
				idx := (c - 1) + W*i

				if idx > len(Ss)-1 {
					break
				}

				chars = append(chars, Ss[idx])
			}

			if strings.Join(chars, "") == T {
				fmt.Fprintln(w, "Yes")
				return
			}
		}
	}

	fmt.Fprintln(w, "No")
}

// wを全探索し、チャンクリストを実際に作る。
// そのチャンクリストに対して可能な縦読みを全て試す。
func alt() {
	defer w.Flush()

	sarr := readStrArr(r)
	S := sarr[0]
	T := sarr[1]
	Ss := strings.Split(S, "")
	Ts := strings.Split(T, "")

	for W := 1; W < len(Ss)-1; W++ {
		chunks := SplitByChunks(Ss, W)
		if len(chunks) < len(Ts) {
			break
		}

		for c := 1; c <= W; c++ {
			vReading := make([]string, 0, len(chunks))
			for _, chunk := range chunks {
				if c <= len(chunk) {
					vReading = append(vReading, chunk[c-1])
				}
			}

			if strings.Join(vReading, "") == T {
				fmt.Fprintln(w, "Yes")
				return
			}
		}
	}

	fmt.Fprintln(w, "No")
}

// ////////////
// Libs    //
// ///////////

// O(n/size)
func SplitByChunks[T any](sl []T, chunkSize int) [][]T {
	if len(sl) == 0 {
		return [][]T{}
	}

	chunks := make([][]T, 0, (len(sl)+chunkSize-1)/chunkSize) // 余りを考慮したlengthの計算
	for chunkSize < len(sl) {
		chunks = append(chunks, sl[0:chunkSize])
		sl = sl[chunkSize:]
	}

	return append(chunks, sl)
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
