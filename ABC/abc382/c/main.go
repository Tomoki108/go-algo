package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, M := read2Ints(r)

	type gourmet struct {
		fno int // 寿司のno,
		pno int // 人のno
		g   int
	}

	gourmets := make([]gourmet, 0, N+M)

	as, _ := r.ReadString('\n')
	ass := strings.Fields(as)
	for i, s := range ass {
		a, _ := strconv.Atoi(s)

		gourmets = append(gourmets, gourmet{
			pno: i + 1,
			g:   a,
		})
	}

	bs, _ := r.ReadString('\n')
	bss := strings.Fields(bs)
	for i, s := range bss {
		b, _ := strconv.Atoi(s)

		gourmets = append(gourmets, gourmet{
			fno: i + 1,
			g:   b,
		})

	}

	sort.Slice(gourmets, func(i, j int) bool {
		if gourmets[i].g < gourmets[j].g {
			return true
		}

		if gourmets[i].g == gourmets[j].g {
			if gourmets[i].pno != 0 && gourmets[j].fno != 0 {
				return true
			}

			if gourmets[i].pno != 0 && gourmets[j].pno != 0 {
				return gourmets[i].pno > gourmets[j].pno
			}
		}

		return false
	})

	fmt.Printf("gourmets: %+v", gourmets)

	ans := make(map[int]int, M)
	for i := 0; i < N+M; i++ {
		fno := gourmets[i].fno
		if fno != 0 {
			ans[fno] = -1
			for j := i - 1; -1 < j; j-- {
				if gourmets[j].pno != 0 {
					ans[fno] = gourmets[j].pno
					break
				}
			}
		}
	}

	for i := 1; i <= M; i++ {
		fmt.Fprintln(w, ans[i])
	}
}

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readString(r *bufio.Reader) string {
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
		grid[i] = readStrArr(r)
	}

	return grid
}

// 文字列グリッドを出力する
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
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

// nCrの計算 O(r)
// (n * (n-1) ... * (n-r+1)) / r!
func combination(n, r int) int {
	if r > n {
		return 0
	}
	if r > n/2 {
		r = n - r // Use smaller r for efficiency
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// slices.Reverce() （Goのバージョンが1.21以前だと使えないため）
// 計算量: O(n)
func slReverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
