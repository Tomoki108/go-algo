package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	ints := ReadIntArr(reader)
	N := ints[0]
	M := ints[1]

	var start = 1
	edges := make(map[int][]int, M)
	m := make(map[int]bool, N) // そのノードからstartへ行けるかどうか

	for i := 0; i < M; i++ {
		ints := ReadIntArr(reader)
		from := ints[0]
		to := ints[1]

		edges[from] = append(edges[from], to)
		if to == start {
			m[from] = true
		}
	}

	visited := make(map[int]bool) // 訪問済みノードを記録

	queue := list.New() // BFSのためのキュー
	type elm struct {   // queueに入れる要素
		node  int
		depth int
	}

	queue.PushBack(elm{node: start, depth: 0})
	visited[start] = true

	foundDepth := 0
	// 頂点1からBFSし、頂点1に行けるノードを見つける
Outer:
	for queue.Len() > 0 {
		current := queue.Front().Value.(elm)
		queue.Remove(queue.Front())

		dests := edges[current.node]
		for _, dest := range dests {
			if visited[dest] {
				continue
			}
			visited[dest] = true

			if m[dest] {
				foundDepth = current.depth + 1
				break Outer
			}

			queue.PushBack(elm{node: dest, depth: current.depth + 1})
		}
	}

	if foundDepth > 0 {
		fmt.Fprint(writer, foundDepth+1) // 頂点1に行けるノードの深さ+1が閉路の辺数
	} else {
		fmt.Fprint(writer, -1)
	}
}

//////////////
// Hepers  //
/////////////

// 一行に1文字のみの入力を読み込む
func ReadString(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
func ReadInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
func ReadStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = s
	}

	return arr
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
func ReadIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}
