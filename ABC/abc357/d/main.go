package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
const intMax = 1 << 62

//lint:ignore U1000 unused
const intMin = -1 << 62

const M = 998244353

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	// K: Nの桁数
	// R: 10^K
	// V_N = N*R^0 + N*R^1 + N*R^2 + ... + N*R^N
	// V_N = N(R^N-1)/(R-1)
	// N(R^N-1)/(R-1) ≡ x (mod M)　// このxを求めたい
	// N(R^N-1) * (R-1)^-1 ≡ x (mod M)  // (R-1)^-1 は(R-1)の逆元（1/(R-1)ではない！）
	// GCD((R-1), M) = 1 // Mは素数なので。よってフェルマーの小定理が使える
	// (R-1)^(M-1) ≡ 1 (mod M)
	// (R-1)*(R-1)^(M-2) ≡ 1 (mod M)
	// よって、(R-1)の逆元は(R-1)^(M-2)となる
	// N(R^N-1) * (R-1)^{M-2} ≡ x (Mod M)
	// x = (N*(R^N -1) * (R-1)^{M-2}) % M

	N := readInt(r)
	K := len(strconv.Itoa(N))
	R := ModExponentiation(10, K, M)

	mN := NewModInt(N, M)
	mR := NewModInt(R, M)

	// 指数は元々の数字をそのまま使うことに注意
	x := mN.Mul(mR.Pow(N).SubI(1)).Mul(mR.SubI(1).Pow(M - 2))

	fmt.Fprintln(w, x.Val())
}

// ////////////
// Libs    //
// ///////////
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
	inverseElm := InverseElm(a.val, mi.modulo)
	return NewModInt(mi.val*inverseElm, mi.modulo)
}

// mod mi.moduloでのaの逆元を求め、mi.valに掛ける。逆元が存在するaを渡すこと
func (mi ModInt) DivI(a int) ModInt {
	return mi.Div(NewModInt(a, mi.modulo))
}

func (mi ModInt) Pow(exp int) ModInt {
	return NewModInt(ModExponentiation(mi.val, exp, mi.modulo), mi.modulo)
}

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

// O(log(m))
// mが素数かつaがmの倍数でない前提で、aのmod mにおける逆元を計算する
//
// フェルマーの小定理より以下が成り立つ。
// a^(m-1) ≡ 1 (mod m)
// a * a^(m-2) ≡ 1 (mod m)
// よってa^(m-2)がaのmod mにおける逆元となる
func InverseElm(a, m int) int {
	return ModExponentiation(a, m-2, m)
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
//
//lint:ignore U1000 unused
func readStr(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
//
//lint:ignore U1000 unused
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
//
//lint:ignore U1000 unused
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
//
//lint:ignore U1000 unused
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
//
//lint:ignore U1000 unused
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
//
//lint:ignore U1000 unused
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
//
//lint:ignore U1000 unused
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身をスペース区切りなしで出力する
//
//lint:ignore U1000 unused
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
//
//lint:ignore U1000 unused
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

//lint:ignore U1000 unused
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//lint:ignore U1000 unused
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
