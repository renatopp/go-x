package dsx

import (
	"container/heap"
	"iter"
)

var _ Container[any] = (*Heap[any])(nil)

// HeapItem represents an item in the heap with its associated value for
// comparison.
type HeapItem[T comparable] struct {
	Item  T
	Value int
}

// Heap is a min-heap data structure that stores items of type T with
// associated integer values for comparison. The heap maintains the property
// that the item with the largest value is at the top of the heap (i.e., the
// last element in the internal slice).
type Heap[T comparable] struct {
	items internalHeap[T]
}

// NewHeap creates and returns a new empty Heap for items of type T.
func NewHeap[T comparable]() *Heap[T] {
	return &Heap[T]{items: internalHeap[T]{}}
}

// Push adds one or more items to the heap with the same specified value. Each
// item will be added individually, maintaining the heap property. The value
// parameter is used to determine the position of the item in the heap - larger
// values will be pushed to the end of the slice.
func (h *Heap[T]) Push(value int, item ...T) {
	for _, item := range item {
		heap.Push(&h.items, HeapItem[T]{Item: item, Value: value})
	}
}

// PushSlice adds multiple items to the heap from a slice with the same specified value.
// Each item will be added individually, maintaining the heap property.
func (h *Heap[T]) PushSlice(value int, items []T) {
	for _, item := range items {
		h.Push(value, item)
	}
}

// Get retrieves the item at the specified index in the heap. The index is
// resolved using python-like negative indexing. If the index is out of range,
// it panics.
func (h *Heap[T]) Get(i int) T {
	i = resolveIndex(i, len(h.items))
	if i < 0 || i >= len(h.items) {
		panic("heap index out of range")
	}
	return h.items[i].Item
}

// GetOr retrieves the item at the specified index in the heap. The index is
// resolved using python-like negative indexing. If the index is out of range,
// it returns the provided default value v.
func (h *Heap[T]) GetOr(i int, v T) T {
	i = resolveIndex(i, len(h.items))
	if i < 0 || i >= len(h.items) {
		return v
	}
	return h.items[i].Item
}

// GetOk retrieves the item at the specified index in the heap. The index is
// resolved using python-like negative indexing. If the index is out of range,
// it returns the zero value of T and false. Otherwise, it returns the item and true.
func (h *Heap[T]) GetOk(i int) (T, bool) {
	i = resolveIndex(i, len(h.items))
	if i < 0 || i >= len(h.items) {
		var zero T
		return zero, false
	}
	return h.items[i].Item, true
}

// First retrieves the item with the smallest value in the heap (i.e., the first
// item in the internal slice). If the heap is empty, it panics.
func (h *Heap[T]) First() T {
	if len(h.items) == 0 {
		panic("heap is empty")
	}
	return h.items[0].Item
}

// FirstOr retrieves the item with the smallest value in the heap (i.e., the first
// item in the internal slice). If the heap is empty, it returns the provided
// default value v.
func (h *Heap[T]) FirstOr(v T) T {
	if len(h.items) == 0 {
		return v
	}
	return h.items[0].Item
}

// FirstOk retrieves the item with the smallest value in the heap (i.e., the first
// item in the internal slice). If the heap is empty, it returns the zero value
// of T and false. Otherwise, it returns the item and true.
func (h *Heap[T]) FirstOk() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	return h.items[0].Item, true
}

// Last retrieves the item with the largest value in the heap (i.e., the last
// item in the internal slice). If the heap is empty, it panics. This is an
// alias for Top.
func (h *Heap[T]) Last() T {
	if len(h.items) == 0 {
		panic("heap is empty")
	}
	return h.items[len(h.items)-1].Item
}

// LastOr retrieves the item with the largest value in the heap (i.e., the last
// item in the internal slice). If the heap is empty, it returns the provided
// default value v. This is an alias for TopOr.
func (h *Heap[T]) LastOr(v T) T {
	if len(h.items) == 0 {
		return v
	}
	return h.items[len(h.items)-1].Item
}

// LastOk retrieves the item with the largest value in the heap (i.e., the last
// item in the internal slice). If the heap is empty, it returns the zero value
// of T and false. Otherwise, it returns the item and true. This is an alias for TopOk.
func (h *Heap[T]) LastOk() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	return h.items[len(h.items)-1].Item, true
}

// Pop removes and returns the item with the largest value in the heap
// (i.e., the last item in the internal slice). If the heap is empty, it panics.
func (h *Heap[T]) Pop() T {
	if len(h.items) == 0 {
		panic("heap is empty")
	}
	item := heap.Pop(&h.items).(HeapItem[T])
	return item.Item
}

// PopOr removes and returns the item with the largest value in the heap
// (i.e., the last item in the internal slice). If the heap is empty, it returns
// the provided default value v.
func (h *Heap[T]) PopOr(v T) T {
	if len(h.items) == 0 {
		return v
	}
	item := heap.Pop(&h.items).(HeapItem[T])
	return item.Item
}

// PopOk removes and returns the item with the largest value in the heap
// (i.e., the last item in the internal slice). If the heap is empty, it returns
// the zero value of T and false. Otherwise, it returns the item and true.
func (h *Heap[T]) PopOk() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(&h.items).(HeapItem[T])
	return item.Item, true
}

// IndexOf returns the index of the first occurrence of the specified item in
// the heap. If the item is not found, it returns -1.
func (h *Heap[T]) IndexOf(item T) int {
	for i, v := range h.items {
		if v.Item == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc returns the index of the first item in the heap that satisfies
// the provided function f. The function f takes an item of type T and returns
// a boolean indicating whether the item satisfies the condition. If no item
// satisfies the condition, it returns -1.
func (h *Heap[T]) IndexOfFunc(f func(T) bool) int {
	for i, v := range h.items {
		if f(v.Item) {
			return i
		}
	}
	return -1
}

// Contains checks if the specified item exists in the heap. It returns true if
// the item is found, and false otherwise.
func (h *Heap[T]) Contains(item T) bool {
	for _, v := range h.items {
		if v.Item == item {
			return true
		}
	}
	return false
}

// ContainsFunc checks if there exists an item in the heap that satisfies the
// provided function f. The function f takes an item of type T and returns a
// boolean indicating whether the item satisfies the condition. It returns true
// if such an item is found, and false otherwise.
func (h *Heap[T]) ContainsFunc(f func(T) bool) bool {
	for _, v := range h.items {
		if f(v.Item) {
			return true
		}
	}
	return false
}

// Size returns the number of items currently in the heap.
func (h *Heap[T]) Size() int {
	return len(h.items)
}

// Clear removes all items from the heap, resulting in an empty heap.
func (h *Heap[T]) Clear() {
	h.items = internalHeap[T]{}
}

// Clone creates and returns a new Heap that is a copy of the current heap. The
// items in the new heap are the same as those in the original heap, but they
// are stored in a new internal slice to ensure that modifications to the new
// heap do not affect the original heap.
func (h *Heap[T]) Clone() *Heap[T] {
	items := make(internalHeap[T], len(h.items))
	copy(items, h.items)
	return &Heap[T]{items: items}
}

// Concat creates and returns a new Heap that contains all the items from the
// current heap followed by all the items from the other heap. The items from
// the other heap are added to the new heap using the Push method, which ensures
// that the heap property is maintained. The original heaps are not modified.
func (h *Heap[T]) Concat(other *Heap[T]) *Heap[T] {
	heap := h.Clone()
	for _, item := range other.items {
		heap.Push(item.Value, item.Item)
	}
	return heap
}

// ToSlice returns a slice containing all the items in the heap. The order of the
// items in the returned slice is the same as their order in the internal slice,
// perserving the value-based ordering, starting with the item with the smallest value.
func (h *Heap[T]) ToSlice() []T {
	items := make([]T, len(h.items))
	for i, item := range h.items {
		items[i] = item.Item
	}
	return items
}

// Iter returns an iterator that yields the index and item of each element in the
// heap. It starts from the first item (the one with the smallest value).
func (h *Heap[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range h.items {
			if !yield(i, item.Item) {
				return
			}
		}
	}
}

// ForEach applies the provided function f to each item in the heap, passing the
// index and the item as arguments. The iteration starts from the first item (the
// one with the smallest value).
func (h *Heap[T]) ForEach(f func(int, T)) {
	for i, item := range h.items {
		f(i, item.Item)
	}
}

// INTERNAL -------------------------------------------------------------------

type internalHeap[T comparable] []HeapItem[T]

func (h internalHeap[T]) Len() int           { return len(h) }
func (h internalHeap[T]) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h internalHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *internalHeap[T]) Push(x any) {
	*h = append(*h, x.(HeapItem[T]))
}
func (h *internalHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
