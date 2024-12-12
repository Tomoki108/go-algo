package prefsum

func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
}
