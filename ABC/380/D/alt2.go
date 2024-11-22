package main

import (
	"fmt"
	"unicode"
)

// TLE
// K以上の長さになるまでSを２倍にする
// 元々のSの長さになるまで半分にしていく。その際にKの位置を補足し、反転回数をカウントする
func main2() {

	var S string
	fmt.Scan(&S)

	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var K int
		fmt.Scan(&K)

		currentLen := len(S)
		for currentLen < K {
			currentLen = currentLen * 2
		}

		flippedCount := 0
		for currentLen > len(S) {
			half := currentLen / 2

			if K > half {
				flippedCount++
				K = K - half
			}

			currentLen = half
		}

		if i != 0 {
			fmt.Print(" ")
		}

		originalChar := S[K-1]

		if flippedCount%2 != 0 {
			if unicode.IsLower(rune(originalChar)) {
				fmt.Print(string(unicode.ToUpper(rune(originalChar))))
				continue
			} else {
				fmt.Print(string(unicode.ToLower(rune(originalChar))))
				continue
			}
		}
		fmt.Print(string(originalChar))
	}
}
