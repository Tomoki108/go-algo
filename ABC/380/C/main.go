package main

import (
	"fmt"
	"strings"
)

type chunk struct {
	startIdx, len int
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var S string
	fmt.Scan(&S)

	sl := strings.Split(S, "")

	chunks := make([]chunk, 0, len(S))

	currentChunkLen := 0
	for i, s := range sl {
		if s == "1" {
			currentChunkLen++

			if i == len(S)-1 {
				chunks = append(chunks, chunk{i - currentChunkLen + 1, currentChunkLen})
			}
		} else {
			if currentChunkLen != 0 {
				chunks = append(chunks, chunk{i - currentChunkLen, currentChunkLen})
				currentChunkLen = 0
			}
		}
	}

	removeIndex := K - 1
	chunks[removeIndex-1].len += chunks[removeIndex].len
	chunks = append(chunks[:removeIndex], chunks[removeIndex+1:]...)

	// chunkMap := make(map[int]int, len(S)) // key: chunk index, value: chunk length

	// for _, c := range chunks {
	// 	chunkMap[c.startIdx] = c.len
	// }

	// fmt.Println(printS(len(S), chunkMap))

	ans := strings.Repeat("0", len(S))
	bytes := []byte(ans)

	for _, v := range chunks {
		copy(bytes[v.startIdx:v.startIdx+v.len], []byte(strings.Repeat("1", v.len)))
	}

	fmt.Println(string(bytes))
}

// func printS(len int, chunkMap map[int]int) string {
// 	s := ""

// 	for i := 0; i < len; i++ {
// 		if len, ok := chunkMap[i]; ok {
// 			s += strings.Repeat("1", len)
// 			i += len - 1
// 			continue
// 		}
// 		s += "0"
// 	}

// 	return s
// }
