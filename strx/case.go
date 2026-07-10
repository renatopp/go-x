// https://github.com/Ma124/strcase
/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strx

import (
	"strings"
	"sync"
)

// "ID": "id",
var uppercaseAcronym = sync.Map{}

// ToLower converts a string to lower case. Example:  `MyHello, World` becomes
// `myhello, world`.
func ToLower(s string) string { return strings.ToLower(s) }

// ToUpper converts a string to UPPER CASE. Example:  `MyHello, World` becomes
// `MYHELLO, WORLD`.
func ToUpper(s string) string { return strings.ToUpper(s) }

// ToTitle converts a string to Title Case. The first letter of each word is
// mapped to upper case. Words are defined as sequences of letters separated
// by non-letter characters. Example: `MyHello, World` becomes `Myhello, World`.
func ToTitle(s string) string {
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := true

	for _, v := range s {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNumber := v >= '0' && v <= '9'
		if capNext && vIsLow {
			v = v - 'a' + 'A'
		} else if !capNext && vIsCap {
			v = v - 'A' + 'a'
		}
		n.WriteRune(v)
		capNext = !(vIsCap || vIsLow || vIsNumber)
	}
	return n.String()
}

// ToDelimiter converts a string to a delimited-case string. For example,
// `MyHello, World` becomes `my⚡hello⚡world` if the delimiter is `⚡`.
func ToDelimiter(s string, d string) string {
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))

	lastType := "delimiter" // "cap", "low", "num", "delimiter"
	for _, v := range s {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNumber := v >= '0' && v <= '9'
		vIsSpecial := !vIsCap && !vIsLow && !vIsNumber

		if vIsCap {
			v = v - 'A' + 'a'
		}

		switch {
		case vIsSpecial:
			if lastType != "delimiter" {
				n.WriteString(d)
				lastType = "delimiter"
			}

		case vIsCap:
			if lastType != "cap" && lastType != "delimiter" {
				n.WriteString(d)
			}
			n.WriteRune(v)
			lastType = "cap"

		case vIsLow:
			if lastType == "number" {
				n.WriteString(d)
			}
			n.WriteRune(v)
			lastType = "low"

		case vIsNumber:
			if lastType != "number" && lastType != "delimiter" {
				n.WriteString(d)
			}
			n.WriteRune(v)
			lastType = "number"
		}
	}
	return strings.Trim(n.String(), d)
}

// ToSnake converts a string to snake_case. Example: `MyHello, World` becomes
// `my_hello_world`.
func ToSnake(s string) string { return ToDelimiter(s, "_") }

// ToUpperSnake converts a string to UPPER_SNAKE_CASE. Example:
// `MyHello, World` becomes `MY_HELLO_WORLD`.
func ToUpperSnake(s string) string { return ToUpper(ToDelimiter(s, "_")) }

// ToKebab converts a string to kebab-case. Example: `MyHello, World` becomes
// `my-hello-world`.
func ToKebab(s string) string { return ToDelimiter(s, "-") }

// ToUpperKebab converts a string to UPPER-KEBAB-CASE. Example:
// `MyHello, World` becomes `MY-HELLO-WORLD`.
func ToUpperKebab(s string) string { return ToUpper(ToDelimiter(s, "-")) }

// ToCamel converts a string to CamelCase. Example: `MyHello, World` becomes
// `myHelloWorld`.
func ToCamel(s string) string { return toInitCase(s, false) }

// ToPascal converts a string to PascalCase. Example: `MyHello, World` becomes
// `MyHelloWorld`.
func ToPascal(s string) string { return toInitCase(s, true) }

// toInitCase converts a string to camelCase or PascalCase depending on the value of
// initCase.
func toInitCase(s string, initCase bool) string {
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	lastWasLow := false
	for _, v := range s {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNumber := v >= '0' && v <= '9'
		vIsSpecial := !vIsCap && !vIsLow && !vIsNumber

		if vIsSpecial {
			capNext = true
			lastWasLow = false

		} else if vIsNumber {
			n.WriteRune(v)
			capNext = true
			lastWasLow = false

		} else if vIsCap {
			if !capNext && !lastWasLow {
				v = v - 'A' + 'a'
			}
			n.WriteRune(v)
			capNext = false
			lastWasLow = false

		} else if vIsLow {
			if capNext {
				v = v - 'a' + 'A'
			}
			n.WriteRune(v)
			capNext = false
			lastWasLow = true
		}
	}
	return n.String()
}
