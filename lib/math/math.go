package math

// a割るbの、数学における剰余を返す。
// a = b * Quotient + RemainderとなるRemainderを返す（Quotientは負でもよく、Remainderは常に0以上という制約がある）
// goのa%bだと、|a|割るbの剰余にaの符号をつけて返すため、負の数が含まれる場合数学上の剰余とは異なる。
func Mod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}
