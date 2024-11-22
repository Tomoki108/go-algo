package main

import (
	"fmt"
	"math/bits"
	"unicode"
)

func main1() {
	var S string
	fmt.Scan(&S)

	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var K int
		fmt.Scan(&K)

		quotient := K / len(S)
		remainder := K % len(S)

		setNo := remainder - 1
		if setNo == -1 {
			setNo = len(S)
		}

		var numOfOperation int
		if quotient == 0 {
			numOfOperation = 0

			fmt.Println("K:", K, "quotient:", quotient, "remainder:", remainder, "setNo:", setNo, "numOfOperation:", numOfOperation)
		} else {
			numOfOperation = bits.Len64(uint64(setNo)) - 1

			bitsWithoutMostLeft := uint64(setNo) << (64 - bits.Len64(uint64(setNo)) + 1)
			pc := bits.OnesCount64(bitsWithoutMostLeft)
			if pc > 0 {
				numOfOperation += 1
			}

			fmt.Println("K:", K, "quotient:", quotient, "remainder:", remainder, "setNo:", setNo, "numOfOperation:", numOfOperation)
		}

		var letterIndex int
		if K <= len(S) {
			letterIndex = K - 1
		} else {
			r := K % len(S)
			if r == 0 {
				letterIndex = len(S) - 1
			} else {
				letterIndex = r - 1
			}
		}

		// fmt.Println("letterIndex:", letterIndex, "original char:", string(rune(S[letterIndex])))

		isUpper := unicode.IsUpper(rune(S[letterIndex]))
		if numOfOperation%2 != 0 {
			isUpper = !isUpper
		}

		if i != 0 {
			fmt.Print(" ")
		}
		if isUpper {
			fmt.Print(string(unicode.ToUpper(rune(S[letterIndex]))))
		} else {
			fmt.Print(string(unicode.ToLower(rune(S[letterIndex]))))
		}

	}
}
