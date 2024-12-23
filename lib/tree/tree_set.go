package tree

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

// MultiSet is a multiset that allows duplicate elements, implemented with a TreeSet.
type MultiSet[T Ordered] struct {
	elements map[T]int // Value to count mapping
	order    *TreeSet[T]
}

func NewMultiSet[T Ordered]() *MultiSet[T] {
	return &MultiSet[T]{
		elements: make(map[T]int),
		order:    NewTreeSet[T](),
	}
}

// O(log N) if new value
// O(1) if already exists
func (ms *MultiSet[T]) Add(value T) {
	if ms.elements[value] == 0 {
		ms.order.Add(value)
	}
	ms.elements[value]++
}

// O(log N) if last value of the value
// O(1) if not last value of the value
func (ms *MultiSet[T]) Remove(value T) {
	if ms.elements[value] == 0 {
		return
	}
	ms.elements[value]--
	if ms.elements[value] == 0 {
		delete(ms.elements, value)
		ms.order.Remove(value)
	}
}

// O(1)
func (ms *MultiSet[T]) Contains(value T) bool {
	return ms.elements[value] > 0
}

// O(1)
func (ms *MultiSet[T]) Count(value T) int {
	return ms.elements[value]
}

// O(k + log n) n: the num of unique elements, k: the total num of elements in the range.
func (ms *MultiSet[T]) Range(min, max T) []T {
	uniqueValues := ms.order.Range(min, max) // Get unique values in range from TreeSet.
	var result []T
	for _, value := range uniqueValues {
		count := ms.Count(value)
		for i := 0; i < count; i++ {
			result = append(result, value)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
