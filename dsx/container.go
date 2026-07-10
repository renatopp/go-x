package dsx

import (
	"iter"
)

type Container[T comparable] interface {
	Get(i int) T
	GetOr(i int, v T) T
	GetOk(i int) (T, bool)
	First() T
	FirstOr(v T) T
	FirstOk() (T, bool)
	Last() T
	LastOr(v T) T
	LastOk() (T, bool)
	Pop() T
	PopOr(v T) T
	PopOk() (T, bool)
	IndexOf(item T) int
	IndexOfFunc(f func(T) bool) int
	Contains(item T) bool
	ContainsFunc(f func(T) bool) bool
	Size() int
	Clear()
	ToSlice() []T
	Iter() iter.Seq2[int, T]
}
