package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// 数がキー。最後に登場した場所(index + 1)が要素
	appearMap := make(map[int]int, N)

	Bs := make([]int, N)
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)

		if idx, ok := appearMap[A]; ok {
			Bs[i] = idx
		} else {
			Bs[i] = -1
		}

		appearMap[A] = i + 1
	}

	// Bsを空白区切りで一行に出力する
	for i, B := range Bs {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(B)
	}
}
