package linkedlist

type NodeBD[T comparable] struct {
	val  T
	prev *NodeBD[T]
	next *NodeBD[T]
}

// O(N)
func (n *NodeBD[T]) Head() *NodeBD[T] {
	for n.prev != nil {
		n = n.prev
	}
	return n
}

// O(N)
func (n *NodeBD[T]) Tail() *NodeBD[T] {
	for n.next != nil {
		n = n.next
	}
	return n
}

// O(1)
func (n *NodeBD[T]) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

// O(1)
func (n *NodeBD[T]) InsertAfter(val T) {
	node := &NodeBD[T]{val: val, prev: n, next: n.next}
	if n.next != nil {
		n.next.prev = node
	}
	n.next = node
}

// O(N)
func CreateBDList[T comparable](sl []T) (head *NodeBD[T], nodeMap map[T]*NodeBD[T]) {
	if len(sl) == 0 {
		return nil, nil
	}

	head = &NodeBD[T]{val: sl[0]}
	nodeMap = make(map[T]*NodeBD[T], len(sl))
	nodeMap[sl[0]] = head

	prev := head
	for i := 1; i < len(sl); i++ {
		node := &NodeBD[T]{val: sl[i], prev: prev}
		prev.next = node
		prev = node
		nodeMap[sl[i]] = node
	}
	return head, nodeMap
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

// O(N)
func (n *Node[T]) Tail() *Node[T] {
	for n.next != nil {
		n = n.next
	}
	return n
}

// O(N)
func (n *Node[T]) Remove() {
	if n.next != nil {
		n.next = n.next.next
	}
}

// O(1)
func (n *Node[T]) InsertAfter(val T) {
	node := &Node[T]{val: val, next: n.next}
	n.next = node
}

// O(N)
func CreateList[T comparable](sl []T) (head *Node[T], nodeMap map[T]*Node[T]) {
	if len(sl) == 0 {
		return nil, nil
	}

	head = &Node[T]{val: sl[0]}
	nodeMap = make(map[T]*Node[T], len(sl))
	nodeMap[sl[0]] = head

	prev := head
	for i := 1; i < len(sl); i++ {
		node := &Node[T]{val: sl[i]}
		prev.next = node
		prev = node
		nodeMap[sl[i]] = node
	}
	return head, nodeMap
}
