package strx

// This file rexports all functions from the standard library's strings package
// for convenience.

import (
	"strings"
)

// Clone returns a fresh copy of s.
// It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring
// of a much larger string. Using Clone can help such programs
// use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory.
// Clone should typically be used only rarely, and only when
// profiling indicates that it is needed.
// For strings of length zero the string "" will be returned
// and no allocation is made.
func Clone(s string) string { return strings.Clone(s) }

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Use Compare when you need to perform a three-way comparison (with
// [slices.SortFunc], for example). It is usually clearer and always faster
// to use the built-in string comparison operators ==, <, >, and so on.
func Compare(a, b string) int { return strings.Compare(a, b) }

// Contains reports whether substr is within s.
func Contains(s, substr string) bool { return strings.Contains(s, substr) }

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool { return strings.ContainsAny(s, chars) }

// ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
func ContainsFunc(s string, f func(rune) bool) bool { return strings.ContainsFunc(s, f) }

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool { return strings.ContainsRune(s, r) }

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr string) int { return strings.Count(s, substr) }

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut(s, sep string) (before, after string, found bool) { return strings.Cut(s, sep) }

// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
func CutPrefix(s, prefix string) (after string, found bool) { return strings.CutPrefix(s, prefix) }

// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
func CutSuffix(s, suffix string) (before string, found bool) { return strings.CutSuffix(s, suffix) }

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func EqualFold(s, t string) bool { return strings.EqualFold(s, t) }

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by [unicode.IsSpace], returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields(s string) []string { return strings.Fields(s) }

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
func FieldsFunc(s string, f func(rune) bool) []string { return strings.FieldsFunc(s, f) }

// HasPrefix reports whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// StartsWith reports whether the string s begins with prefix. Alias to HasPrefix.
func StartsWith(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// HasSuffix reports whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// EndsWith reports whether the string s ends with suffix. Alias to HasSuffix.
func EndsWith(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// IndexOf returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func IndexOf(s, substr string) int { return strings.Index(s, substr) }

// IndexOfAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexOfAny(s, chars string) int { return strings.IndexAny(s, chars) }

// IndexOfByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexOfByte(s string, c byte) int { return strings.IndexByte(s, c) }

// IndexOfFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexOfFunc(s string, f func(rune) bool) int { return strings.IndexFunc(s, f) }

// IndexOfRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is [utf8.RuneError], it returns the first instance of any
// invalid UTF-8 byte sequence.
func IndexOfRune(s string, r rune) int { return strings.IndexRune(s, r) }

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Join(elems []string, sep string) string { return strings.Join(elems, sep) }

// LastIndexOf returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndexOf(s, substr string) int { return strings.LastIndex(s, substr) }

// LastIndexOfAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexOfAny(s, chars string) int { return strings.LastIndexAny(s, chars) }

// LastIndexOfByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexOfByte(s string, c byte) int { return strings.LastIndexByte(s, c) }

// LastIndexOfFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexOfFunc(s string, f func(rune) bool) int { return strings.LastIndexFunc(s, f) }

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string { return strings.Map(mapping, s) }

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func Repeat(s string, count int) string { return strings.Repeat(s, count) }

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string { return strings.Replace(s, old, new, n) }

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new string) string { return strings.ReplaceAll(s, old, new) }

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to [SplitN] with a count of -1.
//
// To split around the first instance of a separator, see [Cut].
func Split(s, sep string) []string { return strings.Split(s, sep) }

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to [SplitAfterN] with a count of -1.
func SplitAfter(s, sep string) []string { return strings.SplitAfter(s, sep) }

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring will be the unsplit remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [SplitAfter].
func SplitAfterN(s, sep string, n int) []string { return strings.SplitAfterN(s, sep, n) }

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring will be the unsplit remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [Split].
//
// To split around the first instance of a separator, see [Cut].
func SplitN(s, sep string, n int) []string { return strings.SplitN(s, sep, n) }

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func ToValidUTF8(s, replacement string) string { return strings.ToValidUTF8(s, replacement) }

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s, cutset string) string { return strings.Trim(s, cutset) }

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s string, f func(rune) bool) string { return strings.TrimFunc(s, f) }

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
func TrimLeft(s, cutset string) string { return strings.TrimLeft(s, cutset) }

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s string, f func(rune) bool) string { return strings.TrimLeftFunc(s, f) }

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string { return strings.TrimPrefix(s, prefix) }

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
func TrimRight(s, cutset string) string { return strings.TrimRight(s, cutset) }

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s string, f func(rune) bool) string { return strings.TrimRightFunc(s, f) }

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string { return strings.TrimSpace(s) }

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string { return strings.TrimSuffix(s, suffix) }
