package iterx

import (
	"iter"

	"github.com/renatopp/go-x/randx"
)

// Window returns a Seq that yields consecutive windows of the specified size from the input sequence.
// For example, Window([]int{1, 2, 3, 4}, 2) would yield []int{1, 2}, then []int{2, 3}, and finally []int{3, 4}.
// If the size is greater than the length of the input sequence, no windows will be yielded.
func Window[V any](seq []V, size int) iter.Seq[[]V] {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	return func(yield func([]V) bool) {
		var window []V
		for _, v := range seq {
			window = append(window, v)
			if len(window) > size {
				window = window[1:]
			}
			if len(window) == size {
				if !yield(window) {
					return
				}
			}
		}
	}
}

// WindowString returns a Seq that yields consecutive windows of the specified size from the input string.
// For example, WindowString("abcd", 2) would yield "ab", then "bc", and finally "cd".
// If the size is greater than the length of the input string, no windows will be yielded.
func WindowString(seq string, size int) iter.Seq[string] {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	return func(yield func(string) bool) {
		var window []rune
		for _, r := range seq {
			window = append(window, r)
			if len(window) > size {
				window = window[1:]
			}
			if len(window) == size {
				if !yield(string(window)) {
					return
				}
			}
		}
	}
}

// Shuffle returns a Seq that yields the elements of the input sequence in random order. Each element is yielded exactly once.
func Shuffle[V any](seq []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		perm := make([]int, len(seq))
		for i := range perm {
			perm[i] = i
		}
		randx.Shuffle(perm)
		for _, i := range perm {
			if !yield(seq[i]) {
				return
			}
		}
	}
}

// RangeInt returns a Seq that yields integers from start (inclusive) to end (exclusive).
// For example, RangeInt(0, 5) would yield 0, 1, 2, 3, and 4.
func RangeInt(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// RangeIntStep returns a Seq that yields integers from start (inclusive) to end (exclusive) with a specified step.
// For example, RangeIntStep(0, 10, 2) would yield 0, 2, 4, 6, and 8. If step is negative, it will yield in reverse order (e.g., RangeIntStep(10, 0, -2) would yield 10, 8, 6, 4, and 2).
func RangeIntStep(start, end, step int) iter.Seq[int] {
	if step == 0 {
		panic("step must be non-zero")
	}
	return func(yield func(int) bool) {
		if step > 0 {
			for i := start; i < end; i += step {
				if !yield(i) {
					return
				}
			}
		} else {
			for i := start; i > end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// RangeFloat returns a Seq that yields float64 values from start (inclusive) to end (exclusive).
// For example, RangeFloat(0.0, 1.0) would yield values starting from 0.0 up to but not
// including 1.0, with a default step of 1.0 (i.e., it would yield 0.0 only). To specify a different step, use RangeFloatStep.
func RangeFloat(start, end float64) iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// RangeFloatStep returns a Seq that yields float64 values from start (inclusive) to end (exclusive) with a specified step.
// For example, RangeFloatStep(0.0, 1.0, 0.2) would yield 0.0, 0.2, 0.4, 0.6, and 0.8.
// If step is negative, it will yield in reverse order (e.g., RangeFloatStep(1.0, 0.0, -0.2)
// would yield 1.0, 0.8, 0.6, 0.4, and 0.2).
func RangeFloatStep(start, end, step float64) iter.Seq[float64] {
	if step == 0 {
		panic("step must be non-zero")
	}
	return func(yield func(float64) bool) {
		if step > 0 {
			for i := start; i < end; i += step {
				if !yield(i) {
					return
				}
			}
		} else {
			for i := start; i > end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// Repeat returns a Seq that yields the specified value indefinitely.
func Repeat[T any](value T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}

// RepeatN returns a Seq that yields the specified value n times. If n is less
// than or equal to 0, it will yield nothing.
func RepeatN[T any](value T, n int) iter.Seq[T] {
	if n <= 0 {
		return func(yield func(T) bool) {}
	}
	return func(yield func(T) bool) {
		for i := 0; i < n; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// RepeatWhile returns a Seq that yields the specified value repeatedly as long as
// the keep function returns true. The keep function is called before each yield,
// and if it returns false, the sequence will stop yielding values.
func RepeatWhile[T any](value T, keep func() bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for keep() {
			if !yield(value) {
				return
			}
		}
	}
}

// RepeatUntil returns a Seq that yields the specified value repeatedly until
// the stop function returns true. The stop function is called before each
// yield, and if it returns true, the sequence will stop yielding values.
func RepeatUntil[T any](value T, stop func() bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if stop() {
				return
			}
			if !yield(value) {
				return
			}
		}
	}
}

// Cycle returns a Seq that yields the elements of the input sequence in order,
// repeating indefinitely. For example, Cycle([]int{1, 2, 3}) would yield 1,
// then 2, then 3, then 1 again, and so on.
func Cycle[T any](seq []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for { // infinite loop
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// CycleN returns a Seq that yields the elements of the input sequence in order,
// repeating n times. For example, CycleN([]int{1, 2, 3}, 2) would yield 1, then 2,
// then 3, then 1 again, then 2 again, and finally 3 again. If n is less than or
// equal to 0, it will yield nothing.
func CycleN[T any](seq []T, n int) iter.Seq[T] {
	if n <= 0 {
		return func(yield func(T) bool) {}
	}
	return func(yield func(T) bool) {
		for i := 0; i < n; i++ {
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// CycleWhile returns a Seq that yields the elements of the input sequence in order,
// repeating as long as the keep function returns true. The keep function is called
// before each cycle, and if it returns false, the sequence will stop yielding values.
func CycleWhile[T any](seq []T, keep func() bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for keep() {
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// CycleUntil returns a Seq that yields the elements of the input sequence in order,
// repeating until the stop function returns true. The stop function is called before
// each cycle, and if it returns true, the sequence will stop yielding values.
func CycleUntil[T any](seq []T, stop func() bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if stop() {
				return
			}
			for _, v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Take returns a Seq that yields the first n elements of the input sequence.
// If n is less than or equal to 0, it will yield nothing. If n is greater than
// the length of the input sequence, it will yield all elements of the input sequence.
func Take[T any](seq []T, n int) iter.Seq[T] {
	if n <= 0 {
		return func(yield func(T) bool) {}
	}
	return func(yield func(T) bool) {
		for i, v := range seq {
			if i >= n {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// TakeWhile returns a Seq that yields elements from the input sequence as long
// as the keep function returns true.
func TakeWhile[T any](seq []T, keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range seq {
			if !keep(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// TakeUntil returns a Seq that yields elements from the input sequence until
// the stop function returns true.
func TakeUntil[T any](seq []T, stop func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range seq {
			if stop(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Zip returns a Seq2 that yields pairs of elements from the input slices a
// and b. The first element of each pair is taken from a, and the second element
// is taken from b. The sequence will yield pairs until one of the input slices
// is exhausted (i.e., it will yield min(len(a), len(b)) pairs).
func Zip[T, Q any](a []T, b []Q) iter.Seq2[T, Q] {
	return func(yield func(T, Q) bool) {
		minLen := len(a)
		if len(b) < minLen {
			minLen = len(b)
		}
		for i := 0; i < minLen; i++ {
			if !yield(a[i], b[i]) {
				return
			}
		}
	}
}

// Enumerate returns a Seq2 that yields pairs of index and element from the input slice seq.
// The first element of each pair is the index of the element in seq, and the second element is the value of the element.
// For example, Enumerate([]string{"a", "b", "c"}) would yield (0, "a"), then (1, "b"), and finally (2, "c").
func Enumerate[T any](seq []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range seq {
			if !yield(i, v) {
				return
			}
		}
	}
}
