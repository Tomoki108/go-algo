package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// 数がキー。登場した場所が要素
	appearMap := make(map[int][]int, N)

	var As []int
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		As = append(As, A)

		appearMap[A] = append([]int{i}, appearMap[A]...)
	}
	fmt.Printf("appearMap: %v\n", appearMap)

	Bs := make([]int, N)
	for i := 0; i < N; i++ {
		found := false

		for _, appear := range appearMap[As[i]] {
			if appear < i {
				Bs[i] = appear + 1
				found = true
				break
			}
		}

		if !found {
			Bs[i] = -1
		}
	}

	// Bsを空白区切りで一行に出力する
	for i, B := range Bs {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(B)
	}
}
