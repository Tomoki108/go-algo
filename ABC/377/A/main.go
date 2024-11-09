package main

import (
	"fmt"
	"strings"
)

func main() {

	alt()
	// simple version
	// var S string
	// fmt.Scan(&S)
	// sl := strings.Split(S, "")
	//
	// cntA := 0
	// cntB := 0
	// cntC := 0
	// for _, s := range sl {
	// 	if s == "A" {
	// 		cntA++
	// 	} else if s == "B" {
	// 		cntB++
	// 	} else if s == "C" {
	// 		cntC++
	// 	}
	// }

	// if cntA == 1 && cntB == 1 && cntC == 1 {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }
}

// more versatile version
func alt() {
	var S string
	fmt.Scan(&S)
	sl := strings.Split(S, "")

	cntMap := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
	}

	for _, s := range sl {
		_, ok := cntMap[s]
		if ok {
			cntMap[s]++
		}
	}

	ans := "Yes"
	for _, v := range cntMap {
		if v != 1 {
			ans = "No"
			break
		}
	}

	fmt.Println(ans)
}
