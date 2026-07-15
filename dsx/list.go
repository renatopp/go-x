package dsx

import (
	"iter"
	"slices"
)

// List is a general-purpose dynamic array that supports random access,
// insertion, and removal at arbitrary indices, in addition to the standard
// Container operations.
type List[T comparable] struct {
	items []T
}

// Compile-time guarantee that List implements the Container interface.
var _ Container[int] = (*List[int])(nil)

// NewList creates a new empty list.
func NewList[T comparable]() *List[T] {
	return &List[T]{
		items: []T{},
	}
}

// NewListFrom creates a new list from the given items.
func NewListFrom[T comparable](items []T) *List[T] {
	return &List[T]{
		items: items,
	}
}

// Push adds one or more items to the end of the list.
func (l *List[T]) Push(items ...T) {
	l.items = append(l.items, items...)
}

// PushSlice adds multiple items to the end of the list from a slice. The items
// are added in the order they appear in the slice.
func (l *List[T]) PushSlice(items []T) {
	l.items = append(l.items, items...)
}

// Get returns the item at the specified index in the list. The index is resolved
// using python-like negative indexing. If the index is out of range, it panics.
func (l *List[T]) Get(i int) T {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		panic("list index out of range")
	}
	return l.items[i]
}

// GetOr returns the item at the specified index in the list. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// provided default value.
func (l *List[T]) GetOr(i int, v T) T {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		return v
	}
	return l.items[i]
}

// GetOk returns the item at the specified index in the list. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// zero value of T and false. Otherwise, it returns the item and true.
func (l *List[T]) GetOk(i int) (T, bool) {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		var zero T
		return zero, false
	}
	return l.items[i], true
}

// Set replaces the item at the specified index in the list. The index is resolved
// using python-like negative indexing. If the index is out of range, it panics.
func (l *List[T]) Set(i int, v T) {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		panic("list index out of range")
	}
	l.items[i] = v
}

// Insert inserts one or more items at the specified index in the list, shifting
// the subsequent items to the right. The index is resolved using python-like
// negative indexing and clamped to the range [0, Size()], so out-of-range
// indices insert at the closest boundary instead of panicking.
func (l *List[T]) Insert(i int, items ...T) {
	i = resolveIndex(i, len(l.items))
	if i < 0 {
		i = 0
	}
	l.items = slices.Insert(l.items, i, items...)
}

// RemoveAt removes and returns the item at the specified index in the list,
// shifting the subsequent items to the left. The index is resolved using
// python-like negative indexing. If the index is out of range, it panics.
func (l *List[T]) RemoveAt(i int) T {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		panic("list index out of range")
	}
	item := l.items[i]
	l.items = slices.Delete(l.items, i, i+1)
	return item
}

// RemoveAtOk removes and returns the item at the specified index in the list,
// shifting the subsequent items to the left. The index is resolved using
// python-like negative indexing. If the index is out of range, it returns the
// zero value of T and false. Otherwise, it returns the item and true.
func (l *List[T]) RemoveAtOk(i int) (T, bool) {
	i = resolveIndex(i, len(l.items))
	if i < 0 || i >= len(l.items) {
		var zero T
		return zero, false
	}
	item := l.items[i]
	l.items = slices.Delete(l.items, i, i+1)
	return item, true
}

// Remove removes the first occurrence of the specified item from the list. It
// returns true if an item was removed, and false if the item was not found.
func (l *List[T]) Remove(item T) bool {
	i := l.IndexOf(item)
	if i == -1 {
		return false
	}
	l.items = slices.Delete(l.items, i, i+1)
	return true
}

// RemoveFunc removes all items from the list that satisfy the provided predicate
// function. It returns the number of items removed.
func (l *List[T]) RemoveFunc(f func(T) bool) int {
	before := len(l.items)
	l.items = slices.DeleteFunc(l.items, f)
	return before - len(l.items)
}

// First returns the item at the start of the list. If the list is empty, it panics.
func (l *List[T]) First() T {
	if len(l.items) == 0 {
		panic("list is empty")
	}
	return l.items[0]
}

// FirstOr returns the item at the start of the list. If the list is empty, it
// returns the provided default value.
func (l *List[T]) FirstOr(v T) T {
	if len(l.items) == 0 {
		return v
	}
	return l.items[0]
}

// FirstOk returns the item at the start of the list. If the list is empty, it
// returns the zero value of T and false. Otherwise, it returns the item and true.
func (l *List[T]) FirstOk() (T, bool) {
	if len(l.items) == 0 {
		var zero T
		return zero, false
	}
	return l.items[0], true
}

// Last returns the item at the end of the list. If the list is empty, it panics.
func (l *List[T]) Last() T {
	if len(l.items) == 0 {
		panic("list is empty")
	}
	return l.items[len(l.items)-1]
}

// LastOr returns the item at the end of the list. If the list is empty, it
// returns the provided default value.
func (l *List[T]) LastOr(v T) T {
	if len(l.items) == 0 {
		return v
	}
	return l.items[len(l.items)-1]
}

// LastOk returns the item at the end of the list. If the list is empty, it
// returns the zero value of T and false. Otherwise, it returns the item and true.
func (l *List[T]) LastOk() (T, bool) {
	if len(l.items) == 0 {
		var zero T
		return zero, false
	}
	return l.items[len(l.items)-1], true
}

// Pop removes and returns the item at the end of the list. If the list is
// empty, it panics.
func (l *List[T]) Pop() T {
	if len(l.items) == 0 {
		panic("list is empty")
	}
	item := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return item
}

// PopOr removes and returns the item at the end of the list. If the list is
// empty, it returns the provided default value.
func (l *List[T]) PopOr(v T) T {
	if len(l.items) == 0 {
		return v
	}
	item := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return item
}

// PopOk removes and returns the item at the end of the list. If the list is
// empty, it returns the zero value of T and false. Otherwise, it returns the
// item and true.
func (l *List[T]) PopOk() (T, bool) {
	if len(l.items) == 0 {
		var zero T
		return zero, false
	}
	item := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return item, true
}

// IndexOf returns the index of the first occurrence of the specified item in
// the list, or -1 if the item is not found.
func (l *List[T]) IndexOf(item T) int {
	for i, v := range l.items {
		if v == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc returns the index of the first item in the list that satisfies
// the provided predicate function, or -1 if no such item is found.
func (l *List[T]) IndexOfFunc(f func(T) bool) int {
	for i, v := range l.items {
		if f(v) {
			return i
		}
	}
	return -1
}

// Contains returns true if the specified item is present in the list, and
// false otherwise.
func (l *List[T]) Contains(item T) bool {
	return l.IndexOf(item) != -1
}

// ContainsFunc returns true if there is an item in the list that satisfies the
// provided predicate function, and false otherwise.
func (l *List[T]) ContainsFunc(f func(T) bool) bool {
	return l.IndexOfFunc(f) != -1
}

// Size returns the number of items currently in the list.
func (l *List[T]) Size() int {
	return len(l.items)
}

// Clear removes all items from the list, leaving it empty.
func (l *List[T]) Clear() {
	l.items = []T{}
}

// Clone creates and returns a new list that is a copy of the current list.
func (l *List[T]) Clone() *List[T] {
	items := make([]T, len(l.items))
	copy(items, l.items)
	return NewListFrom(items)
}

// Concat creates and returns a new list that contains all the items from the
// current list followed by all the items from the other lists. The original
// lists are not modified.
func (l *List[T]) Concat(others ...*List[T]) *List[T] {
	items := make([]T, len(l.items))
	copy(items, l.items)
	for _, other := range others {
		items = append(items, other.items...)
	}
	return NewListFrom(items)
}

// ToSlice returns a slice containing all the items in the list, in order.
// Modifying the returned slice will not affect the original list.
func (l *List[T]) ToSlice() []T {
	return l.items
}

// Iter returns an iterator that yields the index and item of each element in
// the list, in order.
func (l *List[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range l.items {
			if !yield(i, item) {
				return
			}
		}
	}
}

// ForEach calls the given function for each index-item pair in the list, in order.
func (l *List[T]) ForEach(f func(int, T)) {
	for i, item := range l.items {
		f(i, item)
	}
}
