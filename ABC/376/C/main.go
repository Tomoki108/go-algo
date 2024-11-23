package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	ReadInt(r)          // Nは使わないが行を進めるために読み込む
	As := ReadIntArr(r) // おもちゃサイズ
	Bs := ReadIntArr(r) // 箱サイズ
	// 降順ソート
	sort.Slice(As, func(i, j int) bool {
		return As[i] > As[j]
	})
	sort.Slice(Bs, func(i, j int) bool {
		return Bs[i] > Bs[j]
	})

	var notMatched int // 箱が見つからないおもちゃの大きさ
	var numOfNotMatched = 0

	for i, a := range As {
		if i == len(As)-1 && numOfNotMatched == 0 {
			notMatched = a
			numOfNotMatched++
			break
		}

		if Bs[i-numOfNotMatched] < a {
			notMatched = a
			numOfNotMatched++
			if numOfNotMatched == 2 {
				break
			}
		}

	}

	if numOfNotMatched == 2 {
		fmt.Fprint(w, "-1")
		return
	}

	fmt.Fprint(w, notMatched)
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
