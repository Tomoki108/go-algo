package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	type xa struct {
		x, a int // x: マス目番号, a: 石の数
	}

	xas := make([]xa, M)
	for i := 0; i < M; i++ {
		var X int
		fmt.Scan(&X)

		xas[i].x = X
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
	cost := 0
	for _, xa := range xas {
		// check enogh stones exist and placed properly
		if stoneSum < xa.x-1 {
			fmt.Println(-1)
			return
		}

		// for i := 1; i <= xa.a; i++ {
		// 	//右辺：「石が何番目の升目に最終的に行くことになるか」-「現在の升目の番号」
		// 	cost += (stoneSum + i) - xa.x
		// }

		cost += ((stoneSum + 1 - xa.x) + (stoneSum - xa.x + xa.a)) * xa.a / 2

		stoneSum += xa.a
	}

	if stoneSum != N {
		fmt.Println(-1)
		return
	}

	fmt.Println(cost)
}
