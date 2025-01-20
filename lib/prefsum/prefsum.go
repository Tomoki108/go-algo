package prefsum

// O(n)
// 一次元累積和を返す（index0には0を入れる。）
func PrefixSum(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + sl[i]
	}
	return res
}

// O(grid_size)
// 二次元累積和を返す（各次元のindex0には0を入れる。）
func PrefixSum2D(grid [][]int) [][]int {
	H := len(grid) + 1
	W := len(grid[0]) + 1

	sumGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		sumGrid[i] = make([]int, W)

		if i == 0 {
			continue
		}
		copy(sumGrid[i][1:], grid[i-1])
	}

	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			sumGrid[i][j] += sumGrid[i][j-1]
		}
	}
	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			sumGrid[i][j] += sumGrid[i-1][j]
		}
	}
	return sumGrid
}

// O(cube_size)
// 三次元累積和を返す（各次元のindex0には0を入れる。）
func PrefixSum3D(cube [][][]int) [][][]int {
	X := len(cube) + 1
	Y := len(cube[0]) + 1
	Z := len(cube[0][0]) + 1

	sumCube := make([][][]int, X)
	for i := 0; i < X; i++ {
		sumCube[i] = make([][]int, Y)
		for j := 0; j < Y; j++ {
			sumCube[i][j] = make([]int, Z)
		}
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			for k := 0; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j][k-1]
			}
		}
	}
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			for k := 0; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j-1][k]
			}
		}
	}
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			for k := 0; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i-1][j][k]
			}
		}
	}
	return sumCube
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
