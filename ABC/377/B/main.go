package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := make([][]string, 8)
	for i := 0; i < 8; i++ {
		grid[i] = make([]string, 8)
	}

	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 8; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		grid[i] = strings.Split(line, "")
	}

	dangerMap := make([][]bool, 8)
	for i := range dangerMap {
		dangerMap[i] = make([]bool, 8)
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == "#" {
				dangerMap[i] = []bool{true, true, true, true, true, true, true, true}

				for idx := range grid {
					dangerMap[idx][j] = true
				}
			}
		}
	}

	ans := 0
	for _, row := range dangerMap {
		for _, cell := range row {
			if !cell {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
