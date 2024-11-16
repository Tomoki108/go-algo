package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	var Q int
	fmt.Scan(&Q)

	Ks := make([]int, 0, Q)
	for i := 0; i < Q; i++ {
		var K int
		fmt.Scan(&K)
		Ks = append(Ks, K)
	}

	// lenS := big.NewInt(int64(len(S)))
	// numToMultiply := new(big.Int).Exp(big.NewInt(10), big.NewInt(100), nil)
	// finalLegth := new(big.Int).Mul(lenS, numToMultiply)

	// // remainder := new(big.Int).Mod(finalLegth, lenS)

	// ABQ
	// ABQ abq
	// ABQ abq abq ABQ
	// ABQ abq abq ABQ abq ABQ ABQ abq
	// ABQ abq abq ABQ abq ABQ ABQ abq abq ABQ ABQ abq ABQ abq abq ABQ
	// ABQ abq abq ABQ abq ABQ ABQ abq abq ABQ ABQ abq ABQ abq abq ABQ

	s := "ABQ"
	sb := []byte(s)

	// sbを全部小文字に変更する
	for i := 0; i < len(sb); i++ {
		sb[i] = sb[i] + 32
	}
	fmt.Println(string(sb))

	// ABQ
	// ABQ abq(一回反転)
	// ABQ abq(一回反転) abq(一回反転) ABQ(二回反転)
	// ABQ abq(一回反転) abq(一回反転) ABQ(二回反転) abq(一回反転) ABQ(二回反転) ABQ(二回反転) abq(三回反転)
	// ABQ abq(一回反転) abq(一回反転) ABQ(二回反転) abq ABQ ABQ abq abq ABQ ABQ abq ABQ abq abq ABQ

}
