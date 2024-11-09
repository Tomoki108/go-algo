package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	type coordinate struct {
		i, j int
	}

	// 駒がある場所及び危険な場所の座標を保持
	cMap := make(map[coordinate]struct{}, M*9)
	for i := 0; i < M; i++ {
		var i, j int
		fmt.Scan(&i, &j)

		// mark danger
		// (i+2,j+1) に置かれている
		// (i+1,j+2) に置かれている
		// (i−1,j+2) に置かれている
		// (i−2,j+1) に置かれている
		// (i−2,j−1) に置かれている
		// (i−1,j−2) に置かれている
		// (i+1,j−2) に置かれている
		// (i+2,j−1) に置かれている
		cMap[coordinate{i, j}] = struct{}{}

		if i+2 <= N && j+1 <= N {
			cMap[coordinate{i + 2, j + 1}] = struct{}{}
		}
		if i+1 <= N && j+2 <= N {
			cMap[coordinate{i + 1, j + 2}] = struct{}{}
		}
		if i-1 > 0 && j+2 <= N {
			cMap[coordinate{i - 1, j + 2}] = struct{}{}
		}
		if i-2 > 0 && j+1 <= N {
			cMap[coordinate{i - 2, j + 1}] = struct{}{}
		}
		if i-2 > 0 && j-1 > 0 {
			cMap[coordinate{i - 2, j - 1}] = struct{}{}
		}
		if i-1 > 0 && j-2 > 0 {
			cMap[coordinate{i - 1, j - 2}] = struct{}{}
		}
		if i+1 <= N && j-2 > 0 {
			cMap[coordinate{i + 1, j - 2}] = struct{}{}
		}
		if i+2 <= N && j-1 > 0 {
			cMap[coordinate{i + 2, j - 1}] = struct{}{}
		}
	}

	ans := N*N - len(cMap)
	fmt.Println(ans)
}
