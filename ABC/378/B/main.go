package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var qs, rs []int
	for i := 0; i < N; i++ {
		var q, r int
		fmt.Scan(&q, &r)
		qs = append(qs, q)
		rs = append(rs, r)
	}

	var Q int
	fmt.Scan(&Q)

	var ts, ds []int
	for i := 0; i < Q; i++ {
		var t, d int
		fmt.Scan(&t, &d)
		ts = append(ts, t)
		ds = append(ds, d)
	}

	// ここまで入力

	type numToDevAndReminder struct {
		numToDev int // 割る数, qi
		reminder int // 余り, ri
	}
	numToDevAndReminders := make([]numToDevAndReminder, N)
	for i := 0; i < N; i++ {
		numToDevAndReminders[i] = numToDevAndReminder{numToDev: qs[i], reminder: rs[i]}
	}
	// fmt.Printf("numToDevAndReminders: %v\n", numToDevAndReminders)

	type garbageInfo struct {
		kind int // tj
		day  int // dj
	}
	garbageInfos := make([]garbageInfo, Q)
	for i := 0; i < Q; i++ {
		garbageInfos[i] = garbageInfo{day: ds[i], kind: ts[i]}
	}
	// fmt.Printf("garbageInfos: %v\n", garbageInfos)

	for _, garbageInfo := range garbageInfos {
		numToDev := numToDevAndReminders[garbageInfo.kind-1].numToDev
		remainder := numToDevAndReminders[garbageInfo.kind-1].reminder

		if garbageInfo.day < remainder {
			fmt.Println(remainder)
			continue
		}

		remainderFromDay := garbageInfo.day % numToDev
		if remainderFromDay == remainder {
			fmt.Println(garbageInfo.day)
			continue
		}

		quotientFromDay := garbageInfo.day / numToDev
		if remainderFromDay < remainder {
			ans := numToDev*(quotientFromDay) + remainder
			fmt.Println(ans)
		} else {
			ans := numToDev*(quotientFromDay+1) + remainder
			fmt.Println(ans)
		}
	}
}
