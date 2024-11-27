package compress

import (
	"strconv"
	"strings"
)

var Delimiter = "_"

// ランレングス圧縮を行う。[]"数+delimiter+文字種"を返す。
func RunLength(sl []string) []string {
	comp := make([]string, 0, len(sl))

	lastChar := ""
	currentLen := 0
	for i := 0; i < len(sl); i++ {
		s := sl[i]

		if i == 0 {
			lastChar = s
			currentLen = 1
			continue
		}

		if s == lastChar {
			currentLen++
			continue
		}

		comp = append(comp, strconv.Itoa(currentLen)+Delimiter+lastChar)
		lastChar = s
		currentLen = 1
	}
	comp = append(comp, strconv.Itoa(currentLen)+Delimiter+lastChar) // 最後の一文字

	return comp
}

// "数+delimiter+文字種"を分割して数と文字種を返す
func SplitRLStr(s string) (int, string) {
	strs := strings.Split(s, Delimiter)
	num, _ := strconv.Atoi(strs[0])

	return num, strs[1]
}
