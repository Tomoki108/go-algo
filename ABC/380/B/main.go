package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	sl := strings.Split(S, "")

	ans := make([]int, 0, len(S))
	numOfBars := 0
	for i := 1; i < len(sl); i++ {
		if sl[i] == "-" {
			numOfBars++
		} else {
			if numOfBars != 0 {
				ans = append(ans, numOfBars)
				numOfBars = 0
			}
		}
	}

	for i, v := range ans {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
