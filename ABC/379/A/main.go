package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	a := N / 100
	b := (N - a*100) / 10
	c := N - a*100 - b*10

	fmt.Println(b*100+c*10+a, c*100+a*10+b)
}
