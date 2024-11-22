package main

import (
	"fmt"
	"math/bits"
	"unicode"
)

// TLE
// Sが一文字の場合を考える。操作後の文字列の桁数は、2^Kとなる。
// 桁数を二進数表記する場合、1000は2^3であり、３回目の操作でできたもの。1001~1111は4回目の操作でできたもの。
// 最も右側の1を消すと、反対側の対応する桁数に行ける。ただし1000のような区切りの場合は例外であり、そういった場合は、最も右側の乗数回一番最初のSから反転している。
// このことから、桁数を二進数表記し、最も右側の1の位置を取得する。その場所に行くまでに何回1を消すのかのカウント + 最も右側の1の乗数が反転回数となる。
// 今まで一文字で考えていたが、文字数が増えてもSの塊として考えれば同じことができる。
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

		// Kが何セット目のSに含まれるかを判定
		setNo := quotient
		if remainder > 0 {
			setNo++
		}

		binSetNo := uint64(setNo)
		popCount := bits.OnesCount64(binSetNo)

		firstOneOrder := -1

		// 何回反転するかを判定
		firstOneOrder = bits.Len64(binSetNo & -binSetNo)
		flippedCount := popCount - 1 + firstOneOrder - 1

		charIndex := remainder - 1
		if charIndex == -1 {
			charIndex = len(S) - 1
		}
		originalChar := rune(S[charIndex])

		isUpper := unicode.IsUpper(originalChar)
		if flippedCount%2 != 0 {
			isUpper = !isUpper
		}

		if i != 0 {
			fmt.Print(" ")
		}
		if isUpper {
			// fmt.Printf("K: %d setNo: %d flippedCount: %d popCount: %d binSetNo: %d originalChar: %s isUpper: %t binSetNo: %064b reversedBinSetNo: %064b\n", K, setNo, flippedCount, popCount, binSetNo, string(originalChar), isUpper, binSetNo, reversedBinSetNo)

			fmt.Print(string(unicode.ToUpper(originalChar)))
		} else {
			// fmt.Printf("K: %d setNo: %d flippedCount: %d popCount: %d binSetNo: %d originalChar: %s isUpper: %t binSetNo: %064b reversedBinSetNo: %064b\n", K, setNo, flippedCount, popCount, binSetNo, string(originalChar), isUpper, binSetNo, reversedBinSetNo)

			fmt.Print(string(unicode.ToLower(originalChar)))
		}
	}
}
