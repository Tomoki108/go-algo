package in

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadString(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

func ReadInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	i, _ := strconv.Atoi(input)

	return i
}

func ReadStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = s
	}

	return arr
}

func ReadIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}
