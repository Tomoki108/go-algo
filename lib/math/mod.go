package math

import "fmt"

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
