package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// Sが一文字の場合を考える。操作後の文字数は、2^Kとなる。（Kは操作回数）
// 文字数の二進数表記を考える。ただし1文字目を0とする。101は4文字目、111は8文字目。1000は16文字目、1111は32文字目。
// 二進数の桁数が操作回数に対応する。（1文字目以外。）二進数の最も左側の1を消すと反対側の対応する文字に行ける。(直前の操作でできた塊=全体の半分を消すことに等しいから)
// 1文字目を0としているのは、1000のような場合に最も左側の1を消した時、左側の対応する文字（反転元の文字）に行けるようにするため。この場合は0文字目。
// 何回反転元に戻ると一文字目に行けるかどうか（一文字目を何回判定したかどうか） = 二進数のポップカウント。
// またここまで一文字で考えていたが、文字数が増えてもSの塊の何セット目かというように考えれば同じことができる。
func main() {
	defer writer.Flush()

	S := ReadString(reader)
	Q := ReadInt(reader)
	Ks := ReadIntArr(reader)

	for i := 0; i < Q; i++ {
		K := Ks[i]

		quotient := K / len(S)
		remainder := K % len(S)

		setNo := quotient
		if remainder == 0 {
			setNo--
		}

		binSetNo := uint64(setNo)
		pc := bits.OnesCount64(binSetNo)

		originalCharIndex := remainder - 1
		if originalCharIndex == -1 {
			originalCharIndex = len(S) - 1
		}
		originalChar := rune(S[originalCharIndex])

		isUpper := unicode.IsUpper(originalChar)
		if pc%2 != 0 {
			isUpper = !isUpper
		}

		if i != 0 {
			fmt.Fprint(writer, " ")
		}
		if isUpper {
			// fmt.Printf("K: %d setNo: %d pc: %d binSetNo: %064b originalChar: %s isUpper: %t\n", K, setNo, pc, binSetNo, string(originalChar), isUpper)
			fmt.Fprint(writer, string(unicode.ToUpper(originalChar)))
		} else {
			//fmt.Printf("K: %d setNo: %d pc: %d binSetNo: %064b originalChar: %s isUpper: %t\n", K, setNo, pc, binSetNo, string(originalChar), isUpper)
			fmt.Fprint(writer, string(unicode.ToLower(originalChar)))
		}
	}
}

// 一行に1文字のみの入力を読み込む
func ReadString(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
func ReadInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
func ReadStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = s
	}

	return arr
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
func ReadIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}
