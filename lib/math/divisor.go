package math

// O(√n)
// 約数を列挙する。戻り値はソートされていないことに注意
func EnumerateDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
