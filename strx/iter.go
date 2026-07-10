package strx

import (
	"iter"
	"strings"
)

// IterString returns a Seq2 that yields each character of the input string as a
// separate string, along with its index. For example, IterString("abc")
// would yield (0, "a"), then (1, "b"), and finally (2, "c").
func IterString(seq string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, r := range seq {
			if !yield(i, string(r)) {
				return
			}
		}
	}
}

// IterRunes returns a Seq2 that yields each character of the input string as a
// separate rune, along with its index. For example, IterRunes("abc") would
// yield (0, 'a'), then (1, 'b'), and finally (2, 'c').
func IterRunes(seq string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range seq {
			if !yield(i, r) {
				return
			}
		}
	}
}

// IterFields returns a Seq2 that yields each word in the input string, along
// with its index. Words are defined as sequences of non-space characters
// separated by spaces. For example, IterFields("hello world") would yield
// (0, "hello") and then (1, "world").
func IterFields(s string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for word := range strings.FieldsSeq(s) {
			if !yield(i, word) {
				return
			}
			i++
		}
	}
}

// IterFieldsFunc returns a Seq2 that yields each word in the input string, along
// with its index. Words are defined as sequences of characters separated by
// characters for which the provided function returns true. For example,
// IterFieldsFunc("hello,world", func(r rune) bool { return r == ',' }) would
// yield (0, "hello") and then (1, "world").
func IterFieldsFunc(s string, f func(rune) bool) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for word := range strings.FieldsFuncSeq(s, f) {
			if !yield(i, word) {
				return
			}
			i++
		}
	}
}

// IterLines returns a Seq2 that yields each line of the input string, along with
// its index. Lines are defined as sequences of characters separated by newline
// characters. For example, IterLines("line1\nline2") would yield (0, "line1")
// and then (1, "line2").
func IterLines(s string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for line := range strings.Lines(s) {
			if !yield(i, Trim(line, "\r\n")) {
				return
			}
			i++
		}
	}
}

// IterSplit returns a Seq2 that yields each substring of the input string that is
// separated by the specified separator, along with its index. For example,
// IterSplit("a,b,c", ",") would yield (0, "a"), then (1, "b"), and finally (2, "c").
func IterSplit(s, sep string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for part := range strings.SplitSeq(s, sep) {
			if !yield(i, part) {
				return
			}
			i++
		}
	}
}

// IterSplitAfter returns a Seq2 that yields each substring of the input string that is
// separated by the specified separator, along with its index. The separator is
// included at the end of each substring. For example, IterSplitAfter("a,b,c", ",")
// would yield (0, "a,"), then (1, "b,"), and finally (2, "c").
func IterSplitAfter(s, sep string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for part := range strings.SplitAfterSeq(s, sep) {
			if !yield(i, part) {
				return
			}
			i++
		}
	}
}
