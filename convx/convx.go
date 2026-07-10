package convx

import (
	"fmt"
	"strconv"
)

func ToString(v any) string {
	return fmt.Sprintf("%v", v)
}

func ToInt(v any) (int, error) {
	i64, err := ToInt64(v)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

func ForceToInt(v any) int {
	i, _ := ToInt(v)
	return i
}

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

func ForceToInt8(v any) int8 {
	i, _ := ToInt8(v)
	return i
}

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

func ForceToInt16(v any) int16 {
	i, _ := ToInt16(v)
	return i
}

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

func ForceToInt32(v any) int32 {
	i, _ := ToInt32(v)
	return i
}

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
		return strconv.ParseInt(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

func ForceToInt64(v any) int64 {
	i, _ := ToInt64(v)
	return i
}

func ToUint(v any) (uint, error) {
	u, err := ToUint64(v)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

func ForceToUint(v any) uint {
	u, _ := ToUint(v)
	return u
}

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

func ForceToUint8(v any) uint8 {
	u, _ := ToUint8(v)
	return u
}

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

func ForceToUint16(v any) uint16 {
	u, _ := ToUint16(v)
	return u
}

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

func ForceToUint32(v any) uint32 {
	u, _ := ToUint32(v)
	return u
}

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
		return strconv.ParseUint(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", val)
	}
}

func ForceToUint64(v any) uint64 {
	u, _ := ToUint64(v)
	return u
}

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

func ForceToBool(v any) bool {
	b, _ := ToBool(v)
	return b
}

func ToFloat32(v any) (float32, error) {
	f64, err := ToFloat64(v)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}

func ForceToFloat32(v any) float32 {
	f, _ := ToFloat32(v)
	return f
}

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

func ForceToFloat64(v any) float64 {
	f, _ := ToFloat64(v)
	return f
}
