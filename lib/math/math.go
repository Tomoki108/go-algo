package math

import "strconv"

// O(n) n: numの桁数
// numの桁数を返す
func GetDigists(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	return digits
}

// O(n) n: numの桁数
// numの各桁の数字を返す
func ToDigits(n int) []int {
	s := strconv.FormatInt(int64(n), 10)
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i] - '0') // （'x'に対応する数字 - '0'に対応する数字）のruneの数字 = x as int
	}
	return digits
}

// O(log(min(a,b)))
// 拡張ユークリッドの互除法で、最大公約数(Greatest Common Divisor)を求める
// （ax + by = gcd(a, b) となるx, yも返す）
func GCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := GCD(b, a%b)
	x2 := y1
	y2 := x1 - (a/b)*y1
	return gcd, x2, y2
}

// O(log(min(a,b)))
// 最小公倍数（Least Common Multipler）を求める
func LCM(a, b int) int {
	gcd, _, _ := GCD(a, b)
	return a * b / gcd
}
