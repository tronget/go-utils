package collections

import (
	"fmt"
)

type RingBuffer[T any] struct {
	data       []T
	start, end int
	size       int
	count      int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		data: make([]T, capacity),
		size: capacity,
	}
}

func (r *RingBuffer[T]) Push(val T) {

	if r.size == r.count {
		r.start = (r.start + 1) % r.size
	} else {
		r.count++
	}
	r.data[r.end] = val
	r.end = (r.end + 1) % r.size
}

func (r *RingBuffer[T]) Pop() T {
	if res, err := r.TryPop(); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (r *RingBuffer[T]) TryPop() (T, error) {
	var zero T
	if r.count <= 0 {
		return zero, fmt.Errorf("cannot Pop: ringbuffer is empty")
	}
	res := r.data[r.start]
	r.data[r.start] = zero
	r.start = (r.start + 1) % r.size
	r.count--
	return res, nil
}

func (r *RingBuffer[T]) Peek() T {
	if res, err := r.TryPeek(); err != nil {
		panic(err)
	} else {
		return res
	}
}

func (r *RingBuffer[T]) TryPeek() (T, error) {
	var zero T
	if r.count <= 0 {
		return zero, fmt.Errorf("cannot Peek: ringbuffer is empty")
	}
	return r.data[r.start], nil
}

func (r *RingBuffer[T]) Full() bool {
	return r.count == r.size
}

func (r *RingBuffer[T]) Empty() bool {
	return r.count == 0
}

func (r *RingBuffer[T]) Clear() {
	r.start, r.end = 0, 0
	r.count = 0
	var zero T
	for i := range r.data {
		r.data[i] = zero
	}
}

func (r *RingBuffer[T]) Size() int {
	return r.count
}
