package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

var N int

func main() {
	defer w.Flush()

	n, K := read2Ints(r)
	N = n
	Rs := readIntArr(r)

	candidateNums := make([][]int, N)
	for i := 0; i < N; i++ {
		candidateNums[i] = make([]int, 0, N)
		for j := 1; j <= Rs[i]; j++ {
			candidateNums[i] = append(candidateNums[i], j)
		}
	}

	dfs([]int{}, candidateNums)

	sort.Slice(sequences, func(i, j int) bool {
		for k := 0; k < N; k++ {
			if sequences[i][k] == sequences[j][k] {
				continue
			}
			return sequences[i][k] < sequences[j][k]
		}

		panic("unreachable")
	})

	found := false
	for _, seq := range sequences {
		sum := 0
		for j := 0; j < N; j++ {
			sum += seq[j]
		}
		if sum%K == 0 {
			found = true
			writeSlice(w, seq)
		}
	}

	if !found {
		fmt.Fprintln(w)
	}
}

var sequences [][]int

func dfs(currentSeq []int, candidateNums [][]int) {
	if len(currentSeq) == N {
		sequences = append(sequences, currentSeq)
		return
	}

	nextValCandidates := candidateNums[len(currentSeq)]
	for _, v := range nextValCandidates {
		copySeq := make([]int, 0, len(currentSeq))
		copySeq = append(copySeq, currentSeq...)
		copySeq = append(copySeq, v)

		dfs(copySeq, candidateNums)
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
