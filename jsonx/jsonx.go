package jsonx

import (
	"bytes"
	"encoding/json"
)

// force is a helper to ignore errors in functions that return (T, error) and just return T.
func force[T any](a T, _ error) T {
	return a
}

// Marshal is a wrapper around json.Marshal.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalString is a wrapper around json.Marshal that returns a string instead of []byte.
func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// MarshalIndent is a wrapper around json.MarshalIndent.
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalIndentString is a wrapper around json.MarshalIndent that returns a string instead of []byte.
func MarshalIndentString(v any, prefix, indent string) (string, error) {
	b, err := json.MarshalIndent(v, prefix, indent)
	return string(b), err
}

// ForceMarshal is a wrapper around Marshal that ignores the error and just returns the []byte.
func ForceMarshal(v any) []byte {
	return force(json.Marshal(v))
}

// ForceMarshalString is a wrapper around MarshalString that ignores the error and just returns the string.
func ForceMarshalString(v any) string {
	return force(MarshalString(v))
}

// ForceMarshalIndent is a wrapper around MarshalIndent that ignores the error and just returns the []byte.
func ForceMarshalIndent(v any, prefix, indent string) []byte {
	return force(json.MarshalIndent(v, prefix, indent))
}

// ForceMarshalIndentString is a wrapper around MarshalIndentString that ignores the error and just returns the string.
func ForceMarshalIndentString(v any, prefix, indent string) string {
	return force(MarshalIndentString(v, prefix, indent))
}

// Unmarshal is a wrapper around json.Unmarshal.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// UnmarshalString is a wrapper around json.Unmarshal that takes a string instead of []byte.
func UnmarshalString(s string, v any) error {
	return json.Unmarshal([]byte(s), v)
}

// UnmarshalAs is a wrapper around json.Unmarshal that returns the unmarshaled
// value defined by a type parameter instead of taking a pointer.
func UnmarshalAs[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return v, err
}

// UnmarshalStringAs is a wrapper around json.Unmarshal that takes a string
// instead of []byte and returns the unmarshaled value defined by a type parameter.
func UnmarshalStringAs[T any](s string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	return v, err
}

// ForceUnmarshalAs is a wrapper around UnmarshalAs that ignores the error and just returns the unmarshaled value.
func ForceUnmarshalAs[T any](data []byte) T {
	return force(UnmarshalAs[T](data))
}

// ForceUnmarshalStringAs is a wrapper around UnmarshalStringAs that ignores the error and just returns the unmarshaled value.
func ForceUnmarshalStringAs[T any](s string) T {
	return force(UnmarshalStringAs[T](s))
}

// PrettyPrintJSON is a wrapper around json.Indent that returns the pretty-printed JSON as []byte.
func PrettyPrint(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, data, "", "  ")
	return buf.Bytes(), err
}

// PrettyPrintJSONString is a wrapper around PrettyPrintJSON that takes a string and returns a pretty-printed JSON string.
func PrettyPrintString(s string) (string, error) {
	b, err := PrettyPrint([]byte(s))
	return string(b), err
}

// ForcePrettyPrintJSON is a wrapper around PrettyPrintJSON that ignores the error and just returns the pretty-printed JSON as []byte.
func ForcePrettyPrint(data []byte) []byte {
	return force(PrettyPrint(data))
}

// ForcePrettyPrintJSONString is a wrapper around PrettyPrintJSONString that ignores the error and just returns the pretty-printed JSON string.
func ForcePrettyPrintString(s string) string {
	return force(PrettyPrintString(s))
}

// Valid is a wrapper around json.Valid that takes a []byte and returns whether it is valid JSON.
func IsValid(data []byte) bool {
	return json.Valid(data)
}

// ValidString is a wrapper around json.Valid that takes a string and returns whether it is valid JSON.
func IsValidString(s string) bool {
	return json.Valid([]byte(s))
}

// Compact is a wrapper around json.Compact that takes a []byte and returns the compacted JSON as []byte.
func Compact(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	err := json.Compact(&buf, data)
	return buf.Bytes(), err
}

// CompactString is a wrapper around Compact that takes a string and returns the compacted JSON as a string.
func CompactString(s string) (string, error) {
	b, err := Compact([]byte(s))
	return string(b), err
}

// ForceCompact is a wrapper around Compact that ignores the error and just returns the compacted JSON as []byte.
func ForceCompact(data []byte) []byte {
	return force(Compact(data))
}

// ForceCompactString is a wrapper around CompactString that ignores the error and just returns the compacted JSON as a string.
func ForceCompactString(s string) string {
	return force(CompactString(s))
}
