package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var H, W, K int
var grid [][]string
var visited [][]bool
var ans int

func main() {
	fmt.Scanf("%d %d %d", &H, &W, &K)

	grid = make([][]string, H)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < H; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line) // remove '\n' (TrimSpaceは改行も削除してくれる)
		grid[i] = strings.Split(line, "")
	}

	visited = make([][]bool, H)
	for h := 0; h < H; h++ {
		visited[h] = make([]bool, W)
		for w := 0; w < W; w++ {
			visited[h][w] = false
		}
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if grid[h][w] == "." {
				explore(h, w, 0)
			}
		}
	}

	fmt.Println(ans)
}

func explore(i, j, k int) {
	if k == K {
		ans++
		return
	}

	visited[i][j] = true

	for _, delta := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		ni, nj := i+delta[0], j+delta[1]

		if 0 <= ni && ni < H && 0 <= nj && nj < W && grid[ni][nj] == "." && !visited[ni][nj] {
			explore(ni, nj, k+1)
		}

	}

	// ここに到達するのは、ここからはそれ以上もうどこにも進めない場合。訪問済みを解除して、再帰元に戻る
	visited[i][j] = false
}
