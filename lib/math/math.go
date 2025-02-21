package math

import (
	"math"
	"strconv"
)

// O(n) n: numの桁数
// numの桁数を返す
func GetDigits(num int) int {
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

// √n以上の、最も√nに近い整数を返す
func Sqrt(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n))))
}

// O(n)
// nを階乗する
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// O(√n)
// 素因数分解を行い、素因数=>指数のmapを返す（keyは昇順）
func PrimeFactorization(n int) map[int]int {
	pf := make(map[int]int)
	// 因数候補は√nまででいい。
	// √nより大きい数で割った場合、商は√nより小さい数になるため、√n以下の検証時にその割り算は済んでいる。
	for factor := 2; factor*factor <= n; factor++ {
		for n%factor == 0 {
			pf[factor]++
			n /= factor
		}
	}
	if n > 1 {
		pf[n]++
	}
	return pf
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

// O(log(n))
// log_2_nを返す
func Log(n int) int {
	ans := 1
	for {
		if n == 1 {
			break
		}
		n /= 2
		ans++
	}

	return ans
}
