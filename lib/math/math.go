package math

func GetDigists(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	return digits
}
