package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 月毎に、毎年の最高気温、最低気温を管理。それを回帰分析する。
func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	type Missing struct {
		year  float64
		month int
		isMax bool
	}
	missings := make([]Missing, 0, n)

	maxYears := make([][]float64, 13)
	maxVals := make([][]float64, 13)
	minYears := make([][]float64, 13)
	minVals := make([][]float64, 13)

	reader.ReadString('\n')
	for i := 0; i < n; i++ {
		raw, _ := reader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(raw), "\t")
		if len(parts) < 4 {
			continue
		}
		yr, _ := strconv.Atoi(parts[0])
		mo := monthStringToInt(parts[1])
		fy := float64(yr)

		tmax, tmin := parts[2], parts[3]

		if strings.HasPrefix(tmax, "Missing_") {
			missings = append(missings, Missing{fy, mo, true})
		} else {
			v, _ := strconv.ParseFloat(tmax, 64)
			maxYears[mo] = append(maxYears[mo], fy)
			maxVals[mo] = append(maxVals[mo], v)
		}

		if strings.HasPrefix(tmin, "Missing_") {
			missings = append(missings, Missing{fy, mo, false})
		} else {
			v, _ := strconv.ParseFloat(tmin, 64)
			minYears[mo] = append(minYears[mo], fy)
			minVals[mo] = append(minVals[mo], v)
		}
	}

	slopeMax := make([]float64, 13)
	interceptMax := make([]float64, 13)
	slopeMin := make([]float64, 13)
	interceptMin := make([]float64, 13)
	for m := 1; m <= 12; m++ {
		if len(maxYears[m]) > 1 {
			s, c := linearReg(maxYears[m], maxVals[m])
			slopeMax[m], interceptMax[m] = s, c
		}
		if len(minYears[m]) > 1 {
			s, c := linearReg(minYears[m], minVals[m])
			slopeMin[m], interceptMin[m] = s, c
		}
	}

	for _, mm := range missings {
		var pred float64
		if mm.isMax {
			pred = slopeMax[mm.month]*mm.year + interceptMax[mm.month]
		} else {
			pred = slopeMin[mm.month]*mm.year + interceptMin[mm.month]
		}
		fmt.Fprintln(writer, pred)
	}
}

func linearReg(xs, ys []float64) (slope, intercept float64) {
	n := float64(len(xs))
	var sx, sy float64
	for i := range xs {
		sx += xs[i]
		sy += ys[i]
	}
	mx := sx / n
	my := sy / n
	var vx, cov float64
	for i := range xs {
		dx := xs[i] - mx
		dy := ys[i] - my
		vx += dx * dx
		cov += dx * dy
	}
	slope = cov / vx
	intercept = my - slope*mx
	return
}

func monthStringToInt(m string) int {
	switch m {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	}
	return 0
}
