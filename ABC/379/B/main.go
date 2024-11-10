package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	var S string
	fmt.Scan(&N, &K)
	fmt.Scan(&S)

	sl := strings.Split(S, "")

	ans := 0
	for i := 0; i < N; i++ {
		if sl[i] != "O" {
			continue
		}

		if i+K > N {
			break
		}

		ok := true
		for j := i; j < i+K; j++ {
			if sl[j] == "X" {
				ok = false
				break
			}
		}

		if ok {
			ans++

			for j := i; j < i+K; j++ {
				sl[j] = "X"
			}
		}
	}

	fmt.Println(ans)
}
