package dsx

// Optional represents a value that may or may not be present. It provides methods
// to check for presence, retrieve the value, and provide default values if the
// value is not present.
type Optional[T any] struct {
	value   T
	present bool
}

// Some creates an Optional instance that contains a value, indicating that the
// value is present.
func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

// None creates an Optional instance that does not contain a value, indicating
// that the value is not present. The zero value of the type T is used as the
// value, but it is not considered present.
func None[T any]() Optional[T] {
	var zero T
	return Optional[T]{value: zero, present: false}
}

// Has returns true if the Optional contains a value, and false if it does not.
func (o Optional[T]) Has() bool {
	return o.present
}

// Get returns the value contained in the Optional. If the Optional does not
// contain a value, it returns the zero value of type T. It is the caller's
// responsibility to check if the value is present using Has() before calling Get()
// to avoid unintended consequences of using a zero value.
func (o Optional[T]) Get() T {
	return o.value
}

// GetOr returns the value contained in the Optional if it is present. If the
// Optional does not contain a value, it returns the provided defaultValue instead. This method
// allows for a convenient way to specify a fallback value when the Optional is empty.
func (o Optional[T]) GetOr(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// GetOk returns the value contained in the Optional and a boolean indicating whether the value is present.
// This method provides a convenient way to check for presence and retrieve the value in a single call.
func (o Optional[T]) GetOk() (T, bool) {
	return o.value, o.present
}
