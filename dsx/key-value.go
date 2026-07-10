package dsx

// Represents a simple key-value pair.
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

// Creates a new KeyValue instance with the provided key and value.
func NewKeyValue[K comparable, V any](key K, value V) *KeyValue[K, V] {
	return &KeyValue[K, V]{
		Key:   key,
		Value: value,
	}
}
