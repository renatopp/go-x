package dsx

import (
	"iter"
)

// Queue is a simple queue data structure that supports push and pop operations.
type Queue[T comparable] struct {
	items []T
}

// NewQueue creates a new empty queue.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		items: []T{},
	}
}

// NewQueueFrom creates a new queue from the given items.
func NewQueueFrom[T comparable](items []T) *Queue[T] {
	return &Queue[T]{
		items: items,
	}
}

// Push adds one or more items to the end of the queue.
func (q *Queue[T]) Push(items ...T) {
	q.items = append(q.items, items...)
}

// PushSlice adds multiple items to the end of the queue from a slice. The items
// are added in the order they appear in the slice.
func (q *Queue[T]) PushSlice(items []T) {
	q.items = append(q.items, items...)
}

// Get returns the item at the specified index in the queue. The index is resolved
// using python-like negative indexing. If the index is out of range, it panics.
func (q *Queue[T]) Get(i int) T {
	i = resolveIndex(i, len(q.items))
	if i < 0 || i >= len(q.items) {
		panic("queue index out of range")
	}
	return q.items[i]
}

// GetOr returns the item at the specified index in the queue. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// provided default value.
func (q *Queue[T]) GetOr(i int, v T) T {
	i = resolveIndex(i, len(q.items))
	if i < 0 || i >= len(q.items) {
		return v
	}
	return q.items[i]
}

// GetOk returns the item at the specified index in the queue. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// zero value of T and false. Otherwise, it returns the item and true.
func (q *Queue[T]) GetOk(i int) (T, bool) {
	i = resolveIndex(i, len(q.items))
	if i < 0 || i >= len(q.items) {
		var zero T
		return zero, false
	}
	return q.items[i], true
}

// First  returns the item at the front of the queue without
// removing it. If the queue is empty, it panics.
func (q *Queue[T]) First() T {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	return q.items[0]
}

// FirstOr returns the item at the front of the queue without removing it. If the queue is empty, it returns the provided default value.
func (q *Queue[T]) FirstOr(v T) T {
	if len(q.items) == 0 {
		return v
	}
	return q.items[0]
}

// FirstOk returns the item at the front of the queue without removing it. If the queue is empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (q *Queue[T]) FirstOk() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// Last returns the item at the end of the queue without removing it. If the queue is
// empty, it panics.
func (q *Queue[T]) Last() T {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	return q.items[len(q.items)-1]
}

// LastOr returns the item at the end of the queue without removing it.
// If the queue is empty, it returns the provided default value.
func (q *Queue[T]) LastOr(v T) T {
	if len(q.items) == 0 {
		return v
	}
	return q.items[len(q.items)-1]
}

// LastOk returns the item at the end of the queue without removing it. If the queue is empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (q *Queue[T]) LastOk() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[len(q.items)-1], true
}

// Pop removes and returns the item at the front of the queue. If the queue is empty, it panics.
func (q *Queue[T]) Pop() T {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// PopOr removes and returns the item at the front of the queue. If the queue is empty, it returns the provided default value.
func (q *Queue[T]) PopOr(v T) T {
	if len(q.items) == 0 {
		return v
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// PopOk removes and returns the item at the front of the queue. If the queue is empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (q *Queue[T]) PopOk() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IndexOf returns the index of the first occurrence of the specified item in
// the queue. If the item is not found, it returns -1.
func (q *Queue[T]) IndexOf(item T) int {
	for i, v := range q.items {
		if v == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc returns the index of the first item in the queue that
// satisfies the provided predicate function. If no such item is found, it returns -1.
func (q *Queue[T]) IndexOfFunc(f func(T) bool) int {
	for i, v := range q.items {
		if f(v) {
			return i
		}
	}
	return -1
}

// Contains returns true if the specified item is present in the queue, and
// false otherwise.
func (q *Queue[T]) Contains(item T) bool {
	return q.IndexOf(item) != -1
}

// ContainsFunc returns true if there is an item in the queue that satisfies
// the provided predicate function, and false otherwise.
func (q *Queue[T]) ContainsFunc(f func(T) bool) bool {
	return q.IndexOfFunc(f) != -1
}

// Size returns the number of items currently in the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue, resulting in an empty queue.
func (q *Queue[T]) Clear() {
	q.items = []T{}
}

// Clone creates a new queue that is a copy of the current queue. The items
// in the new queue are the same as those in the original queue, but they are
// stored in a separate slice, so modifications to one queue will not affect
// the other.
func (q *Queue[T]) Clone() *Queue[T] {
	items := make([]T, len(q.items))
	copy(items, q.items)
	return NewQueueFrom(items)
}

// ToSlice returns a slice containing all the items in the queue in the same
// order. Modifying the returned slice will not affect the original queue.
func (q *Queue[T]) ToSlice() []T {
	return q.items
}

// Iter returns an iterator that yields the index and item of each element in
// the queue in order. Starting from the front of the queue (index 0) to the end.
func (q *Queue[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range q.items {
			if !yield(i, item) {
				break
			}
		}
	}
}
