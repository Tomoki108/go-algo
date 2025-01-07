package math

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
