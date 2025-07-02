package types

import (
	"sync"
)

// OrderedQueue holds orderable items in sorted order
type OrderedQueue[T any] struct {
	queue   []T
	mu      sync.RWMutex
	compare func(a, b T) int // comparison function: -1 if a < b, 0 if a == b, 1 if a > b
}

// NewOrderedQueue creates a new OrderedQueue with a comparison function
func NewOrderedQueue[T any](compare func(a, b T) int) *OrderedQueue[T] {
	return &OrderedQueue[T]{
		queue:   make([]T, 0),
		compare: compare,
	}
}

// Push adds an item to the queue in the correct sorted position
// If the item already exists, it is ignored
func (q *OrderedQueue[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Find insertion position and check for duplicates
	insertPos := 0
	for i, existing := range q.queue {
		cmp := q.compare(existing, item)
		if cmp == 0 {
			// Item already exists, ignore
			return
		}
		if cmp < 0 {
			// existing < item, continue searching
			insertPos = i + 1
		} else {
			// existing > item, found insertion position
			break
		}
	}

	// Insert item at the correct position
	if insertPos == len(q.queue) {
		// Append to end
		q.queue = append(q.queue, item)
	} else {
		// Insert at position
		q.queue = append(q.queue, item)
		copy(q.queue[insertPos+1:], q.queue[insertPos:len(q.queue)-1])
		q.queue[insertPos] = item
	}
}

// Pop removes and returns the first (smallest) item from the queue
// Returns false if queue is empty
func (q *OrderedQueue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		var zero T
		return zero, false
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, true
}

// Peek returns the first (smallest) item without removing it
// Returns false if queue is empty
func (q *OrderedQueue[T]) Peek() (T, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if len(q.queue) == 0 {
		var zero T
		return zero, false
	}

	return q.queue[0], true
}

// Len returns the number of items in the queue
func (q *OrderedQueue[T]) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return len(q.queue)
}

// IsEmpty returns true if the queue is empty
func (q *OrderedQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

// Contains checks if an item exists in the queue
func (q *OrderedQueue[T]) Contains(item T) bool {
	q.mu.RLock()
	defer q.mu.RUnlock()

	for _, existing := range q.queue {
		if q.compare(existing, item) == 0 {
			return true
		}
	}
	return false
}

// Clear removes all items from the queue
func (q *OrderedQueue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = make([]T, 0)
}

// ToSlice returns a copy of all items in the queue
func (q *OrderedQueue[T]) ToSlice() []T {
	q.mu.RLock()
	defer q.mu.RUnlock()

	result := make([]T, len(q.queue))
	copy(result, q.queue)
	return result
}
