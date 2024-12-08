package primenum

// エラトステネスの篩でN以下の素数を昇順で列挙する: O(n*log_log_n)
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
