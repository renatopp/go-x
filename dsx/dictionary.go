package dsx

import (
	"iter"
	"maps"
)

// Dictionary is a simple string-keyed collection of arbitrary values, similar
// to a JSON object. It wraps a map[string]any and provides convenience methods
// for accessing, mutating, and iterating over its entries.
type Dictionary struct {
	data map[string]any
}

// NewDictionary creates a new empty dictionary.
func NewDictionary() *Dictionary {
	return &Dictionary{
		data: make(map[string]any),
	}
}

// NewDictionaryFrom creates a new dictionary from the given map. The map is
// used directly, so modifications to the dictionary will affect the original
// map and vice versa.
func NewDictionaryFrom(data map[string]any) *Dictionary {
	if data == nil {
		data = make(map[string]any)
	}
	return &Dictionary{
		data: data,
	}
}

// Get returns the value associated with the given key. If the key does not
// exist, it returns nil.
func (d *Dictionary) Get(key string) any {
	return d.data[key]
}

// GetOr returns the value associated with the given key. If the key does not
// exist, it returns the provided default value.
func (d *Dictionary) GetOr(key string, v any) any {
	if val, ok := d.data[key]; ok {
		return val
	}
	return v
}

// GetOk returns the value associated with the given key and true if the key
// exists. Otherwise, it returns nil and false.
func (d *Dictionary) GetOk(key string) (any, bool) {
	v, ok := d.data[key]
	return v, ok
}

// Set associates the given value with the given key, overwriting any existing
// value for that key.
func (d *Dictionary) Set(key string, value any) {
	d.data[key] = value
}

// Delete removes the given keys from the dictionary. Keys that do not exist
// are ignored.
func (d *Dictionary) Delete(keys ...string) {
	for _, key := range keys {
		delete(d.data, key)
	}
}

// ContainsKey returns true if the dictionary contains the given key.
func (d *Dictionary) ContainsKey(key string) bool {
	_, ok := d.data[key]
	return ok
}

// ContainsValue returns true if the dictionary contains the given value.
func (d *Dictionary) ContainsValue(value any) bool {
	for _, v := range d.data {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc returns true if the dictionary contains a key-value pair that
// satisfies the given predicate function.
func (d *Dictionary) ContainsFunc(fn func(string, any) bool) bool {
	for k, v := range d.data {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Size returns the number of key-value pairs in the dictionary.
func (d *Dictionary) Size() int {
	return len(d.data)
}

// Clear removes all key-value pairs from the dictionary, leaving it empty.
func (d *Dictionary) Clear() {
	d.data = make(map[string]any)
}

// Clone creates and returns a new dictionary that is a shallow copy of the
// current dictionary. If the values are reference types, they will be shared
// between the original and the clone.
func (d *Dictionary) Clone() *Dictionary {
	return NewDictionaryFrom(maps.Clone(d.data))
}

// Keys returns a slice of all keys in the dictionary.
func (d *Dictionary) Keys() []string {
	keys := make([]string, 0, len(d.data))
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values in the dictionary.
func (d *Dictionary) Values() []any {
	values := make([]any, 0, len(d.data))
	for _, v := range d.data {
		values = append(values, v)
	}
	return values
}

// ToMap returns the underlying map[string]any. Modifying the returned map
// will affect the dictionary and vice versa.
func (d *Dictionary) ToMap() map[string]any {
	return d.data
}

// Concat returns a new dictionary that is the concatenation of the current
// dictionary with the given dictionaries. If there are duplicate keys, the
// value from the last dictionary wins.
func (d *Dictionary) Concat(others ...*Dictionary) *Dictionary {
	data := maps.Clone(d.data)
	for _, other := range others {
		maps.Copy(data, other.data)
	}
	return NewDictionaryFrom(data)
}

// Assign copies the key-value pairs from the given dictionaries into the
// current dictionary. If there are duplicate keys, the value from the last
// dictionary wins.
func (d *Dictionary) Assign(others ...*Dictionary) {
	for _, other := range others {
		maps.Copy(d.data, other.data)
	}
}

// Equal returns true if the two dictionaries have the same keys and the same
// values for those keys.
func (d *Dictionary) Equal(other *Dictionary) bool {
	return maps.EqualFunc(d.data, other.data, func(a, b any) bool {
		return a == b
	})
}

// Iter returns an iterator that yields the key and value of each entry in the
// dictionary.
func (d *Dictionary) Iter() iter.Seq2[string, any] {
	return maps.All(d.data)
}

// ForEach calls the given function for each key-value pair in the dictionary.
func (d *Dictionary) ForEach(fn func(string, any)) {
	for k, v := range d.data {
		fn(k, v)
	}
}
