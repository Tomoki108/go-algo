package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	l := []int{a, b, c, d}
	wasted := []bool{false, false, false, false}
	ans := 0
	for i := 0; i <= 3; i++ {
		if wasted[i] {
			continue
		}

		for j := i + 1; j <= 3; j++ {
			if wasted[j] {
				continue
			}

			if l[i] == l[j] {
				wasted[i] = true
				wasted[j] = true
				ans++
				break
			}
		}
	}

	fmt.Println(ans)
}
