package runex

import (
	"slices"
)

// Alias for IsDigit
var IsNumber = IsDigit

// Alias for IsDigit
var IsNumeric = IsDigit

// Alias for IsLetter
var IsAlpha = IsLetter

// Check if a rune is a digit (0-9)
func IsDigit(r rune) bool { return r >= '0' && r <= '9' }

// Check if a rune is a letter (a-z or A-Z)
func IsLetter(r rune) bool { return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' }

// Check if a rune is a newline
func IsNewline(r rune) bool { return r == '\n' }

// Check if a rune is a whitespace (space, tab, newline or carriage return)
func IsWhitespace(r rune) bool { return r == ' ' || r == '\n' || r == '\r' }

// Check if a rune is a space (space or carriage return)
func IsSpace(r rune) bool { return r == ' ' || r == '\r' }

// Check if a rune is the end of file (rune 0)
func IsEof(r rune) bool { return r == 0 }

// Check if a rune is a letter or a digit
func IsAlphaNumeric(r rune) bool { return IsLetter(r) || IsDigit(r) }

// Check if a rune is any of the runes in the given slice
func IsOneOf(r rune, runes ...rune) bool { return slices.Contains(runes, r) }

// Check if a rune is a hexadecimal digit (0-9, a-f or A-F)
func IsHexadecimal(r rune) bool { return IsDigit(r) || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F' }

// Check if a rune is an octal digit (0-7)
func IsOctal(r rune) bool { return r >= '0' && r <= '7' }

// Check if a rune is a binary digit (0 or 1)
func IsBinary(r rune) bool { return r == '0' || r == '1' }

// Check if a rune is a lowercase letter (a-z)
func IsLower(r rune) bool { return r >= 'a' && r <= 'z' }

// Check if a rune is an uppercase letter (A-Z)
func IsUpper(r rune) bool { return r >= 'A' && r <= 'Z' }

// Check if a rune is a punctuation mark
func IsPunctuation(r rune) bool {
	return IsOneOf(r, '.', ',', '!', '?', ';', ':', '"', '\'', '-', '(', ')', '[', ']', '{', '}')
}

// Check if a rune is a symbol or operator
func IsSymbol(r rune) bool {
	return IsOneOf(r, '+', '-', '*', '/', '=', '<', '>', '&', '|', '^', '~', '!', '@', '#', '$', '%')
}

// Check if a rune is a tab character
func IsTab(r rune) bool { return r == '\t' }

// Check if a rune is an ASCII character (0-127)
func IsAscii(r rune) bool { return r >= 0 && r <= 127 }

// Check if a rune is printable ASCII
func IsPrintableAscii(r rune) bool { return r >= 32 && r <= 126 }

// Check if a rune is a vowel (a, e, i, o, u - case insensitive)
func IsVowel(r rune) bool { return IsOneOf(r, 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U') }

// Check if a rune is a consonant (letter but not vowel)
func IsConsonant(r rune) bool { return IsLetter(r) && !IsVowel(r) }

// Check if a rune is a quote character (single or double)
func IsQuote(r rune) bool { return r == '"' || r == '\'' || r == '`' }

// Convert a rune to lowercase if it's uppercase
func ToLower(r rune) rune {
	if IsUpper(r) {
		return r + ('a' - 'A')
	}
	return r
}

// Convert a rune to uppercase if it's lowercase
func ToUpper(r rune) rune {
	if IsLower(r) {
		return r - ('a' - 'A')
	}
	return r
}
