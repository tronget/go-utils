package collections

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() T {
	res, err := s.TryPop()
	if err != nil {
		panic(err)
	}
	return res
}

func (s *Stack[T]) Peek() T {
	res, err := s.TryPeek()
	if err != nil {
		panic(err)
	}
	return res
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Clear() {
	var zero T
	for i := range s.data {
		s.data[i] = zero
	}
	s.data = s.data[:0]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) TryPop() (T, error) {
	var zero T
	if s.Empty() {
		return zero, fmt.Errorf("cannot Pop: stack is empty")
	}
	last := len(s.data) - 1
	res := s.data[last]
	s.data[last] = zero
	s.data = s.data[:last]
	return res, nil
}

func (s *Stack[T]) TryPeek() (T, error) {
	if s.Empty() {
		var zero T
		return zero, fmt.Errorf("cannot Peek: stack is empty")
	}
	return s.data[len(s.data)-1], nil
}
