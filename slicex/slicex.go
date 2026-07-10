package slicex

import (
	"cmp"
	"iter"
	"slices"
)

func resolveIndex(i, len int) int {
	if i < 0 {
		i += len
	}
	if i > len {
		i = len
	}
	return i
}

// New returns a new slice with the given elements.
func New[T any](elements ...T) []T {
	return elements
}

// Append appends the given values to the end of the slice and returns the resulting slice.
func Append[T any](s []T, v ...T) []T {
	return append(s, v...)
}

// AppendSlice appends the given slice to the end of the slice and returns the resulting slice.
func AppendSlice[T any](s []T, v []T) []T {
	return append(s, v...)
}

// Prepend prepends the given values to the beginning of the slice and returns the resulting slice.
func Prepend[T any](s []T, v ...T) []T {
	return append(v, s...)
}

// PrependSlice prepends the given slice to the beginning of the slice and returns the resulting slice.
func PrependSlice[T any](s []T, v []T) []T {
	return append(v, s...)
}

// Insert inserts the given values at the specified index in the slice and returns the resulting slice.
// It accepts python-like negative indeces.
func Insert[T any](s []T, i int, v ...T) []T {
	i = resolveIndex(i, len(s))
	return slices.Insert(s, i, v...)
}

// InsertSlice inserts the given slice at the specified index in the slice and returns the resulting slice.
// It accepts python-like negative indeces.
func InsertSlice[T any](s []T, i int, v []T) []T {
	i = resolveIndex(i, len(s))
	return slices.Insert(s, i, v...)
}

// Remove removes the element at the specified index from the slice and returns the resulting slice.
// It accepts python-like negative indeces.
func Remove[T any](s []T, i int) []T {
	i = resolveIndex(i, len(s))
	return append(s[:i], s[i+1:]...)
}

// RemoveRange removes the elements in the specified index range from the slice and returns the resulting slice.
// It accepts python-like negative indeces.
func RemoveRange[T any](s []T, i, j int) []T {
	i = resolveIndex(i, len(s))
	j = resolveIndex(j, len(s))
	if i > j {
		i, j = j, i
	}
	return append(s[:i], s[j:]...)
}

// RemoveFunc removes the elements that satisfy the given predicate function from the slice and returns the resulting slice.
func RemoveFunc[T any](s []T, f func(T) bool) []T {
	var j int
	for i, v := range s {
		if !f(v) {
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}

// RemoveValue removes all occurrence of the given value from the slice and returns the resulting slice.
func RemoveValue[T comparable](s []T, v T) []T {
	result := s[:0]
	for _, x := range s {
		if x != v {
			result = append(result, x)
		}
	}
	return result
}

// Assign appends the given slices to the destination slice. It is useful
// for variadic functions that accept multiple slices as arguments.
func Assign[T any](dst []T, src ...[]T) []T {
	for _, s := range src {
		dst = append(dst, s...)
	}
	return dst
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[T any](s []T) []T { return slices.Clip(s) }

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
// The result may have additional unused capacity.
func Clone[T any](s []T) []T { return slices.Clone(s) }

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func Compact[T comparable](s []T) []T { return slices.Compact(s) }

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func CompactFunc[T any](s []T, eq func(T, T) bool) []T { return slices.CompactFunc(s, eq) }

// Concat returns a new slice concatenating the passed in slices.
func Concat[T any](val ...[]T) []T { return slices.Concat(val...) }

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Empty and nil slices are considered equal.
// Floating point NaNs are not considered equal.
func Equal[T comparable](s1, s2 []T) bool { return slices.Equal(s1, s2) }

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[T any](s1, s2 []T, eq func(T, T) bool) bool { return slices.EqualFunc(s1, s2, eq) }

// Compare compares the elements of s1 and s2, using [cmp.Compare] on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare[T cmp.Ordered](s1, s2 []T) int { return slices.Compare(s1, s2) }

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[T any](s1, s2 []T, cmp func(T, T) int) int { return slices.CompareFunc(s1, s2, cmp) }

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool { return slices.Contains(s, v) }

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[T any](s []T, f func(T) bool) bool { return slices.ContainsFunc(s, f) }

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[T any](s []T, n int) []T { return slices.Grow(s, n) }

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func IndexOf[T comparable](s []T, v T) int { return slices.Index(s, v) }

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexOfFunc[T any](s []T, f func(T) bool) int { return slices.IndexFunc(s, f) }

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[T cmp.Ordered](s []T) bool { return slices.IsSorted(s) }

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func IsSortedFunc[T any](s []T, cmp func(a, b T) int) bool { return slices.IsSortedFunc(s, cmp) }

// Max returns the maximal value in x. It panics if x is empty.
// For floating-point E, Max propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Max[T cmp.Ordered](s []T) T { return slices.Max(s) }

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func MaxFunc[T any](s []T, cmp func(a, b T) int) T { return slices.MaxFunc(s, cmp) }

// Min returns the minimal value in x. It panics if x is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Min[T cmp.Ordered](s []T) T { return slices.Min(s) }

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func MinFunc[T any](s []T, cmp func(a, b T) int) T { return slices.MinFunc(s, cmp) }

// Repeat returns a new slice that repeats the provided slice the given number of times.
// The result has length and capacity (len(x) * count).
// The result is never nil.
// Repeat panics if count is negative or if the result of (len(x) * count)
// overflows.
func Repeat[T any](s []T, count int) []T { return slices.Repeat(s, count) }

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func Replace[T any](s []T, i, j int, v ...T) []T { return slices.Replace(s, i, j, v...) }

// Reverse reverses the elements of the slice in place.
func Reverse[T any](s []T) { slices.Reverse(s) }

// Reversed returns a new slice with the elements of the original slice in reverse order.
func Reversed[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	slices.Reverse(result)
	return result
}

// Sort sorts a slice of any ordered type in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func Sort[T cmp.Ordered](s []T) { slices.Sort(s) }

// SortFunc sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b or a and b are incomparable in the sense of
// a strict weak ordering.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
// The function should return 0 for incomparable items.
func SortFunc[T any](s []T, cmp func(a, b T) int) { slices.SortFunc(s, cmp) }

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func SortStableFunc[T any](s []T, cmp func(a, b T) int) { slices.SortStableFunc(s, cmp) }

// Sorted collects values from seq into a new slice, sorts the slice,
// and returns it.
func Sorted[T cmp.Ordered](seq iter.Seq[T]) []T { return slices.Sorted(seq) }

// SortedFunc collects values from seq into a new slice, sorts the slice
// using the comparison function, and returns it.
func SortedFunc[T any](seq iter.Seq[T], cmp func(a, b T) int) []T { return slices.SortedFunc(seq, cmp) }

// SortedStableFunc collects values from seq into a new slice.
// It then sorts the slice while keeping the original order of equal elements,
// using the comparison function to compare elements.
// It returns the new slice.
func SortedStableFunc[T any](seq iter.Seq[T], cmp func(a, b T) int) []T {
	return slices.SortedStableFunc(seq, cmp)
}

// Iter returns an iterator over index-value pairs in the slice
// in the usual order.
func Iter[T any](s []T) iter.Seq2[int, T] {
	return slices.All(s)
}

// IterBackward returns an iterator over index-value pairs in the slice,
// traversing it backward with descending indices.
func IterBackward[T any](s []T) iter.Seq2[int, T] {
	return slices.Backward(s)
}

// IterValues returns an iterator that yields the slice elements in order.
func IterValues[T any](s []T) iter.Seq[T] {
	return slices.Values(s)
}

// Map applies the function f to each element of the slice,
// and returns a new slice with the results.
func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter applies the function f to each element of the slice,
// and returns a new slice with the elements for which f returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies the function f to each element of the slice,
// and returns the accumulated result.
func Reduce[T any, U any](s []T, f func(U, T) U, initial U) U {
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

// ForEach applies the function f to each element of the slice.
func ForEach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

// ForEachIndex applies the function f to each index-value pair of the slice.
func ForEachIndex[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}
