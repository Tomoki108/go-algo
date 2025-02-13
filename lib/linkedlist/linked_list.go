package linkedlist

type NodeD[T comparable] struct {
	val  T
	prev *NodeD[T]
	next *NodeD[T]
}

// O(N)
func (n *NodeD[T]) Head() *NodeD[T] {
	for n.prev != nil {
		n = n.prev
	}
	return n
}

// O(N)
func (n *NodeD[T]) Tail() *NodeD[T] {
	for n.next != nil {
		n = n.next
	}
	return n
}

// O(1)
func (n *NodeD[T]) Remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

// O(1)
func (n *NodeD[T]) InsertAfter(val T) {
	node := &NodeD[T]{val: val, prev: n, next: n.next}
	if n.next != nil {
		n.next.prev = node
	}
	n.next = node
}

// O(N)
func DoublyLinkedList[T comparable](sl []T) (head *NodeD[T], nodeMap map[T]*NodeD[T]) {
	if len(sl) == 0 {
		return nil, nil
	}

	head = &NodeD[T]{val: sl[0]}
	nodeMap = make(map[T]*NodeD[T], len(sl))
	nodeMap[sl[0]] = head

	prev := head
	for i := 1; i < len(sl); i++ {
		node := &NodeD[T]{val: sl[i], prev: prev}
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
func LinkedList[T comparable](sl []T) (head *Node[T], nodeMap map[T]*Node[T]) {
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
