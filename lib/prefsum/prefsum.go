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
			if i == 0 || j == 0 {
				continue
			}
			copy(sumCube[i][j][1:], cube[i-1][j-1])
		}
	}

	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j][k-1]
			}
		}
	}
	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i][j-1][k]
			}
		}
	}
	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			for k := 1; k < Z; k++ {
				sumCube[i][j][k] += sumCube[i-1][j][k]
			}
		}
	}
	return sumCube
}

// O(1)
// 二次元累積和から、任意の範囲の和を求める
// 左上(i, j) から 右下(k, l) までの範囲の和を求める.
// i, j, k, lには累積和グリッドではなく元々のグリッドのものを渡すこと
func SumFrom2DPrefixSum(sumGrid [][]int, i, j, k, l int) int {
	// k, lは累積和グリッドのindexに合わせるために+1
	// i, jはその一つ左下の範囲の累積和を引きたいのでそのまま
	k++
	l++

	// 包除原理
	result := sumGrid[k][l]

	result -= sumGrid[i][l]
	result -= sumGrid[k][j]

	result += sumGrid[i][j]

	return result
}

// O(1)
// 三次元累積和から、任意の範囲の和を求める.
// 左上(i, j, k) から 右下(l, m, n) までの範囲の和を求める.
// i, j, k, l, m, nには累積和グリッドではなく元々のグリッドのものを渡すこと
func SumFrom3DPrefixSum(sumCube [][][]int, i, j, k, l, m, n int) int {
	// l, m, nは累積和グリッドのindexに合わせるために+1
	// i, j, kはその一つ左下の範囲の累積和を引きたいのでそのまま
	l++
	m++
	n++

	// 包除原理
	result := sumCube[l][m][n]

	result -= sumCube[i][m][n]
	result -= sumCube[l][j][n]
	result -= sumCube[l][m][k]

	result += sumCube[i][j][n]
	result += sumCube[i][m][k]
	result += sumCube[l][j][k]

	result -= sumCube[i][j][k]

	return result
}

// O(n)
// 一次元累積XORを返す（index0には0を入れる。）
func PrefixXOR(sl []int) []int {
	n := len(sl)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] ^ sl[i]
	}
	return res
}
