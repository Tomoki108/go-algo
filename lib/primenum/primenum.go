package primenum

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

// O(n*log_log_n)
// エラトステネスの篩でN以下の素数を昇順で列挙する
func Eratos(n int) []int {
	if n < 2 {
		return []int{}
	}

	// Create a boolean slice to track prime numbers
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Mark non-prime numbers
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect prime numbers
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
