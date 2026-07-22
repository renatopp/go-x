package mapx

// This file reexports all functions from the standard library's maps package
// for convenience.

import (
	"iter"
	"maps"
)

// Clone returns a new map that is a copy of the given map. This is a shallow
// copy, so if the values are reference types, they will be shared between the
// original and the clone.
func Clone[K comparable, V any](m map[K]V) map[K]V { return maps.Clone(m) }

// RemoveFunc removes the key-value pairs that satisfy the given predicate
// function.
func RemoveFunc[K comparable, V any](m map[K]V, fn func(K, V) bool) {
	maps.DeleteFunc(m, fn)
}

// Equal returns true if the two maps are equal, i.e. they have the same keys
// and the same values for those keys.
func Equal[K comparable, V comparable](m1, m2 map[K]V) bool { return maps.Equal(m1, m2) }

// EqualFunc returns true if the two maps are equal according to the given
// comparison function for values.
func EqualFunc[K comparable, V any](m1, m2 map[K]V, fn func(V, V) bool) bool {
	return maps.EqualFunc(m1, m2, fn)
}

// Iter returns a sequence of key-value pairs in the map.
func Iter[K comparable, V any](m map[K]V) iter.Seq2[K, V] { return maps.All(m) }

// IterKeys returns a sequence of keys in the map.
func IterKeys[K comparable, V any](m map[K]V) iter.Seq[K] { return maps.Keys(m) }

// IterValues returns a sequence of values in the map.
func IterValues[K comparable, V any](m map[K]V) iter.Seq[V] { return maps.Values(m) }
