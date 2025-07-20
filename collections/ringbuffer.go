package collections

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

// func (r *RingBuffer[T]) Pop(val T) T {
// 	if r.count <= 0 {
// 		panic("empty ringbuffer")
// 	}
// 	r.end = r.end - 1
// 	sort.Sort()
// }
