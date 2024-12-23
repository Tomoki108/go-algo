package orderedset

import (
	"sort"
)

type OrderedSet[T comparable] struct {
	elements map[T]struct{}
	order    []T
}

func NewOrderedSet[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		elements: make(map[T]struct{}),
		order:    []T{},
	}
}

// O(1)
func (s *OrderedSet[T]) Add(value T) {
	if _, exists := s.elements[value]; !exists {
		s.elements[value] = struct{}{}
		s.order = append(s.order, value)
	}
}

// O(N)
func (s *OrderedSet[T]) Remove(value T) {
	if _, exists := s.elements[value]; exists {
		delete(s.elements, value)
		for i, v := range s.order {
			if v == value {
				// Remove from the order slice
				s.order = append(s.order[:i], s.order[i+1:]...)
				break
			}
		}
	}
}

// O(N)
func (s *OrderedSet[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// GetAll returns all elements in insertion order.
func (s *OrderedSet[T]) GetAll() []T {
	return s.order
}

// Sort reorders the elements in the set according to a custom sort function.
func (s *OrderedSet[T]) Sort(less func(a, b T) bool) {
	sort.Slice(s.order, func(i, j int) bool {
		return less(s.order[i], s.order[j])
	})
}
