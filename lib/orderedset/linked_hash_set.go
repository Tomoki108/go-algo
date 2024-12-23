package orderedset

import "fmt"

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type LinkedHashSet[T comparable] struct {
	elements map[T]*Node[T]
	head     *Node[T] // GetAllに必要
	tail     *Node[T] // Add, Remove に必要
}

func NewLinkedHashSet[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{
		elements: make(map[T]*Node[T]),
	}
}

// O(1)
func (s *LinkedHashSet[T]) Add(value T) {
	if _, exists := s.elements[value]; exists {
		return // Element already exists, do nothing.
	}
	// Create a new node.
	newNode := &Node[T]{value: value}
	// Add the node to the end of the list.
	if s.tail == nil {
		// First element in the list.
		s.head = newNode
		s.tail = newNode
	} else {
		// Append to the tail.
		s.tail.next = newNode
		newNode.prev = s.tail
		s.tail = newNode
	}
	// Add to the map.
	s.elements[value] = newNode
}

// O(1)
func (s *LinkedHashSet[T]) Remove(value T) {
	node, exists := s.elements[value]
	if !exists {
		return // Element not found, do nothing.
	}
	// Remove the node from the list.
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// Node is the head.
		s.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// Node is the tail.
		s.tail = node.prev
	}
	// Remove from the map.
	delete(s.elements, value)
}

// O(1)
func (s *LinkedHashSet[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// O(N)
func (s *LinkedHashSet[T]) GetAll() []T {
	var result []T
	for node := s.head; node != nil; node = node.next {
		result = append(result, node.value)
	}
	return result
}

func ExampleForLinkedHashSet() {
	// Example usage
	set := NewLinkedHashSet[string]()
	set.Add("apple")
	set.Add("banana")
	set.Add("cherry")
	set.Add("apple") // Duplicate, will not be added

	fmt.Println(set.GetAll()) // Output: [apple banana cherry]

	set.Remove("banana")
	fmt.Println(set.GetAll()) // Output: [apple cherry]

	fmt.Println(set.Contains("apple"))  // Output: true
	fmt.Println(set.Contains("banana")) // Output: false
}
