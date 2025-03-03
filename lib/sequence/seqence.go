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

// O(|seq|)
// 数列の、全ての連続部分列の総和の総和を返す
func SumOfAllContiguousSubsequences(seq []int) int {
	n := len(seq)
	sum := 0
	for i := 1; i <= n; i++ {
		// 数列のある要素が登場する連続部分列は、左端を自身以前の要素から選ぶ数 * 右端を自身以降の要素から選ぶ数
		sum += seq[i-1] * i * (n - i + 1)
	}
	return sum
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
