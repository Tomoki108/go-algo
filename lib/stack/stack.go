package stack

import "container/list"

type Stack[T any] struct {
	list *list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: list.New(),
	}
}

func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	back := s.list.Back()
	if back == nil {
		var zero T
		return zero, false
	}
	s.list.Remove(back)
	return back.Value.(T), true
}

// Peek returns the back element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	back := s.list.Back()
	if back == nil {
		var zero T
		return zero, false
	}
	return back.Value.(T), true
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}
