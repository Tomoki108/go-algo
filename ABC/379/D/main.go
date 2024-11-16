package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	plantHeights := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < Q; i++ {

		line, _ := reader.ReadString('\n')
		sl := strings.Fields(line)

		if sl[0] == "1" {
			plantHeights = append(plantHeights, 0)
		} else if sl[0] == "2" {
			t, _ := strconv.Atoi(sl[1])

			waited := make([]int, 0, len(plantHeights))
			for _, height := range plantHeights {
				waited = append(waited, height+t)
			}
			plantHeights = waited
		} else if sl[0] == "3" {
			h, _ := strconv.Atoi(sl[1])

			sort.Slice(plantHeights, func(i, j int) bool {
				return plantHeights[i] < plantHeights[j]
			})

			idx := sort.Search(len(plantHeights), func(i int) bool {
				return plantHeights[i] >= h
			})
			fmt.Println(len(plantHeights) - idx)

			plantHeights = plantHeights[:idx]
		}
	}
}
