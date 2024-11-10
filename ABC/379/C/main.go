package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	type xa struct {
		x, a int
	}

	xas := make([]xa, 0, M)
	for i := 0; i < M; i++ {
		var X int
		fmt.Scan(&X)

		xas = append(xas, xa{x: X})
	}

	for i := 0; i < M; i++ {
		var A int
		fmt.Scan(&A)

		xas[i].a = A
	}

	sort.Slice(xas, func(i, j int) bool {
		return xas[i].x < xas[j].x
	})

	stoneSum := 0
	for i := 0; i < M; i++ {
		stoneSum += xas[i].a
	}
	if stoneSum != N {
		fmt.Println(-1)

		return
	}

	lastEmptyCellNum := N
	for i := M - 1; i >= 0; i-- {
		if xas[i].x == lastEmptyCellNum {
			lastEmptyCellNum--
		} else {
			break
		}
	}

	if lastEmptyCellNum == 0 {
		fmt.Println(0)
		return
	}

	ans := 0
	for i := M - 1; i >= 0; i-- {
		cellNum := xas[i].x
		stoneNum := xas[i].a

		if cellNum > lastEmptyCellNum {
			continue
		}

		if stoneNum > lastEmptyCellNum-cellNum+1 {
			ans = -1
			break
		}

		ans += ((lastEmptyCellNum - cellNum + (lastEmptyCellNum - stoneNum - cellNum + 1)) * stoneNum) / 2
		lastEmptyCellNum = lastEmptyCellNum - stoneNum
	}

	fmt.Println(ans)
}
