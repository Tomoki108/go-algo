package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	xs := make([]int, 0, M)
	for i := 0; i < M; i++ {
		var X int
		fmt.Scan(&X)

		xs = append(xs, X)
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})

	as := make([]int, 0, M)
	for i := 0; i < M; i++ {
		var A int
		fmt.Scan(&A)

		as = append(as, A)
	}

	lastEmptyCellNum := N
	for i := M - 1; i >= 0; i-- {
		if xs[i] == lastEmptyCellNum {
			lastEmptyCellNum--
		} else {
			break
		}
	}

	ans := 0

	stoneSum := 0
	for i := 0; i < M; i++ {
		stoneSum += as[i]
	}
	if stoneSum != N {
		ans = -1
		fmt.Println(ans)

		return
	}

	for i := M - 1; i >= 0; i-- {
		cellNum := xs[i]
		stoneNum := as[i]

		if stoneNum > lastEmptyCellNum-cellNum+1 {
			ans = -1
			break
		}

		for j := 0; j < stoneNum; j++ {
			ans += lastEmptyCellNum - cellNum
			lastEmptyCellNum--
		}

	}

	fmt.Println(ans)
}
