package dsx_test

import (
	"testing"
)

// equalSlice compares two slices element by element. testx.Equal cannot be used
// directly on slices because it compares values with `any(a) != any(b)`, which
// panics at runtime for uncomparable types like slices.
func equalSlice[T comparable](t *testing.T, expected, actual []T) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatalf("Expected slice %v (len %d), but got %v (len %d).", expected, len(expected), actual, len(actual))
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("Expected %v at index %d, but got %v. Expected: %v, Actual: %v", expected[i], i, actual[i], expected, actual)
		}
	}
}

// sameElements compares two slices ignoring order, for results coming from
// map iteration (e.g. Dictionary.Keys/Values) where order is not guaranteed.
func sameElements[T comparable](t *testing.T, expected, actual []T) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatalf("Expected %v (len %d), but got %v (len %d).", expected, len(expected), actual, len(actual))
	}
	for _, e := range expected {
		found := false
		for _, a := range actual {
			if e == a {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Expected %v to contain %v, but it did not.", actual, e)
		}
	}
}
