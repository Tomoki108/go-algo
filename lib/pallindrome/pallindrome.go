package pallindrome

// O(n/2 * (n-k+1))
func ContainsPallindrome(ss []string, k int) bool {
	for i := 0; i <= len(ss)-k; i++ { // k文字ずつチェック、インデックスを一個ずつずらす
		toCheck := ss[i : i+k]
		if IsPallindrome(toCheck) {
			return true
		}
	}

	return false
}

// O(n/2)
func IsPallindrome(ss []string) bool {
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}

	return true
}
