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
	totalGrowth := 0 // 遅延評価用の成長値

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < Q; i++ {
		line, _ := reader.ReadString('\n')
		sl := strings.Fields(line)

		if sl[0] == "1" {
			plantHeights = append(plantHeights, -totalGrowth) // 遅延成長分を差し引いて追加
		} else if sl[0] == "2" {
			t, _ := strconv.Atoi(sl[1])
			totalGrowth += t
		} else if sl[0] == "3" {
			h, _ := strconv.Atoi(sl[1])
			h -= totalGrowth // 実際の高さに変換

			sort.Ints(plantHeights)
			idx := sort.Search(len(plantHeights), func(i int) bool {
				return plantHeights[i] >= h
			})
			fmt.Println(len(plantHeights) - idx)
			plantHeights = plantHeights[:idx]
		}
	}
}
