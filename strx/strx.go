// Package strx provides various string manipulation functions that are not
// included in the standard library's strings package. It also reexports all
// functions from the strings package for convenience.
package strx

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Indent returns a string with the specified number of spaces
// inserted before the string.
func Indent(s string, pad int) string {
	return IndentWith(s, pad, " ")
}

// IndentWith returns a string with the specified number of characters
// inserted before the string. The character used for padding is
// specified by the second parameter.
func IndentWith(s string, pad int, with string) string {
	ident := strings.Repeat(with, pad)
	return ident + strings.ReplaceAll(s, "\n", "\n"+ident)
}

// Escape replaces newlines and tabs in a string with their escaped
// representations. It returns the modified string.
func Escape(s string) string {
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

// JoinFunc is a helper function that applies a function to each element of the
// slice and then joins the resulting strings with the specified separator. It
// takes a slice of items, a function that takes an item and returns a string,
// and a separator string. It returns a string that is the result of calling
// Join with the resulting strings.
func JoinFunc[T any](items []T, fn func(T) string, sep string) string {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = fn(item)
	}
	return strings.Join(strs, sep)
}

// HumanList joins a slice of strings with "and" or "or" depending on the value of
// the last parameter. If the slice is empty, it returns an empty string. If the
// slice has one element, it returns that element. If the slice has two elements,
// it joins them with "and" or "or". If the slice has more than two elements,
// it joins them with ", " and the last element with "and" or "or".
func HumanList(items []string, and string) string {
	if len(items) == 0 {
		return ""
	}

	if len(items) == 1 {
		return items[0]
	}

	res := strings.Join(items[:len(items)-1], ", ") + " " + and + " " + items[len(items)-1]
	return res
}

// HumanListFunc is a helper function that applies a function to each element of
// the slice and then calls HumanList with the resulting strings. It takes a
// slice of items, a function that takes an item and returns a string, and a
// string to use for the last element. It returns a string that is the result of
// calling HumanList with the resulting strings.
func HumanListFunc[T any](items []T, fn func(T) string, and string) string {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = fn(item)
	}
	return HumanList(strs, and)
}

// Pads the string with spaces to the left until the string reaches the desired
// length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadLeft(s string, n int) string {
	return PadLeftWith(s, n, " ")
}

// Pads the string with the specified character to the left until the string
// reaches the desired length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadLeftWith(s string, n int, with string) string {
	n = n - Length(s)
	if n <= 0 {
		return s
	}
	return strings.Repeat(with, n) + s
}

// Pads the string with spaces to the right until the string reaches the
// desired
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadRight(s string, n int) string {
	return PadRightWith(s, n, " ")
}

// Pads the string with the specified character to the right until the string
// reaches the desired length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadRightWith(s string, n int, with string) string {
	n = n - Length(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(with, n)
}

// Pads the string with spaces to the left and right until the string reaches
// the desired length. In case of an odd number of characters, the left side
// will have one more character than the right side.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadCenter(s string, n int) string {
	return PadCenterWith(s, n, " ")
}

// Pads the string with the specified character to the left and right until the
// string reaches the desired length. In case of an odd number of characters,
// the left side will have one more character than the right side.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadCenterWith(s string, n int, with string) string {
	n = n - Length(s)
	if n <= 0 {
		return s
	}
	left := n / 2
	right := left
	if n%2 == 1 {
		left++
	}
	return strings.Repeat(with, left) + s + strings.Repeat(with, right)
}

// FirstUp capitalizes the first letter of the string and returns the modified
// string. If the string is empty, it will be returned as is.
func FirstUp(s string) string {
	if utf8.RuneCountInString(s) == 0 {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return ToUpper(string(r)) + s[size:]
}

// FirstLow lowercases the first letter of the string and returns the modified
// string. If the string is empty, it will be returned as is.
func FirstLow(s string) string {
	if utf8.RuneCountInString(s) == 0 {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return ToLower(string(r)) + s[size:]
}

// TrimSpaces removes all spaces from the start and end of the string.
func TrimSpaces(s string) string {
	return strings.Trim(s, " ")
}

// IsBlank returns true if the string is empty or contains only whitespace
// characters.
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// Truncate returns a string that is at most maxLen characters long. If the
// input string is longer than maxLen, it will be truncated.
func Truncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen])
}

// Ellipsis guarantees that the returned string will not be longer than maxLen,
// adding "..." at the end if the string is truncated. Notice that the ellipsis
// counts towards the maxLen. Also notice that if the maxLen is less than or
// equal to 3, the function will always return "...".
func Ellipsis(s string, maxLen int) string {
	size := utf8.RuneCountInString(s)
	if size <= maxLen {
		return s
	}

	if maxLen <= 3 {
		return "..."
	}

	return string([]rune(s)[:maxLen-3]) + "..."
}

// Format is a simple wrapper around fmt.Sprintf that allows you to use it
// without importing the fmt package.
func Format(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

// Length returns the number of characters in the string. It counts Unicode
// characters correctly, so it will return the expected length for strings with
// multi-byte characters. For example, Length("hello") would return 5, and
// Length("你好") would return 2.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Length returns the number of characters in the string ignoring escape codes,
// including color codes. It counts Unicode characters correctly, so it will
// return the expected length for strings with multi-byte characters.
func VisualLength(s string) int {
	count := 0
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
		} else if inEscape && r == 'm' {
			inEscape = false
		} else if !inEscape && r != utf8.RuneError {
			count++
		}
	}
	return count
}

// Lines splits the input string into lines using the newline character as a
// separator. It returns a slice of strings, where each string is a line from the
// input string. For example, Lines("hello\nworld") would return
// []string{"hello", "world"}.
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// WrapWord wraps the input string to the specified maximum length, breaking at
// word boundaries.
//
// Example: `apple, bananas and oranges` wrapped to a max length of 10 would
// become: `apple,\nbananas\nand\noranges`
func WrapWord(s string, maxLen int) string {
	return WrapWordWith(s, maxLen, "\n")
}

// WrapWordWith wraps the input string to the specified maximum length,
// breaking at word boundaries and using the specified separator.
//
// Example: `apple, bananas and oranges` wrapped to a max length of 10 with a
// separator of " | " would become: `apple, | bananas | and | oranges`
func WrapWordWith(s string, maxLen int, sep string) string {
	if maxLen <= 0 {
		return s
	}
	var b strings.Builder
	lineLen := 0
	for i, word := range IterFields(s) {
		wordLen := Length(word)
		if lineLen > 0 && lineLen+1+wordLen > maxLen {
			b.WriteString(sep)
			lineLen = 0
		} else if i > 0 {
			b.WriteString(" ")
			lineLen++
		}
		b.WriteString(word)
		lineLen += wordLen
	}
	return b.String()
}

// WrapLetter wraps the input string to the specified maximum length, breaking
// at letter boundaries.
//
// Example: `apple, bananas and oranges` wrapped to a max length of 10 would
// become: `apple, ban\nanas and\n oranges`
func WrapLetter(s string, maxLen int) string {
	return WrapLetterWith(s, maxLen, "\n")
}

// WrapLetterWith wraps the input string to the specified maximum length, breaking
// at letter boundaries and using the specified separator.
// Example: `apple, bananas and oranges` wrapped to a max length of 10 with a
// separator of " | " would become: `apple, ba | nanas and | oranges`
func WrapLetterWith(s string, maxLen int, sep string) string {
	if maxLen <= 0 {
		return s
	}

	var b strings.Builder
	lineLen := 0
	for _, r := range s {
		if lineLen > 0 && lineLen+1 > maxLen {
			b.WriteString("\n")
			lineLen = 0
		}
		b.WriteRune(r)
		lineLen++
	}
	return b.String()
}

// WrapHyphen wraps the input string to the specified maximum length, breaking
// at letter boundaries and adding a hyphen at the end of the line if the word
// is broken.
//
// Example: `apple, bananas and oranges` wrapped to a max length of 10 would
// become: `apple, bana-\nanas and \noranges`
func WrapHyphen(s string, maxLen int) string {
	return WrapHyphenWith(s, maxLen, "\n")
}

// WrapHyphenWith wraps the input string to the specified maximum length, breaking
// at letter boundaries and adding a hyphen at the end of the line if the word
// is broken. It uses the specified separator between lines.
// Example: `apple, bananas and oranges` wrapped to a max length of 10 with a
// separator of " | " would become: `apple, bana- | nas and | oranges`
func WrapHyphenWith(s string, maxLen int, sep string) string {
	if maxLen <= 1 {
		return s
	}

	var b strings.Builder
	lineLen := 0
	for _, word := range IterFields(s) {
	reset:
		wordLen := Length(word)

		if lineLen+wordLen <= maxLen {
			b.WriteString(word)
			lineLen += wordLen
			if lineLen < maxLen {
				b.WriteString(" ")
				lineLen++
			}
			continue
		}

		if maxLen-lineLen <= 1 {
			b.WriteString(sep)
			lineLen = 0
			goto reset
		}

		runes := []rune(word)
		i := maxLen - lineLen - 1
		b.WriteString(string(runes[:i]))
		b.WriteString("-\n")
		word = string(runes[i:])
		lineLen = 0
		goto reset
	}
	return b.String()
}

// Reverse returns a new string which is the reverse of the input string. It
// handles Unicode characters correctly, so it will reverse the string as expected.
// For example, Reverse("hello") would return "olleh", and Reverse("你好") would
// return "好你".
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToPrintableAscii returns a new string containing only the printable ASCII
// characters from the input string. It removes any non-printable or non-ASCII
// characters. For example, ToPrintableAscii("hello\nworld") would return
// "helloworld", and ToPrintableAscii("你好") would return an empty string.
func ToPrintableAscii(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 0x20 && c <= 0x7E {
			b = append(b, c)
		}
	}
	return string(b)
}
