// Package yamlx extends the goccy/go-yaml package with a set of helpers
// mirroring the interface of the jsonx package.
package yamlx

import (
	"github.com/goccy/go-yaml"
)

// force is a helper to ignore errors in functions that return (T, error) and just return T.
func force[T any](a T, _ error) T {
	return a
}

// Marshal is a wrapper around yaml.Marshal.
func Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

// MarshalString is a wrapper around yaml.Marshal that returns a string instead of []byte.
func MarshalString(v any) (string, error) {
	b, err := yaml.Marshal(v)
	return string(b), err
}

// MarshalIndent is a wrapper around yaml.MarshalWithOptions that sets the indentation width, in spaces.
func MarshalIndent(v any, indent int) ([]byte, error) {
	return yaml.MarshalWithOptions(v, yaml.Indent(indent))
}

// MarshalIndentString is a wrapper around MarshalIndent that returns a string instead of []byte.
func MarshalIndentString(v any, indent int) (string, error) {
	b, err := MarshalIndent(v, indent)
	return string(b), err
}

// ForceMarshal is a wrapper around Marshal that ignores the error and just returns the []byte.
func ForceMarshal(v any) []byte {
	return force(yaml.Marshal(v))
}

// ForceMarshalString is a wrapper around MarshalString that ignores the error and just returns the string.
func ForceMarshalString(v any) string {
	return force(MarshalString(v))
}

// ForceMarshalIndent is a wrapper around MarshalIndent that ignores the error and just returns the []byte.
func ForceMarshalIndent(v any, indent int) []byte {
	return force(MarshalIndent(v, indent))
}

// ForceMarshalIndentString is a wrapper around MarshalIndentString that ignores the error and just returns the string.
func ForceMarshalIndentString(v any, indent int) string {
	return force(MarshalIndentString(v, indent))
}

// Unmarshal is a wrapper around yaml.Unmarshal.
func Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

// UnmarshalString is a wrapper around yaml.Unmarshal that takes a string instead of []byte.
func UnmarshalString(s string, v any) error {
	return yaml.Unmarshal([]byte(s), v)
}

// UnmarshalAs is a wrapper around yaml.Unmarshal that returns the unmarshaled
// value defined by a type parameter instead of taking a pointer.
func UnmarshalAs[T any](data []byte) (T, error) {
	var v T
	err := yaml.Unmarshal(data, &v)
	return v, err
}

// UnmarshalStringAs is a wrapper around yaml.Unmarshal that takes a string
// instead of []byte and returns the unmarshaled value defined by a type parameter.
func UnmarshalStringAs[T any](s string) (T, error) {
	var v T
	err := yaml.Unmarshal([]byte(s), &v)
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

// PrettyPrint reformats YAML data with consistent indentation and key ordering preserved.
func PrettyPrint(data []byte) ([]byte, error) {
	var v yaml.MapSlice
	if err := yaml.UnmarshalWithOptions(data, &v, yaml.UseOrderedMap()); err != nil {
		return nil, err
	}
	return yaml.Marshal(v)
}

// PrettyPrintString is a wrapper around PrettyPrint that takes a string and returns a pretty-printed YAML string.
func PrettyPrintString(s string) (string, error) {
	b, err := PrettyPrint([]byte(s))
	return string(b), err
}

// ForcePrettyPrint is a wrapper around PrettyPrint that ignores the error and just returns the pretty-printed YAML as []byte.
func ForcePrettyPrint(data []byte) []byte {
	return force(PrettyPrint(data))
}

// ForcePrettyPrintString is a wrapper around PrettyPrintString that ignores the error and just returns the pretty-printed YAML string.
func ForcePrettyPrintString(s string) string {
	return force(PrettyPrintString(s))
}

// IsValid reports whether data is valid YAML.
func IsValid(data []byte) bool {
	var v any
	return yaml.Unmarshal(data, &v) == nil
}

// IsValidString reports whether s is valid YAML.
func IsValidString(s string) bool {
	return IsValid([]byte(s))
}

// Compact reformats YAML data into a single-line flow-style representation.
func Compact(data []byte) ([]byte, error) {
	var v yaml.MapSlice
	if err := yaml.UnmarshalWithOptions(data, &v, yaml.UseOrderedMap()); err != nil {
		return nil, err
	}
	return yaml.MarshalWithOptions(v, yaml.Flow(true))
}

// CompactString is a wrapper around Compact that takes a string and returns the compacted YAML as a string.
func CompactString(s string) (string, error) {
	b, err := Compact([]byte(s))
	return string(b), err
}

// ForceCompact is a wrapper around Compact that ignores the error and just returns the compacted YAML as []byte.
func ForceCompact(data []byte) []byte {
	return force(Compact(data))
}

// ForceCompactString is a wrapper around CompactString that ignores the error and just returns the compacted YAML as a string.
func ForceCompactString(s string) string {
	return force(CompactString(s))
}
