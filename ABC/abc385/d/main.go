package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const intMax = 1 << 62
const intMin = -1 << 62

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// ordered setを使った解法
// 家の、x座標ごとのy座標、y座標ごとのx座標について、ordered set（AVLを使った平衡二分木で実装）で管理する
func main() {
	defer w.Flush()

	iarr := readIntArr(r)
	N, M, Sx, Sy := iarr[0], iarr[1], iarr[2], iarr[3]

	houseXYMap := make(map[int]*TreeSet[int], N)
	houseYXMap := make(map[int]*TreeSet[int], N)

	for i := 0; i < N; i++ {
		X, Y := read2Ints(r)
		if _, ok := houseXYMap[X]; !ok {
			houseXYMap[X] = NewTreeSet[int]()
		}
		if _, ok := houseYXMap[Y]; !ok {
			houseYXMap[Y] = NewTreeSet[int]()
		}
		houseXYMap[X].Add(Y)
		houseYXMap[Y].Add(X)
	}

	ans := 0
	current := [2]int{Sx, Sy}
	for i := 0; i < M; i++ {
		sarr := readStrArr(r)
		D := sarr[0]
		CS := sarr[1]
		C, _ := strconv.Atoi(CS)

		var next [2]int

		switch D {
		case "U":
			next = [2]int{current[0], current[1] + C}

			x := current[0]
			fromY := current[1]
			toY := next[1]

			yTree, ok := houseXYMap[x]
			if ok {
				houseYs := yTree.Range(fromY, toY)
				ans += len(houseYs)

				for _, y := range houseYs {
					yTree.Remove(y)
					houseYXMap[y].Remove(x)
				}
			}
		case "D":
			next = [2]int{current[0], current[1] - C}

			x := current[0]
			fromY := next[1]
			toY := current[1]

			yTree, ok := houseXYMap[x]
			if ok {
				houseYs := yTree.Range(fromY, toY)
				ans += len(houseYs)

				for _, y := range houseYs {
					yTree.Remove(y)
					houseYXMap[y].Remove(x)
				}
			}
		case "L":
			next = [2]int{current[0] - C, current[1]}

			y := current[1]
			fromX := next[0]
			toX := current[0]

			xTree, ok := houseYXMap[y]
			if ok {
				houseXs := xTree.Range(fromX, toX)
				ans += len(houseXs)

				for _, x := range houseXs {
					xTree.Remove(x)
					houseXYMap[x].Remove(y)
				}
			}
		case "R":
			next = [2]int{current[0] + C, current[1]}

			y := current[1]
			fromX := current[0]
			toX := next[0]

			xTree, ok := houseYXMap[y]
			if ok {
				houseXs := xTree.Range(fromX, toX)
				ans += len(houseXs)

				for _, x := range houseXs {
					xTree.Remove(x)
					houseXYMap[x].Remove(y)
				}
			}
		}

		current = next
	}

	fmt.Fprintf(w, "%d %d %d\n", current[0], current[1], ans)
}

// 工夫で家の重複カウントを防ぐ解法
// サンタの移動を線分として記録する。
// 家の、x座標ごとのy座標について、スライスで管理する。
// 垂直なサンタの移動線のみ処理し、スライスから家を削除する。
// スライスから残った家の、y座標ごとのx座標のスライスを作成し、水平なサンタの移動線を処理する。
func Alt() {
	defer w.Flush()

	iarr := readIntArr(r)
	N, M, Sx, Sy := iarr[0], iarr[1], iarr[2], iarr[3]

	houseXYMap := make(map[int][]int, N)
	for i := 0; i < N; i++ {
		X, Y := read2Ints(r)
		houseXYMap[X] = append(houseXYMap[X], Y)
	}

	for x, _ := range houseXYMap {
		sort.Ints(houseXYMap[x])
	}

	xPaths := make([][2][2]int, 0, M) // from, to 横の移動
	yPaths := make([][2][2]int, 0, M) // from, to　縦の移動

	current := [2]int{Sx, Sy}
	for i := 0; i < M; i++ {
		sarr := readStrArr(r)
		D := sarr[0]
		CS := sarr[1]
		C, _ := strconv.Atoi(CS)

		var next [2]int

		switch D {
		case "U":
			next = [2]int{current[0], current[1] + C}
			yPaths = append(yPaths, [2][2]int{current, next})
		case "D":
			next = [2]int{current[0], current[1] - C}
			yPaths = append(yPaths, [2][2]int{current, next})
		case "L":
			next = [2]int{current[0] - C, current[1]}
			xPaths = append(xPaths, [2][2]int{current, next})
		case "R":
			next = [2]int{current[0] + C, current[1]}
			xPaths = append(xPaths, [2][2]int{current, next})
		}

		current = next
	}

	count := 0
	for _, yPath := range yPaths {
		from := yPath[0]
		to := yPath[1]

		x := from[0]

		fy := from[1]
		ty := to[1]
		fromY := min(fy, ty)
		toY := max(fy, ty)

		houseYs, ok := houseXYMap[x]
		if !ok {
			continue
		}

		idx1 := sort.Search(len(houseYs), func(i int) bool {
			return houseYs[i] >= fromY
		})
		if idx1 != len(houseYs) {
			idx2 := sort.Search(len(houseYs), func(i int) bool {
				return houseYs[i] > toY
			})

			passedHouses := len(houseYs[idx1:idx2])
			count += passedHouses

			newHouseYs := houseYs[:idx1]
			if idx2 != len(houseYs) {
				newHouseYs = append(newHouseYs, houseYs[idx2:]...)
			}
			houseXYMap[x] = newHouseYs
		}
	}

	housYXMap := make(map[int][]int, N)
	for X, Ys := range houseXYMap {
		for _, Y := range Ys {
			housYXMap[Y] = append(housYXMap[Y], X)
		}
	}
	for y, _ := range housYXMap {
		sort.Ints(housYXMap[y])
	}

	for _, xPath := range xPaths {
		from := xPath[0]
		to := xPath[1]

		y := from[1]

		fx := from[0]
		tx := to[0]
		fromX := min(fx, tx)
		toX := max(fx, tx)

		houseXs, ok := housYXMap[y]
		if !ok {
			continue
		}

		idx1 := sort.Search(len(houseXs), func(i int) bool {
			return houseXs[i] >= fromX
		})
		if idx1 != len(houseXs) {
			idx2 := sort.Search(len(houseXs), func(i int) bool {
				return houseXs[i] > toX
			})

			passedHouses := len(houseXs[idx1:idx2])
			count += passedHouses

			newHouseXs := houseXs[:idx1]
			if idx2 != len(houseXs) {
				newHouseXs = append(newHouseXs, houseXs[idx2:]...)
			}
			housYXMap[y] = newHouseXs
		}
	}

	fmt.Fprintf(w, "%d %d %d\n", current[0], current[1], count)
}

// ////////////
// Libs    //
// ///////////
type Ordered interface {
	~int | ~float64 | ~string
}

// TreeNode represents a node in the AVL tree.
type TreeNode[T Ordered] struct {
	value  T
	left   *TreeNode[T]
	right  *TreeNode[T]
	height int
}

// TreeSet is a generic AVL tree-based set. It is alwways sorted automatically.
// can Add, Remove, Contains in O(log n), GetAll in O(n).
type TreeSet[T Ordered] struct {
	root *TreeNode[T]
	size int
}

func NewTreeSet[T Ordered]() *TreeSet[T] {
	return &TreeSet[T]{}
}

// O(1)
func (t *TreeSet[T]) Height(node *TreeNode[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

// O(log n)
func (t *TreeSet[T]) Add(value T) {
	t.root = t.add(t.root, value)
}

// O(log n)
func (t *TreeSet[T]) Remove(value T) {
	t.root = t.remove(t.root, value)
}

// O(log n)
func (t *TreeSet[T]) Contains(value T) bool {
	return t.contains(t.root, value)
}

// O(n)
func (t *TreeSet[T]) GetAll() []T {
	var result []T
	t.inOrderTraversal(t.root, &result)
	return result
}

// O(k + log n)　k: num of nodes in the range
func (t *TreeSet[T]) Range(min, max T) []T {
	var result []T
	t.rangeQuery(t.root, min, max, &result)
	return result
}

// add recursively adds a value to the AVL tree and balances it.
func (t *TreeSet[T]) add(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		t.size++
		return &TreeNode[T]{value: value, height: 1}
	}
	if value < node.value {
		node.left = t.add(node.left, value)
	} else if value > node.value {
		node.right = t.add(node.right, value)
	} else {
		// Value already exists, do nothing.
		return node
	}

	// Update height and balance the tree.
	node.height = 1 + max(t.Height(node.left), t.Height(node.right))
	return t.balance(node)
}

// remove recursively deletes a value from the AVL tree and balances it.
func (t *TreeSet[T]) remove(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = t.remove(node.left, value)
	} else if value > node.value {
		node.right = t.remove(node.right, value)
	} else {
		// Node to be removed found.
		t.size--
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		// Replace with the in-order successor.
		minNode := t.findMin(node.right)
		node.value = minNode.value
		node.right = t.remove(node.right, minNode.value)
	}

	// Update height and balance the tree.
	node.height = 1 + max(t.Height(node.left), t.Height(node.right))
	return t.balance(node)
}

// findMin finds the node with the smallest value in a subtree.
func (t *TreeSet[T]) findMin(node *TreeNode[T]) *TreeNode[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// contains recursively checks for the existence of a value in the AVL tree.
func (t *TreeSet[T]) contains(node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}
	if value < node.value {
		return t.contains(node.left, value)
	} else if value > node.value {
		return t.contains(node.right, value)
	}
	return true
}

// inOrderTraversal performs an in-order traversal of the AVL tree.
func (t *TreeSet[T]) inOrderTraversal(node *TreeNode[T], result *[]T) {
	if node == nil {
		return
	}
	t.inOrderTraversal(node.left, result)
	*result = append(*result, node.value)
	t.inOrderTraversal(node.right, result)
}

// balance balances the AVL tree.
func (t *TreeSet[T]) balance(node *TreeNode[T]) *TreeNode[T] {
	balanceFactor := t.Height(node.left) - t.Height(node.right)

	// Left heavy
	if balanceFactor > 1 {
		if t.Height(node.left.left) >= t.Height(node.left.right) {
			node = t.rotateRight(node)
		} else {
			node.left = t.rotateLeft(node.left)
			node = t.rotateRight(node)
		}
	}

	// Right heavy
	if balanceFactor < -1 {
		if t.Height(node.right.right) >= t.Height(node.right.left) {
			node = t.rotateLeft(node)
		} else {
			node.right = t.rotateRight(node.right)
			node = t.rotateLeft(node)
		}
	}

	return node
}

// rotateLeft performs a left rotation.
func (t *TreeSet[T]) rotateLeft(node *TreeNode[T]) *TreeNode[T] {
	right := node.right
	node.right = right.left
	right.left = node

	// Update heights
	node.height = 1 + max(t.Height(node.left), t.Height(node.right))
	right.height = 1 + max(t.Height(right.left), t.Height(right.right))
	return right
}

// rotateRight performs a right rotation.
func (t *TreeSet[T]) rotateRight(node *TreeNode[T]) *TreeNode[T] {
	left := node.left
	node.left = left.right
	left.right = node

	// Update heights
	node.height = 1 + max(t.Height(node.left), t.Height(node.right))
	left.height = 1 + max(t.Height(left.left), t.Height(left.right))
	return left
}

// rangeQuery recursively collects all values within the range [min, max].
func (t *TreeSet[T]) rangeQuery(node *TreeNode[T], min, max T, result *[]T) {
	if node == nil {
		return
	}
	if min < node.value {
		t.rangeQuery(node.left, min, max, result)
	}
	if min <= node.value && node.value <= max {
		*result = append(*result, node.value)
	}
	if max > node.value {
		t.rangeQuery(node.right, min, max, result)
	}
}

//////////////
// Helpers  //
/////////////

// 一行に1文字のみの入力を読み込む
func readStr(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

// 一行に1つの整数のみの入力を読み込む
func readInt(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	str := strings.TrimSpace(input)
	i, _ := strconv.Atoi(str)

	return i
}

// 一行に2つの整数のみの入力を読み込む
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])

	return i1, i2
}

// 一行に複数の文字列が入力される場合、スペース区切りで文字列を読み込む
func readStrArr(r *bufio.Reader) []string {
	input, _ := r.ReadString('\n')
	return strings.Fields(input)
}

// 一行に複数の整数が入力される場合、スペース区切りで整数を読み込む
func readIntArr(r *bufio.Reader) []int {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}

// height行の文字列グリッドを読み込む
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readStr(r)
		grid[i] = strings.Split(str, "")
	}

	return grid
}

// 文字列グリッドを出力する
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}

// スライスの中身をスペース区切りで出力する
func writeSlice[T any](w *bufio.Writer, sl []T) {
	vs := make([]any, len(sl))
	for i, v := range sl {
		vs[i] = v
	}
	fmt.Fprintln(w, vs...)
}

// スライスの中身をスペース区切りなしで出力する
func writeSliceWithoutSpace[T any](w *bufio.Writer, sl []T) {
	for idx, v := range sl {
		fmt.Fprint(w, v)
		if idx == len(sl)-1 {
			fmt.Fprintln(w)
		}
	}
}

// スライスの中身を一行づつ出力する
func writeSliceByLine[T any](w *bufio.Writer, sl []T) {
	for _, v := range sl {
		fmt.Fprintln(w, v)
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
