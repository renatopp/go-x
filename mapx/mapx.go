package mapx

import (
	"maps"
)

// ACCESSORS ------------------------------------------------------------------

// GetOr returns the value for the given key if it exists, otherwise returns
// the default value.
func GetOr[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

// Keys returns a slice of all keys in the map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values in the map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// KeyOf returns the key for the given value if it exists, otherwise returns
// the zero value of the key type and false.
func KeyOf[K comparable, V comparable](m map[K]V, value V) (K, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// KeyOfFunc returns the key for the first value that satisfies the given
// predicate function, otherwise returns the zero value of the key type and
// false.
func KeyOfFunc[K comparable, V any](m map[K]V, fn func(K, V) bool) (K, bool) {
	for k, v := range m {
		if fn(k, v) {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// Size returns the number of key-value pairs in the map.
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}

// MANIPULATORS ---------------------------------------------------------------

// Concat returns a new map that is the concatenation of the given maps. If
// there are duplicate keys, the value from the last map will be used.
func Concat[K comparable, V any](m map[K]V, others ...map[K]V) map[K]V {
	items := make(map[K]V)
	maps.Copy(items, m)
	for _, other := range others {
		maps.Copy(items, other)
	}
	return items
}

// Assign copies the key-value pairs from the given maps into the first map. If
// there are duplicate keys, the value from the last map will be used.
func Assign[K comparable, V any](m map[K]V, others ...map[K]V) {
	for _, other := range others {
		maps.Copy(m, other)
	}
}

// Remove removes the given keys from the map.
func Remove[K comparable, V any](m map[K]V, keys ...K) {
	for _, key := range keys {
		delete(m, key)
	}
}

// Clear removes all key-value pairs from the map.
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// CHECKS ---------------------------------------------------------------------

// ContainsKey returns true if the map contains the given key.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsValue returns true if the map contains the given value.
func ContainsValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc returns true if the map contains a key-value pair that
// satisfies the given predicate function.
func ContainsFunc[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// ITERATORS ------------------------------------------------------------------

// ForEach calls the given function for each key-value pair in the map.
func ForEach[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

// ForEachKey calls the given function for each key in the map.
func ForEachValue[K comparable, V any](m map[K]V, fn func(V)) {
	for _, v := range m {
		fn(v)
	}
}

// ForEachKey calls the given function for each key in the map.
func ForEachKey[K comparable, V any](m map[K]V, fn func(K)) {
	for k := range m {
		fn(k)
	}
}
