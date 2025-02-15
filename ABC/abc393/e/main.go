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

// 1000000000000000000, 19 digits, 10^18
const INF = int(1e18)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, K := read2Ints(r)
	As := readIntArr(r)

	factorCount := make(map[int64]int, N)
	fss := make([]map[int64]int, 0, N)
	for _, a := range As {
		fs := factorize(int64(a))
		fss = append(fss, fs)
		for k := range fs {
			factorCount[k] += 1
		}
	}

	for i := 0; i < N; i++ {
		fs := fss[i]

		ans := 1
		for k := range fs {
			if factorCount[k] >= K {
				ans = max(ans, int(k))
			}
		}

		fmt.Fprintln(w, ans)
	}

}

//////////////
// Libs    //
/////////////

func gcd(a, b int64) int64 {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func modPow(base, exponent, mod int64) int64 {
	result := int64(1)
	base %= mod
	for exponent > 0 {
		if exponent&1 == 1 {
			result = (result * base) % mod
		}
		exponent >>= 1
		base = (base * base) % mod
	}
	return result
}

func isPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n == 1 || n%2 == 0 {
		return false
	}

	m := n - 1
	// n - 1 = d * 2^s となるように分解
	s := 0
	d := m
	for d%2 == 0 {
		d /= 2
		s++
	}

	testNumbers := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, a := range testNumbers {
		if a == n {
			continue
		}
		x := modPow(a, d, n)
		if x == 1 {
			continue
		}
		r := 0
		for x != m {
			x = modPow(x, 2, n)
			r++
			if x == 1 || r == s {
				return false
			}
		}
	}
	return true
}

func findPrimeFactor(n int64) int64 {
	if n%2 == 0 {
		return 2
	}

	// m = int(n**0.125) + 1
	m := int64(math.Pow(float64(n), 0.125)) + 1

	// c を 1 から n-1 まで試す
	for c := int64(1); c < n; c++ {
		// f(a) = (a^2 mod n + c) mod n
		f := func(a int64) int64 {
			return (modPow(a, 2, n) + c) % n
		}

		y := int64(0)
		g, q, r, k := int64(1), int64(1), int64(1), int64(0)
		var x, ys int64

		for g == 1 {
			x = y
			// 内側のループ： k < 3*r/4 の間 f を適用
			for k < (3*r)/4 {
				y = f(y)
				k++
			}
			// k < r かつ g == 1 の間、q に累積して積をとる
			for k < r && g == 1 {
				ys = y
				steps := m
				if r-k < m {
					steps = r - k
				}
				for i := int64(0); i < steps; i++ {
					y = f(y)
					diff := x - y
					if diff < 0 {
						diff = -diff
					}
					q = (q * diff) % n
				}
				g = gcd(q, n)
				k += m
			}
			k = r
			r *= 2
		}

		if g == n {
			g = 1
			y = ys
			for g == 1 {
				y = f(y)
				diff := x - y
				if diff < 0 {
					diff = -diff
				}
				g = gcd(diff, n)
			}
		}

		if g == n {
			continue
		}
		if isPrime(g) {
			return g
		} else if isPrime(n / g) {
			return n / g
		} else {
			return findPrimeFactor(g)
		}
	}
	return n
}

func factorize(n int64) map[int64]int {
	res := make(map[int64]int)
	for !isPrime(n) && n > 1 {
		p := findPrimeFactor(n)
		s := 0
		for n%p == 0 {
			n /= p
			s++
		}
		res[p] = s
	}
	if n > 1 {
		res[n] = 1
	}
	return res
}

//////////////
// Helpers //
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

// height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArr(r)
	}

	return grid
}

// height行、width列のT型グリッドを作成
func createGrid[T any](height, width int, val T) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
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
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	if len(sl) == 0 {
		fmt.Fprintln(w)
		return
	}

	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sort2Ints(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func sort2IntsDesc(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func mapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
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

func updateToMin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func updateToMax(a *int, b int) {
	if *a < b {
		*a = b
	}
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
	if exp == 0 {
		return 1
	}

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
