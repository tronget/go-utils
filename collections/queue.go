package collections

import "fmt"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Dequeue() T {
	if res, err := q.TryDequeue(); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (q *Queue[T]) Peek() T {
	if res, err := q.TryPeek(); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (q *Queue[T]) Clear() {
	var zero T
	for i := range q.data {
		q.data[i] = zero
	}
	q.data = q.data[:0]
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) TryDequeue() (T, error) {
	var zero T
	if q.Empty() {
		return zero, fmt.Errorf("cannot Dequeue: queue is empty")
	}
	res := q.data[0]
	q.data[0] = zero
	q.data = q.data[1:]
	return res, nil
}

func (q *Queue[T]) TryPeek() (T, error) {
	var zero T
	if q.Empty() {
		return zero, fmt.Errorf("cannot Peek: queue is empty")
	}
	return q.data[0], nil
}
