package slicex

import (
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

// Reversed returns a new slice with the elements of the original slice in reverse order.
func Reversed[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	slices.Reverse(result)
	return result
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
