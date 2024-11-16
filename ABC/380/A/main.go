package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	sl := make([]int, 0, 6)

	tmp := N

	sixth := N / 100000
	tmp -= sixth * 100000
	sl = append(sl, sixth)

	fifth := tmp / 10000
	tmp -= fifth * 10000
	sl = append(sl, fifth)

	fourth := tmp / 1000
	tmp -= fourth * 1000
	sl = append(sl, fourth)

	third := tmp / 100
	tmp -= third * 100
	sl = append(sl, third)

	second := tmp / 10
	tmp -= second * 10
	sl = append(sl, second)

	first := tmp
	sl = append(sl, first)

	count1 := 0
	count2 := 0
	count3 := 0

	for _, v := range sl {
		if v == 1 {
			count1++
		} else if v == 2 {
			count2++
		} else if v == 3 {
			count3++
		}
	}

	if count1 == 1 && count2 == 2 && count3 == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
