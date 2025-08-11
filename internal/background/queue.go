package background

type Queue[T any] struct {
	que chan T
}

// Enqueue an item
func (q *Queue[T]) Enq(item T) bool {
	select {
	case q.que <- item:
		return true
	default:
		return false
	}
}

// Dequeue an item
func (q *Queue[T]) Deq() (T, bool) {
	item, ok := <-q.que
	return item, ok
}

func (q *Queue[T]) Len() int {
	return len(q.que)
}

func NewQueue[T any](size int64) *Queue[T] {
	var q chan T
	if size != 0 {
		q = make(chan T, size)
	} else {
		q = make(chan T)
	}
	return &Queue[T]{que: q}
}
