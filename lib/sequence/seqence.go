package sequence

// 等差数列の和を返す
// O(1)
func ArithmeticSequenceSum(seq []int) int {
	n := len(seq)
	if n == 0 {
		return 0
	}

	return n * (seq[0] + seq[n-1]) / 2
}

// 等比数列の和を返す（公比が整数の前提）
// O(1)
func GeometricSequenceSum(seq []int) int {
	n := len(seq)
	if n == 0 {
		return 0
	}

	if seq[0] == 0 {
		return 0
	}

	if seq[0] == 1 {
		return n
	}

	// 公式：a * (r^n - 1) / (r - 1)
	// a: 初項, r: 公比, n: 項数

	first := seq[0]
	ratio := seq[1] / seq[0]
	return first * (pow(ratio, n) - 1) / (ratio - 1)
}

// NOTE: template.goにも存在する関数
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
