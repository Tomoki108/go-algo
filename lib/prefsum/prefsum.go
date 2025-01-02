package prefsum

// O(n)
// 一次元配列の累積和を返す（index0には0を入れる。）
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
}

// 二次元累積和から、任意の範囲の和を求める
// sumGridには、x, y, z方向に番兵（余分な空の一行）が含まれているものとする
// Lx, Rxは、その軸における範囲指定 => x方向には、Rxの累積和からLx-1の累積和を引く
func SumFrom2DPrefixSum(sumGrid [][]int, Lx, Rx, Ly, Ry int) int {
	Lx--
	Ly--

	// 包除原理
	result := sumGrid[Rx][Ry]

	result -= sumGrid[Lx][Ry]
	result -= sumGrid[Rx][Ly]

	result += sumGrid[Lx][Ly]

	return result
}

// 三次元累積和から、任意の範囲の和を求める
// sumCubには、x, y, z方向に番兵（余分な空の一行）が含まれているものとする
// Lx, Rxは、その軸における範囲指定 => x方向には、Rxの累積和からLx-1の累積和を引く
func SumFrom3DPrefixSum(sumCube [][][]int, Lx, Rx, Ly, Ry, Lz, Rz int) int {
	Lx--
	Ly--
	Lz--

	// 包除原理
	result := sumCube[Rx][Ry][Rz]

	result -= sumCube[Lx][Ry][Rz]
	result -= sumCube[Rx][Ly][Rz]
	result -= sumCube[Rx][Ry][Lz]

	result += sumCube[Lx][Ly][Rz]
	result += sumCube[Lx][Ry][Lz]
	result += sumCube[Rx][Ly][Lz]

	result -= sumCube[Lx][Ly][Lz]

	return result
}
