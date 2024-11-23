package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// ポイント：円について、from~toが1~(1+距離)になるように回転して調節する。from:8 to:1のような状況は考えづらいため。
// from < toではない時にfromとtoを入れ替えるやり方も可。位置が入れ替わっても、移動コストや移動可能な道は変わらない。図を書くと分かる。
func main() {
	defer w.Flush()

	strs := ReadIntArr(r)
	N := strs[0]
	Q := strs[1]

	L := 1
	R := 2
	cost := 0
	for i := 0; i < Q; i++ {
		strs := ReadStrArr(r)
		H := strs[0]
		T, _ := strconv.Atoi(strs[1])

		if H == "L" {
			cost += getCost(L, T, R, N)
			L = T
		} else {
			cost += getCost(R, T, L, N)
			R = T
		}
	}

	fmt.Fprint(w, cost)
}

func getCost(from, to, ng, N int) int {
	// from = 1 になるように円を回転させる
	diff := from - 1
	from = 1
	to = to - diff
	if to < 1 {
		to = N + to
	}

	ng = ng - diff
	if ng < 1 {
		ng = N + ng
	}

	if to < ng {
		return to - from
	}

	return N - (to - from)
}

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func ReadString(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
func ReadInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
func ReadStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = s
	}

	return arr
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
func ReadIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}
