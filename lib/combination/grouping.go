package combination

import "github.com/Tomoki108/go-algo/lib/math"

// NOTE: 計算量の見積もり用. 答えが大きくなる引数の場合、オーバーフローすることに注意.
// n人を区別のある丁度k個のグループに分ける場合の数（全射の個数）を求める
func CalcSurjectionNum(n, k int) int {
	sum := 0
	for i := 1; i <= k; i++ {
		sum += pow(-1, k-i) * CombinationNum(k, i) * pow(i, n)
	}
	return sum
}

// NOTE: 計算量の見積もり用. 答えが大きくなる引数の場合、オーバーフローすることに注意.
// n人を区別のない丁度k 個のグループに分ける場合の数（スターリング数）を求める
func CalcStirlingNum(n, k int) int {
	kf := math.Factorial(k)
	sum := 0
	for i := 1; i <= k; i++ {
		sum += pow(-1, k-i) * CombinationNum(k, i) * pow(i, n)
	}

	return sum / kf
}

// NOTE: 計算量の見積もり用. 答えが大きくなる引数の場合、オーバーフローすることに注意.
// n人を区別のないk個以下のグループに分ける場合の数（ベル数）を求める
func CalcBellNum(n, k int) int {
	sum := 0
	for i := 0; i <= k; i++ {
		sum += CalcStirlingNum(n, i)
	}
	return sum
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
