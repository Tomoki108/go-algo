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

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	N, X := read2Ints(r)
	Ts := readIntArr(r)

	M := 998244353

	pickOdd := NewModInt(1, M).Div(NewModInt(N, M))

	// do[i]: 丁度時刻iで曲の再生が停止する確率
	dp := make([]ModInt, X+1)
	dp[0] = NewModInt(1, M)
	for i := 1; i <= X; i++ {
		dp[i] = NewModInt(0, M)
	}

	for i := 0; i < X; i++ {
		for j := 0; j < N; j++ {
			ni := i + Ts[j]
			if ni > X {
				continue
			}
			toAdd := dp[i].Mul(pickOdd)
			dp[ni] = dp[ni].Add(toAdd)
		}
	}

	ans := NewModInt(0, M)

	minTime := X - Ts[0] + 1
	if minTime < 0 {
		minTime = 0
	}
	for i := minTime; i <= X; i++ {
		ans = ans.Add((dp[i]).Mul(pickOdd))
	}

	fmt.Println(ans.Val())
}

//////////////
// Libs    //
/////////////

// a割るbの、数学における剰余を返す。
// a = b * Quotient + RemainderとなるRemainderを返す（Quotientは負でもよく、Remainderは常に0以上という制約がある）
// goのa%bだと、|a|割るbの剰余にaの符号をつけて返すため、負の数が含まれる場合数学上の剰余とは異なる。
func Mod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

// O(log(exp))
// Calc (base^exp) % mod efficiently
func ModExponentiation(base, exp, mod int) int {
	result := 1
	base = base % mod // 基数を mod で割った余りに変換

	for exp > 0 {
		// exp の最下位ビットが 1 なら結果に base を掛ける
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		// base を二乗し、exp を半分にする
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

type ModInt struct {
	val, modulo int
}

func NewModInt(v, m int) ModInt {
	return ModInt{val: Mod(v, m), modulo: m}
}

func (mi ModInt) Val() int {
	return mi.val
}

func (mi ModInt) Add(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val+a.val, mi.modulo)
}

func (mi ModInt) AddI(a int) ModInt {
	return mi.Add(NewModInt(a, mi.modulo))
}

func (mi ModInt) Sub(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val-a.val, mi.modulo)
}

func (mi ModInt) SubI(a int) ModInt {
	return mi.Sub(NewModInt(a, mi.modulo))
}

func (mi ModInt) Mul(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val*a.val, mi.modulo)
}

func (mi ModInt) MulI(a int) ModInt {
	return mi.Mul(NewModInt(a, mi.modulo))
}

// mod mi.moduloでのaの逆元を求め、mi.valに掛ける。逆元が存在するaを渡すこと
func (mi ModInt) Div(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	inverseElm, err := InverseElmByGCD(a.val, mi.modulo)
	if err != nil {
		panic(err)
	}

	return NewModInt(mi.val*inverseElm, mi.modulo)
}

// mod mi.moduloでのaの逆元を求め、mi.valに掛ける。逆元が存在するaを渡すこと
func (mi ModInt) DivI(a int) ModInt {
	return mi.Div(NewModInt(a, mi.modulo))
}

// 指数expはexp mod Mに置き換えられないので、int型のまま受け取る
func (mi ModInt) Pow(exp int) ModInt {
	return NewModInt(ModExponentiation(mi.val, exp, mi.modulo), mi.modulo)
}

// O(log(min(a,m)))
// 拡張ユークリッドの互除法で、aのmにおける逆元を求める（aとmが互いに素でなければエラーを返す）
// a*x + m*y = 1 となるx, yがわかる。
// a*x + m*y ≡ 1 (Mod m)
// a*x ≡ 1 (Mod m)
// よってxがaのmにおける逆元となる （ただし拡張ユークリッドの互除法で求まるxは負の数の場合もあるので、調整する）
func InverseElmByGCD(a, m int) (int, error) {
	gcd, x, _ := extendedGCD(a, m)
	if gcd != 1 {
		return 0, fmt.Errorf("逆元は存在しません (gcd(%d, %d) = %d)", a, m, gcd)
	}
	return Mod(x, m), nil
}

// O(log(min(a,b)))
// 拡張ユークリッドの互除法で、最大公約数を求める
// （ax + by = gcd(a, b) となるx, yも返す）
func extendedGCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x2 := y1
	y2 := x1 - (a/b)*y1
	return gcd, x2, y2
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

// O(log(exp))
// 繰り返し二乗法で x^y を計算する関数
func pow(base, exp int) int {
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

func dump(format string, a ...interface{}) {
	if dumpFlag {
		fmt.Printf(format, a...)
	}
}
