package dsx

import (
	"iter"
)

var _ Container[any] = (*Stack[any])(nil)

// Stack is a simple stack data structure that supports push and pop operations.
type Stack[T comparable] struct {
	items []T
}

// NewStack creates a new empty stack.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

// NewStackFrom creates a new stack from the given items.
func NewStackFrom[T comparable](items []T) *Stack[T] {
	return &Stack[T]{
		items: items,
	}
}

// Push adds one or more items to the top of the stack.
func (s *Stack[T]) Push(items ...T) {
	s.items = append(s.items, items...)
}

// PushSlice adds multiple items to the top of the stack from a slice. The items
// are added in the order they appear in the slice.
func (s *Stack[T]) PushSlice(items []T) {
	s.items = append(s.items, items...)
}

// Get returns the item at the specified index in the stack. The index is resolved
// using python-like negative indexing. If the index is out of range, it panics.
func (s *Stack[T]) Get(i int) T {
	i = resolveIndex(i, len(s.items))
	if i < 0 || i >= len(s.items) {
		panic("stack index out of range")
	}
	return s.items[i]
}

// GetOr returns the item at the specified index in the stack. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// provided default value.
func (s *Stack[T]) GetOr(i int, v T) T {
	i = resolveIndex(i, len(s.items))
	if i < 0 || i >= len(s.items) {
		return v
	}
	return s.items[i]
}

// GetOk returns the item at the specified index in the stack. The index is resolved
// using python-like negative indexing. If the index is out of range, it returns the
// zero value of T and false. Otherwise, it returns the item and true.
func (s *Stack[T]) GetOk(i int) (T, bool) {
	i = resolveIndex(i, len(s.items))
	if i < 0 || i >= len(s.items) {
		var zero T
		return zero, false
	}
	return s.items[i], true
}

// First returns the item at the top of the stack without removing it. If the stack is
// empty, it panics.
func (s *Stack[T]) First() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	return s.items[len(s.items)-1]
}

// FirstOr returns the item at the top of the stack without removing it. If the stack is
// empty, it returns the provided default value.
func (s *Stack[T]) FirstOr(v T) T {
	if len(s.items) == 0 {
		return v
	}
	return s.items[len(s.items)-1]
}

// FirstOk returns the item at the top of the stack without removing it. If the stack is
// empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (s *Stack[T]) FirstOk() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Last returns the item at the bottom of the stack without removing it. If the stack is
// empty, it panics.
func (s *Stack[T]) Last() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	return s.items[0]
}

// LastOr returns the item at the bottom of the stack without removing it. If the stack is
// empty, it returns the provided default value.
func (s *Stack[T]) LastOr(v T) T {
	if len(s.items) == 0 {
		return v
	}
	return s.items[0]
}

// LastOk returns the item at the bottom of the stack without removing it. If the stack is
// empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (s *Stack[T]) LastOk() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[0], true
}

// Pop removes and returns the item at the top of the stack. If the stack is empty, it panics.
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// PopOr removes and returns the item at the top of the stack. If the stack is
// empty, it returns the provided default value.
func (s *Stack[T]) PopOr(v T) T {
	if len(s.items) == 0 {
		return v
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// PopOk removes and returns the item at the top of the stack. If the stack is
// empty, it returns the zero value of T and false. Otherwise, it returns the item and true.
func (s *Stack[T]) PopOk() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// IndexOf returns the index of the first occurrence of the specified item in
// the stack, or -1 if the item is not found. The search is performed from the
// top of the stack
func (s *Stack[T]) IndexOf(item T) int {
	for i := len(s.items) - 1; i >= 0; i-- {
		if s.items[i] == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc returns the index of the first item in the stack that satisfies the
// provided predicate function, or -1 if no such item is found. The search is
// performed from the top of the stack.
func (s *Stack[T]) IndexOfFunc(f func(T) bool) int {
	for i := len(s.items) - 1; i >= 0; i-- {
		if f(s.items[i]) {
			return i
		}
	}
	return -1
}

// Contains returns true if the specified item is present in the stack, and false
// otherwise. The search is performed from the top of the stack.
func (s *Stack[T]) Contains(item T) bool {
	return s.IndexOf(item) != -1
}

// ContainsFunc returns true if there is an item in the stack that satisfies the
// provided predicate function, and false otherwise. The search is performed from
// the top of the stack.
func (s *Stack[T]) ContainsFunc(f func(T) bool) bool {
	return s.IndexOfFunc(f) != -1
}

// Size returns the number of items currently in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack, leaving it empty.
func (s *Stack[T]) Clear() {
	s.items = []T{}
}

// Clone creates and returns a new Stack that is a copy of the current stack.
func (s *Stack[T]) Clone() *Stack[T] {
	items := make([]T, len(s.items))
	copy(items, s.items)
	return NewStackFrom(items)
}

// ToSlice returns a slice containing all the items in the stack. The order of the
// items in the returned slice is the same as their order in the internal slice,
// with the top of the stack at the end of the slice and the bottom of the stack
// at the beginning.
func (s *Stack[T]) ToSlice() []T {
	return s.items
}

// Iter returns an iterator that yields the index and item of each element in the
// stack. It starts from the top of the stack (the last item in the internal slice)
// and goes down to the bottom (the first item in the internal slice).
func (s *Stack[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(i, s.items[i]) {
				return
			}
		}
	}
}

// ForEach applies the provided function f to each item in the stack, passing the
// index and the item as arguments. The iteration starts from the top of the stack (the last item in the internal slice)
// and goes down to the bottom (the first item in the internal slice).
func (s *Stack[T]) ForEach(f func(int, T)) {
	for i := len(s.items) - 1; i >= 0; i-- {
		f(i, s.items[i])
	}
}
