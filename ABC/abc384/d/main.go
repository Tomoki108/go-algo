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

// 円環の性質を利用した解法（A=>B: x, then B=>A: 全周-x）
func alt() {
	defer w.Flush()

	N, S := read2Ints(r)
	As := readIntArr(r)

	cSum := make([]int, 0, N+1)
	cSum = append(cSum, 0)
	totalSum := 0
	for i := 0; i < N; i++ {
		totalSum += As[i]
		cSum = append(cSum, totalSum)
	}

	remainder := S % totalSum

	// cSum[i] - cSum[j] == remainder or totalSum - remainder
	cSumMap := make(map[int]struct{}, N+1)
	for i := 0; i < N+1; i++ {
		cSum_i := cSum[i]

		toFind1 := cSum_i - remainder
		toFind2 := cSum_i - (totalSum - remainder)

		_, exisits1 := cSumMap[toFind1]
		if exisits1 {
			fmt.Fprintln(w, "Yes")
			return
		}
		_, exisits2 := cSumMap[toFind2]
		if exisits2 {
			fmt.Fprintln(w, "Yes")
			return
		}

		cSumMap[cSum[i]] = struct{}{}
	}

	fmt.Fprintln(w, "No")
}

// 配列を２倍にし、そこを尺取する方法での解法
func main() {
	defer w.Flush()

	N, S := read2Ints(r)
	As := readIntArr(r)

	sum := 0
	minA := intMax
	for i := 0; i < N; i++ {
		sum += As[i]
		minA = min(minA, As[i])
	}

	if minA > S {
		fmt.Fprintln(w, "No")
		return
	}

	toFind := S % sum
	if toFind == 0 {
		fmt.Fprintln(w, "Yes")
		return
	}

	wAs := append(As, As...)

	// (left, right]の区間で尺取
	left, right := 0, 0
	currentSum := 0
	for right < len(wAs) {
		currentSum += wAs[right]
		right++

		if currentSum < toFind {
			continue
		}

		if currentSum == toFind {
			fmt.Fprintln(w, "Yes")
			return
		}

		for currentSum > toFind {
			currentSum -= wAs[left]
			left++

			if left == right {
				currentSum = 0
				break
			}

			if currentSum == toFind {
				fmt.Fprintln(w, "Yes")
				return
			}
		}

	}

	fmt.Fprintln(w, "No")
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
