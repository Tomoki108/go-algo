package compress

import (
	"strconv"
	"strings"
)

// O(n)
// ランレングス圧縮を行う。[]"数+delimiter+文字種"を返す。
func RunLength(sl []string, delimiter string) []string {
	comp := make([]string, 0, len(sl))
	if len(sl) == 0 {
		return comp
	}

	lastChar := sl[0]
	currentLen := 0
	for i := 0; i < len(sl); i++ {
		s := sl[i]
		if s == lastChar {
			currentLen++
		} else {
			comp = append(comp, strconv.Itoa(currentLen)+delimiter+lastChar)
			lastChar = s
			currentLen = 1
		}
	}
	comp = append(comp, strconv.Itoa(currentLen)+delimiter+lastChar) // 最後の一文字

	return comp
}

// O(1)
// "数+delimiter+文字種"を分割して数と文字種を返す
func SplitRLStr(s, delimiter string) (int, string) {
	strs := strings.Split(s, delimiter)
	num, _ := strconv.Atoi(strs[0])

	return num, strs[1]
}
