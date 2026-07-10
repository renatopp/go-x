package runex_test

import (
	"testing"

	"github.com/renatopp/go-x/runex"
	"github.com/renatopp/go-x/testx"
)

func TestIsDigit(t *testing.T) {
	testx.True(t, runex.IsDigit('0'))
	testx.True(t, runex.IsDigit('1'))
	testx.True(t, runex.IsDigit('2'))
	testx.True(t, runex.IsDigit('3'))
	testx.True(t, runex.IsDigit('4'))
	testx.True(t, runex.IsDigit('5'))
	testx.True(t, runex.IsDigit('6'))
	testx.True(t, runex.IsDigit('7'))
	testx.True(t, runex.IsDigit('8'))
	testx.True(t, runex.IsDigit('9'))
	testx.False(t, runex.IsDigit('a'))
	testx.False(t, runex.IsDigit('A'))
	testx.False(t, runex.IsDigit(' '))
	testx.False(t, runex.IsDigit(';'))
}

func TestIsLetter(t *testing.T) {
	testx.True(t, runex.IsLetter('a'))
	testx.True(t, runex.IsLetter('b'))
	testx.True(t, runex.IsLetter('c'))
	testx.True(t, runex.IsLetter('A'))
	testx.True(t, runex.IsLetter('B'))
	testx.True(t, runex.IsLetter('W'))
	testx.True(t, runex.IsLetter('X'))
	testx.False(t, runex.IsLetter('0'))
	testx.False(t, runex.IsLetter('9'))
	testx.False(t, runex.IsLetter(' '))
	testx.False(t, runex.IsLetter(';'))
}

func TestIsNewline(t *testing.T) {
	testx.True(t, runex.IsNewline('\n'))
	testx.False(t, runex.IsNewline('a'))
}

func TestIsWhitespace(t *testing.T) {
	testx.True(t, runex.IsWhitespace(' '))
	testx.True(t, runex.IsWhitespace('\n'))
	testx.True(t, runex.IsWhitespace('\r'))
	testx.False(t, runex.IsWhitespace('a'))
	testx.False(t, runex.IsWhitespace('A'))
	testx.False(t, runex.IsWhitespace('0'))
	testx.False(t, runex.IsWhitespace('9'))
	testx.False(t, runex.IsWhitespace(';'))
}

func TestIsSpace(t *testing.T) {
	testx.True(t, runex.IsSpace(' '))
	testx.True(t, runex.IsSpace('\r'))
}

func TestIsEof(t *testing.T) {
	testx.True(t, runex.IsEof(0))
	testx.False(t, runex.IsEof('a'))
	testx.False(t, runex.IsEof('A'))
	testx.False(t, runex.IsEof('0'))
	testx.False(t, runex.IsEof('9'))
	testx.False(t, runex.IsEof(';'))
}

func TestIsAlphaNumeric(t *testing.T) {
	testx.True(t, runex.IsAlphaNumeric('a'))
	testx.True(t, runex.IsAlphaNumeric('A'))
	testx.True(t, runex.IsAlphaNumeric('0'))
	testx.True(t, runex.IsAlphaNumeric('9'))
	testx.False(t, runex.IsAlphaNumeric(' '))
	testx.False(t, runex.IsAlphaNumeric(';'))
}

func TestIsOneOf(t *testing.T) {
	testx.True(t, runex.IsOneOf('a', 'a', 'b', 'c'))
	testx.True(t, runex.IsOneOf('b', 'a', 'b', 'c'))
	testx.True(t, runex.IsOneOf('c', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf('d', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf(' ', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf(';', 'a', 'b', 'c'))
}

func TestIsHexadecimal(t *testing.T) {
	testx.True(t, runex.IsHexadecimal('0'))
	testx.True(t, runex.IsHexadecimal('1'))
	testx.True(t, runex.IsHexadecimal('2'))
	testx.True(t, runex.IsHexadecimal('3'))
	testx.True(t, runex.IsHexadecimal('4'))
	testx.True(t, runex.IsHexadecimal('5'))
	testx.True(t, runex.IsHexadecimal('6'))
	testx.True(t, runex.IsHexadecimal('7'))
	testx.True(t, runex.IsHexadecimal('8'))
	testx.True(t, runex.IsHexadecimal('9'))
	testx.True(t, runex.IsHexadecimal('a'))
	testx.True(t, runex.IsHexadecimal('b'))
	testx.True(t, runex.IsHexadecimal('c'))
	testx.True(t, runex.IsHexadecimal('d'))
	testx.True(t, runex.IsHexadecimal('e'))
	testx.True(t, runex.IsHexadecimal('f'))
	testx.True(t, runex.IsHexadecimal('A'))
	testx.True(t, runex.IsHexadecimal('B'))
	testx.True(t, runex.IsHexadecimal('C'))
	testx.True(t, runex.IsHexadecimal('D'))
	testx.True(t, runex.IsHexadecimal('E'))
	testx.True(t, runex.IsHexadecimal('F'))
	testx.False(t, runex.IsHexadecimal(' '))
	testx.False(t, runex.IsHexadecimal(';'))
}

func TestIsOctal(t *testing.T) {
	testx.True(t, runex.IsOctal('0'))
	testx.True(t, runex.IsOctal('1'))
	testx.True(t, runex.IsOctal('2'))
	testx.True(t, runex.IsOctal('3'))
	testx.True(t, runex.IsOctal('4'))
	testx.True(t, runex.IsOctal('5'))
	testx.True(t, runex.IsOctal('6'))
	testx.True(t, runex.IsOctal('7'))
	testx.False(t, runex.IsOctal('8'))
	testx.False(t, runex.IsOctal('9'))
	testx.False(t, runex.IsOctal('a'))
	testx.False(t, runex.IsOctal('A'))
	testx.False(t, runex.IsOctal(' '))
	testx.False(t, runex.IsOctal(';'))
}

func TestIsBinary(t *testing.T) {
	testx.True(t, runex.IsBinary('0'))
	testx.True(t, runex.IsBinary('1'))
	testx.False(t, runex.IsBinary('2'))
	testx.False(t, runex.IsBinary('3'))
	testx.False(t, runex.IsBinary('4'))
	testx.False(t, runex.IsBinary('5'))
	testx.False(t, runex.IsBinary('6'))
	testx.False(t, runex.IsBinary('7'))
	testx.False(t, runex.IsBinary('8'))
	testx.False(t, runex.IsBinary('9'))
	testx.False(t, runex.IsBinary('a'))
	testx.False(t, runex.IsBinary('A'))
	testx.False(t, runex.IsBinary(' '))
	testx.False(t, runex.IsBinary(';'))
}

func TestIsLower(t *testing.T) {
	testx.True(t, runex.IsLower('a'))
	testx.True(t, runex.IsLower('b'))
	testx.True(t, runex.IsLower('z'))
	testx.False(t, runex.IsLower('A'))
	testx.False(t, runex.IsLower('Z'))
	testx.False(t, runex.IsLower('0'))
	testx.False(t, runex.IsLower('9'))
	testx.False(t, runex.IsLower(' '))
	testx.False(t, runex.IsLower(';'))
}

func TestIsUpper(t *testing.T) {
	testx.True(t, runex.IsUpper('A'))
	testx.True(t, runex.IsUpper('B'))
	testx.True(t, runex.IsUpper('Z'))
	testx.False(t, runex.IsUpper('a'))
	testx.False(t, runex.IsUpper('z'))
	testx.False(t, runex.IsUpper('0'))
	testx.False(t, runex.IsUpper('9'))
	testx.False(t, runex.IsUpper(' '))
	testx.False(t, runex.IsUpper(';'))
}

func TestIsPunctuation(t *testing.T) {
	testx.True(t, runex.IsPunctuation('.'))
	testx.True(t, runex.IsPunctuation(','))
	testx.True(t, runex.IsPunctuation('!'))
	testx.True(t, runex.IsPunctuation('?'))
	testx.True(t, runex.IsPunctuation(';'))
	testx.True(t, runex.IsPunctuation(':'))
	testx.True(t, runex.IsPunctuation('"'))
	testx.True(t, runex.IsPunctuation('\''))
	testx.True(t, runex.IsPunctuation('-'))
	testx.True(t, runex.IsPunctuation('('))
	testx.True(t, runex.IsPunctuation(')'))
	testx.True(t, runex.IsPunctuation('['))
	testx.True(t, runex.IsPunctuation(']'))
	testx.True(t, runex.IsPunctuation('{'))
	testx.True(t, runex.IsPunctuation('}'))
	testx.False(t, runex.IsPunctuation('a'))
	testx.False(t, runex.IsPunctuation('A'))
	testx.False(t, runex.IsPunctuation('0'))
	testx.False(t, runex.IsPunctuation('9'))
	testx.False(t, runex.IsPunctuation(' '))
}

func TestIsSymbolOrOperator(t *testing.T) {
	testx.True(t, runex.IsSymbol('+'))
	testx.True(t, runex.IsSymbol('-'))
	testx.True(t, runex.IsSymbol('*'))
	testx.True(t, runex.IsSymbol('/'))
}

func TestIsTab(t *testing.T) {
	testx.True(t, runex.IsTab('\t'))
	testx.False(t, runex.IsTab('t'))
}

func TestIsAscii(t *testing.T) {
	testx.True(t, runex.IsAscii('a'))
	testx.True(t, runex.IsAscii('A'))
	testx.False(t, runex.IsAscii('ñ'))
	testx.False(t, runex.IsAscii('€'))
	testx.False(t, runex.IsAscii('🧹'))
}

func TestIsPrintableAscii(t *testing.T) {
	testx.True(t, runex.IsPrintableAscii('a'))
	testx.True(t, runex.IsPrintableAscii('A'))
	testx.True(t, runex.IsPrintableAscii('0'))
	testx.True(t, runex.IsPrintableAscii('9'))
	testx.True(t, runex.IsPrintableAscii(' '))
	testx.False(t, runex.IsPrintableAscii('\n'))
}

func TestIsVowel(t *testing.T) {
	testx.True(t, runex.IsVowel('a'))
	testx.True(t, runex.IsVowel('e'))
	testx.True(t, runex.IsVowel('i'))
	testx.True(t, runex.IsVowel('o'))
	testx.True(t, runex.IsVowel('u'))
}

func TestIsConsonant(t *testing.T) {
	testx.True(t, runex.IsConsonant('b'))
	testx.True(t, runex.IsConsonant('c'))
	testx.False(t, runex.IsConsonant('a'))
}

func TestIsQuote(t *testing.T) {
	testx.True(t, runex.IsQuote('"'))
	testx.True(t, runex.IsQuote('\''))
	testx.True(t, runex.IsQuote('`'))
	testx.False(t, runex.IsQuote('a'))
	testx.False(t, runex.IsQuote('A'))
	testx.False(t, runex.IsQuote('0'))
	testx.False(t, runex.IsQuote('9'))
	testx.False(t, runex.IsQuote(' '))
	testx.False(t, runex.IsQuote(';'))
}

func TestToLower(t *testing.T) {
	testx.Equal(t, runex.ToLower('A'), 'a')
	testx.Equal(t, runex.ToLower('Z'), 'z')
	testx.Equal(t, runex.ToLower('a'), 'a')
	testx.Equal(t, runex.ToLower('z'), 'z')
	testx.Equal(t, runex.ToLower('0'), '0')
	testx.Equal(t, runex.ToLower('9'), '9')
	testx.Equal(t, runex.ToLower(' '), ' ')
	testx.Equal(t, runex.ToLower(';'), ';')
}

func TestToUpper(t *testing.T) {
	testx.Equal(t, runex.ToUpper('a'), 'A')
	testx.Equal(t, runex.ToUpper('z'), 'Z')
	testx.Equal(t, runex.ToUpper('A'), 'A')
	testx.Equal(t, runex.ToUpper('Z'), 'Z')
	testx.Equal(t, runex.ToUpper('0'), '0')
	testx.Equal(t, runex.ToUpper('9'), '9')
	testx.Equal(t, runex.ToUpper(' '), ' ')
	testx.Equal(t, runex.ToUpper(';'), ';')
}
