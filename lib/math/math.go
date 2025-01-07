package math

import "strconv"

// O(n) n: numの桁数
// numの桁数を返す
func GetDigists(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	return digits
}

// O(n) n: numの桁数
// numの各桁の数字を返す
func ToDigits(n int) []int {
	s := strconv.FormatInt(int64(n), 10)
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i] - '0') // （'x'に対応する数字 - '0'に対応する数字）のruneの数字 = x as int
	}
	return digits
}
