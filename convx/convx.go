package convx

import (
	"fmt"
	"strconv"
)

// ToString converts any value to its string representation.
// It uses the default %v format specifier to convert the value to a string.
func ToString(v any) string {
	return fmt.Sprintf("%v", v)
}

// ToInt converts any valid value to an int.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ToInt(v any) (int, error) {
	i64, err := ToInt64(v)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

// ForceToInt converts any valid value to an int, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToInt(v any) int {
	i, _ := ToInt(v)
	return i
}

// ToInt8 converts any valid value to an int.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ToInt8(v any) (int8, error) {
	i64, err := ToInt64(v)
	if err != nil {
		return 0, err
	}
	if i64 < -128 || i64 > 127 {
		return 0, fmt.Errorf("value %d overflows int8", i64)
	}
	return int8(i64), nil
}

// ForceToInt8 converts any valid value to an int8, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ForceToInt8(v any) int8 {
	i, _ := ToInt8(v)
	return i
}

// ToInt16 converts any valid value to an int16.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ToInt16(v any) (int16, error) {
	i64, err := ToInt64(v)
	if err != nil {
		return 0, err
	}
	if i64 < -32768 || i64 > 32767 {
		return 0, fmt.Errorf("value %d overflows int16", i64)
	}
	return int16(i64), nil
}

// ForceToInt16 converts any valid value to an int16, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToInt16(v any) int16 {
	i, _ := ToInt16(v)
	return i
}

// ToInt32 converts any valid value to an int32.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ToInt32(v any) (int32, error) {
	i64, err := ToInt64(v)
	if err != nil {
		return 0, err
	}
	if i64 < -2147483648 || i64 > 2147483647 {
		return 0, fmt.Errorf("value %d overflows int32", i64)
	}
	return int32(i64), nil
}

// ForceToInt32 converts any valid value to an int32, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToInt32(v any) int32 {
	i, _ := ToInt32(v)
	return i
}

// ToInt64 converts any valid value to an int64.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows.
func ToInt64(v any) (int64, error) {
	switch val := any(v).(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		if val > 9223372036854775807 {
			return 0, fmt.Errorf("value %d overflows int64", val)
		}
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case string:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string to int64: %w", err)
		}
		return int64(v), nil
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// ForceToInt64 converts any valid value to an int64, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToInt64(v any) int64 {
	i, _ := ToInt64(v)
	return i
}

// ToUint converts any valid value to an uint.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows or is negative.
func ToUint(v any) (uint, error) {
	u, err := ToUint64(v)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

// ForceToUint converts any valid value to an uint, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToUint(v any) uint {
	u, _ := ToUint(v)
	return u
}

// ToUint8 converts any valid value to an uint8.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows or is negative.
func ToUint8(v any) (uint8, error) {
	u64, err := ToUint64(v)
	if err != nil {
		return 0, err
	}
	if u64 > 255 {
		return 0, fmt.Errorf("value %d overflows uint8", u64)
	}
	return uint8(u64), nil
}

// ForceToUint8 converts any valid value to an uint8, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToUint8(v any) uint8 {
	u, _ := ToUint8(v)
	return u
}

// ToUint16 converts any valid value to an uint16.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows or is negative.
func ToUint16(v any) (uint16, error) {
	u64, err := ToUint64(v)
	if err != nil {
		return 0, err
	}
	if u64 > 65535 {
		return 0, fmt.Errorf("value %d overflows uint16", u64)
	}
	return uint16(u64), nil
}

// ForceToUint16 converts any valid value to an uint16, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToUint16(v any) uint16 {
	u, _ := ToUint16(v)
	return u
}

// ToUint32 converts any valid value to an uint32.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number overflows or is negative.
func ToUint32(v any) (uint32, error) {
	u64, err := ToUint64(v)
	if err != nil {
		return 0, err
	}
	if u64 > 4294967295 {
		return 0, fmt.Errorf("value %d overflows uint32", u64)
	}
	return uint32(u64), nil
}

// ForceToUint32 converts any valid value to an uint32, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToUint32(v any) uint32 {
	u, _ := ToUint32(v)
	return u
}

// ToUint64 converts any valid value to an uint64.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
//
// Returns an error if the number is negative.
func ToUint64(T any) (uint64, error) {
	switch val := any(T).(type) {
	case int:
		if val < 0 {
			return 0, fmt.Errorf("value %d is negative", val)
		}
		return uint64(val), nil
	case int8:
		if val < 0 {
			return 0, fmt.Errorf("value %d is negative", val)
		}
		return uint64(val), nil
	case int16:
		if val < 0 {
			return 0, fmt.Errorf("value %d is negative", val)
		}
		return uint64(val), nil
	case int32:
		if val < 0 {
			return 0, fmt.Errorf("value %d is negative", val)
		}
		return uint64(val), nil
	case int64:
		if val < 0 {
			return 0, fmt.Errorf("value %d is negative", val)
		}
		return uint64(val), nil
	case uint:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint32:
		return uint64(val), nil
	case uint64:
		return val, nil
	case float32:
		if val < 0 {
			return 0, fmt.Errorf("value %f is negative", val)
		}
		return uint64(val), nil
	case float64:
		if val < 0 {
			return 0, fmt.Errorf("value %f is negative", val)
		}
		return uint64(val), nil
	case string:
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			v, err := strconv.ParseFloat(val, 64)
			if err == nil {
				return ToUint64(v) // as float
			}
		}
		return v, err
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", val)
	}
}

// ForceToUint64 converts any valid value to an uint64, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Float are truncated, booleans are converted to 1 (true) or 0
// (false).
func ForceToUint64(v any) uint64 {
	u, _ := ToUint64(v)
	return u
}

// ToBool converts any valid value to a bool.
//
// Valid values are booleans, numeric types (int*, uint*, float), and strings
// with values: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func ToBool(v any) (bool, error) {
	switch val := any(v).(type) {
	case bool:
		return val, nil
	case string:
		return strconv.ParseBool(val)
	case int, int8, int16, int32, int64:
		return ForceToInt64(val) != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		return ForceToUint64(val) != 0, nil
	case float32, float64:
		return ForceToFloat64(val) != 0, nil
	default:
		return false, fmt.Errorf("unsupported type: %T", v)
	}
}

// ForceToBool converts any valid value to a bool, ignoring any errors.
// It returns false if the conversion fails.
//
// Valid values are booleans, numeric types (int*, uint*, float), and strings
// with values: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func ForceToBool(v any) bool {
	b, _ := ToBool(v)
	return b
}

// ToFloat32 converts any valid value to a float32.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Booleans are converted to 1 (true) or 0 (false).
//
// Returns an error if the number overflows.
func ToFloat32(v any) (float32, error) {
	f64, err := ToFloat64(v)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}

// ForceToFloat32 converts any valid value to a float32, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Booleans are converted to 1 (true) or 0 (false).
func ForceToFloat32(v any) float32 {
	f, _ := ToFloat32(v)
	return f
}

// ToFloat64 converts any valid value to a float64.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Booleans are converted to 1 (true) or 0 (false).
//
// Returns an error if the number overflows.
func ToFloat64(v any) (float64, error) {
	switch val := any(v).(type) {
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case int, int8, int16, int32, int64:
		return float64(ForceToInt64(val)), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(ForceToUint64(val)), nil
	case string:
		return strconv.ParseFloat(val, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// ForceToFloat64 converts any valid value to a float64, ignoring any errors.
// It returns 0 if the conversion fails.
//
// Valid values are other numeric types (int*, uint*, float), numberic strings,
// and booleans. Booleans are converted to 1 (true) or 0 (false).
func ForceToFloat64(v any) float64 {
	f, _ := ToFloat64(v)
	return f
}
