package main

import (
	"fmt"
	"strconv"
)

// f counts the number of "heavy numbers" less than or equal to x.
func f(x int64) int64 {
	s := strconv.FormatInt(x, 10)
	n := len(s)

	// dp[j][strict][nonzero] tracks the number of valid numbers:
	// - j: the largest digit seen so far
	// - strict: whether the current number is still constrained by the input's digits
	// - nonzero: whether any non-zero digit has been used
	dp := make([][][]int64, 10)
	for i := range dp {
		dp[i] = make([][]int64, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 2)
		}
	}

	// Initialize DP for the first digit
	dp[0][1][0] = 1
	for i := 1; i < int(s[0]-'0'); i++ {
		dp[i][1][1] = 1
	}
	dp[s[0]-'0'][0][1] = 1

	// Process subsequent digits
	for i := 1; i < n; i++ {
		// dpn is the next state of dp
		dpn := make([][][]int64, 10)
		for j := range dpn {
			dpn[j] = make([][]int64, 2)
			for k := range dpn[j] {
				dpn[j][k] = make([]int64, 2)
			}
		}

		// Transition dp -> dpn
		for largestDigit := 0; largestDigit < 10; largestDigit++ {
			for isStrict := 0; isStrict < 2; isStrict++ {
				for hasNonZero := 0; hasNonZero < 2; hasNonZero++ {
					for currentDigit := 0; currentDigit < 10; currentDigit++ {
						// Skip invalid transitions
						if hasNonZero == 1 && largestDigit <= currentDigit {
							continue
						}
						if isStrict == 0 && currentDigit > int(s[i]-'0') {
							continue
						}

						// Determine the next state
						nextStrict := isStrict
						if currentDigit < int(s[i]-'0') {
							nextStrict = 1
						}
						nextNonZero := hasNonZero
						if currentDigit > 0 {
							nextNonZero = 1
						}
						nextLargestDigit := largestDigit
						if hasNonZero == 0 && currentDigit != 0 {
							nextLargestDigit = currentDigit
						}

						// Update the next state
						dpn[nextLargestDigit][nextStrict][nextNonZero] += dp[largestDigit][isStrict][hasNonZero]
					}
				}
			}
		}
		dp = dpn
	}

	// Sum up all valid numbers
	result := int64(0)
	for largestDigit := 0; largestDigit < 10; largestDigit++ {
		for isStrict := 0; isStrict < 2; isStrict++ {
			result += dp[largestDigit][isStrict][1] // Only consider states with at least one non-zero digit
		}
	}

	return result
}

func main() {
	var l, r int64
	fmt.Scan(&l, &r)
	// Output the number of heavy numbers in the range [l, r]
	fmt.Println(f(r) - f(l-1))
}
