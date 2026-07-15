package convx_test

import (
	"testing"

	"github.com/renatopp/go-x/convx"
	"github.com/renatopp/go-x/testx"
)

func TestToString(t *testing.T) {
	testx.Equal(t, "42", convx.ToString(42))
	testx.Equal(t, "3.14", convx.ToString(3.14))
	testx.Equal(t, "true", convx.ToString(true))
	testx.Equal(t, "hello", convx.ToString("hello"))
	testx.Equal(t, "<nil>", convx.ToString(nil))
	testx.Equal(t, "[1 2 3]", convx.ToString([]int{1, 2, 3}))
}

func TestToInt64(t *testing.T) {
	i, err := convx.ToInt64(42)
	testx.Nil(t, err)
	testx.Equal(t, int64(42), i)

	i, err = convx.ToInt64(int8(-8))
	testx.Nil(t, err)
	testx.Equal(t, int64(-8), i)

	i, err = convx.ToInt64(int16(-16))
	testx.Nil(t, err)
	testx.Equal(t, int64(-16), i)

	i, err = convx.ToInt64(int32(-32))
	testx.Nil(t, err)
	testx.Equal(t, int64(-32), i)

	i, err = convx.ToInt64(int64(-64))
	testx.Nil(t, err)
	testx.Equal(t, int64(-64), i)

	i, err = convx.ToInt64(uint(8))
	testx.Nil(t, err)
	testx.Equal(t, int64(8), i)

	i, err = convx.ToInt64(uint8(8))
	testx.Nil(t, err)
	testx.Equal(t, int64(8), i)

	i, err = convx.ToInt64(uint16(16))
	testx.Nil(t, err)
	testx.Equal(t, int64(16), i)

	i, err = convx.ToInt64(uint32(32))
	testx.Nil(t, err)
	testx.Equal(t, int64(32), i)

	i, err = convx.ToInt64(uint64(64))
	testx.Nil(t, err)
	testx.Equal(t, int64(64), i)

	i, err = convx.ToInt64(uint64(18446744073709551615))
	testx.NotNil(t, err)
	testx.Equal(t, int64(0), i)

	i, err = convx.ToInt64(float32(3.9))
	testx.Nil(t, err)
	testx.Equal(t, int64(3), i)

	i, err = convx.ToInt64(float64(-3.9))
	testx.Nil(t, err)
	testx.Equal(t, int64(-3), i)

	i, err = convx.ToInt64("42")
	testx.Nil(t, err)
	testx.Equal(t, int64(42), i)

	i, err = convx.ToInt64("-42")
	testx.Nil(t, err)
	testx.Equal(t, int64(-42), i)

	i, err = convx.ToInt64("not a number")
	testx.NotNil(t, err)
	testx.Equal(t, int64(0), i)

	i, err = convx.ToInt64("3.14")
	testx.NotNil(t, err)
	testx.Equal(t, int64(0), i)

	i, err = convx.ToInt64(true)
	testx.Nil(t, err)
	testx.Equal(t, int64(1), i)

	i, err = convx.ToInt64(false)
	testx.Nil(t, err)
	testx.Equal(t, int64(0), i)

	i, err = convx.ToInt64(struct{}{})
	testx.NotNil(t, err)
	testx.Equal(t, int64(0), i)
}

func TestForceToInt64(t *testing.T) {
	testx.Equal(t, int64(42), convx.ForceToInt64(42))
	testx.Equal(t, int64(0), convx.ForceToInt64("not a number"))
}

func TestToInt(t *testing.T) {
	i, err := convx.ToInt(42)
	testx.Nil(t, err)
	testx.Equal(t, 42, i)

	i, err = convx.ToInt(3.9)
	testx.Nil(t, err)
	testx.Equal(t, 3, i)

	i, err = convx.ToInt("42")
	testx.Nil(t, err)
	testx.Equal(t, 42, i)

	i, err = convx.ToInt(true)
	testx.Nil(t, err)
	testx.Equal(t, 1, i)

	i, err = convx.ToInt(uint64(18446744073709551615))
	testx.NotNil(t, err)
	testx.Equal(t, 0, i)
}

func TestForceToInt(t *testing.T) {
	testx.Equal(t, 42, convx.ForceToInt("42"))
	testx.Equal(t, 0, convx.ForceToInt("not a number"))
}

func TestToInt8(t *testing.T) {
	i, err := convx.ToInt8(100)
	testx.Nil(t, err)
	testx.Equal(t, int8(100), i)

	i, err = convx.ToInt8(-128)
	testx.Nil(t, err)
	testx.Equal(t, int8(-128), i)

	i, err = convx.ToInt8(127)
	testx.Nil(t, err)
	testx.Equal(t, int8(127), i)

	i, err = convx.ToInt8(128)
	testx.NotNil(t, err)
	testx.Equal(t, int8(0), i)

	i, err = convx.ToInt8(-129)
	testx.NotNil(t, err)
	testx.Equal(t, int8(0), i)

	i, err = convx.ToInt8("not a number")
	testx.NotNil(t, err)
	testx.Equal(t, int8(0), i)
}

func TestForceToInt8(t *testing.T) {
	testx.Equal(t, int8(100), convx.ForceToInt8(100))
	testx.Equal(t, int8(0), convx.ForceToInt8(200))
}

func TestToInt16(t *testing.T) {
	i, err := convx.ToInt16(1000)
	testx.Nil(t, err)
	testx.Equal(t, int16(1000), i)

	i, err = convx.ToInt16(-32768)
	testx.Nil(t, err)
	testx.Equal(t, int16(-32768), i)

	i, err = convx.ToInt16(32767)
	testx.Nil(t, err)
	testx.Equal(t, int16(32767), i)

	i, err = convx.ToInt16(32768)
	testx.NotNil(t, err)
	testx.Equal(t, int16(0), i)

	i, err = convx.ToInt16(-32769)
	testx.NotNil(t, err)
	testx.Equal(t, int16(0), i)
}

func TestForceToInt16(t *testing.T) {
	testx.Equal(t, int16(1000), convx.ForceToInt16(1000))
	testx.Equal(t, int16(0), convx.ForceToInt16(40000))
}

func TestToInt32(t *testing.T) {
	i, err := convx.ToInt32(100000)
	testx.Nil(t, err)
	testx.Equal(t, int32(100000), i)

	i, err = convx.ToInt32(-2147483648)
	testx.Nil(t, err)
	testx.Equal(t, int32(-2147483648), i)

	i, err = convx.ToInt32(2147483647)
	testx.Nil(t, err)
	testx.Equal(t, int32(2147483647), i)

	i, err = convx.ToInt32(int64(2147483648))
	testx.NotNil(t, err)
	testx.Equal(t, int32(0), i)

	i, err = convx.ToInt32(int64(-2147483649))
	testx.NotNil(t, err)
	testx.Equal(t, int32(0), i)
}

func TestForceToInt32(t *testing.T) {
	testx.Equal(t, int32(100000), convx.ForceToInt32(100000))
	testx.Equal(t, int32(0), convx.ForceToInt32(int64(9999999999)))
}

func TestToUint64(t *testing.T) {
	u, err := convx.ToUint64(42)
	testx.Nil(t, err)
	testx.Equal(t, uint64(42), u)

	u, err = convx.ToUint64(-1)
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(int8(-1))
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(int16(-1))
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(int32(-1))
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(int64(-1))
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(uint(8))
	testx.Nil(t, err)
	testx.Equal(t, uint64(8), u)

	u, err = convx.ToUint64(uint8(8))
	testx.Nil(t, err)
	testx.Equal(t, uint64(8), u)

	u, err = convx.ToUint64(uint16(16))
	testx.Nil(t, err)
	testx.Equal(t, uint64(16), u)

	u, err = convx.ToUint64(uint32(32))
	testx.Nil(t, err)
	testx.Equal(t, uint64(32), u)

	u, err = convx.ToUint64(uint64(64))
	testx.Nil(t, err)
	testx.Equal(t, uint64(64), u)

	u, err = convx.ToUint64(float32(3.9))
	testx.Nil(t, err)
	testx.Equal(t, uint64(3), u)

	u, err = convx.ToUint64(float64(-3.9))
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64("42")
	testx.Nil(t, err)
	testx.Equal(t, uint64(42), u)

	u, err = convx.ToUint64("3.14")
	testx.Nil(t, err)
	testx.Equal(t, uint64(3), u)

	u, err = convx.ToUint64("-3.14")
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64("not a number")
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(true)
	testx.Nil(t, err)
	testx.Equal(t, uint64(1), u)

	u, err = convx.ToUint64(false)
	testx.Nil(t, err)
	testx.Equal(t, uint64(0), u)

	u, err = convx.ToUint64(struct{}{})
	testx.NotNil(t, err)
	testx.Equal(t, uint64(0), u)
}

func TestForceToUint64(t *testing.T) {
	testx.Equal(t, uint64(42), convx.ForceToUint64(42))
	testx.Equal(t, uint64(0), convx.ForceToUint64(-1))
}

func TestToUint(t *testing.T) {
	u, err := convx.ToUint(42)
	testx.Nil(t, err)
	testx.Equal(t, uint(42), u)

	u, err = convx.ToUint(-1)
	testx.NotNil(t, err)
	testx.Equal(t, uint(0), u)
}

func TestForceToUint(t *testing.T) {
	testx.Equal(t, uint(42), convx.ForceToUint("42"))
	testx.Equal(t, uint(0), convx.ForceToUint("not a number"))
}

func TestToUint8(t *testing.T) {
	u, err := convx.ToUint8(200)
	testx.Nil(t, err)
	testx.Equal(t, uint8(200), u)

	u, err = convx.ToUint8(255)
	testx.Nil(t, err)
	testx.Equal(t, uint8(255), u)

	u, err = convx.ToUint8(256)
	testx.NotNil(t, err)
	testx.Equal(t, uint8(0), u)

	u, err = convx.ToUint8(-1)
	testx.NotNil(t, err)
	testx.Equal(t, uint8(0), u)
}

func TestForceToUint8(t *testing.T) {
	testx.Equal(t, uint8(200), convx.ForceToUint8(200))
	testx.Equal(t, uint8(0), convx.ForceToUint8(300))
}

func TestToUint16(t *testing.T) {
	u, err := convx.ToUint16(60000)
	testx.Nil(t, err)
	testx.Equal(t, uint16(60000), u)

	u, err = convx.ToUint16(65535)
	testx.Nil(t, err)
	testx.Equal(t, uint16(65535), u)

	u, err = convx.ToUint16(65536)
	testx.NotNil(t, err)
	testx.Equal(t, uint16(0), u)

	u, err = convx.ToUint16(-1)
	testx.NotNil(t, err)
	testx.Equal(t, uint16(0), u)
}

func TestForceToUint16(t *testing.T) {
	testx.Equal(t, uint16(60000), convx.ForceToUint16(60000))
	testx.Equal(t, uint16(0), convx.ForceToUint16(70000))
}

func TestToUint32(t *testing.T) {
	u, err := convx.ToUint32(100000)
	testx.Nil(t, err)
	testx.Equal(t, uint32(100000), u)

	u, err = convx.ToUint32(4294967295)
	testx.Nil(t, err)
	testx.Equal(t, uint32(4294967295), u)

	u, err = convx.ToUint32(int64(4294967296))
	testx.NotNil(t, err)
	testx.Equal(t, uint32(0), u)

	u, err = convx.ToUint32(-1)
	testx.NotNil(t, err)
	testx.Equal(t, uint32(0), u)
}

func TestForceToUint32(t *testing.T) {
	testx.Equal(t, uint32(100000), convx.ForceToUint32(100000))
	testx.Equal(t, uint32(0), convx.ForceToUint32(int64(9999999999)))
}

func TestToBool(t *testing.T) {
	b, err := convx.ToBool(true)
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool(false)
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool("true")
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool("True")
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool("TRUE")
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool("t")
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool("1")
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool("false")
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool("0")
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool("not a bool")
	testx.NotNil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool(42)
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool(int8(0))
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool(uint(42))
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool(uint8(0))
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool(3.14)
	testx.Nil(t, err)
	testx.True(t, b)

	b, err = convx.ToBool(float32(0))
	testx.Nil(t, err)
	testx.False(t, b)

	b, err = convx.ToBool(struct{}{})
	testx.NotNil(t, err)
	testx.False(t, b)
}

func TestForceToBool(t *testing.T) {
	testx.True(t, convx.ForceToBool("true"))
	testx.False(t, convx.ForceToBool("not a bool"))
}

func TestToFloat64(t *testing.T) {
	f, err := convx.ToFloat64(float32(3.5))
	testx.Nil(t, err)
	testx.AlmostEqual(t, 3.5, f, 0.0001)

	f, err = convx.ToFloat64(3.14)
	testx.Nil(t, err)
	testx.AlmostEqual(t, 3.14, f, 0.0001)

	f, err = convx.ToFloat64(42)
	testx.Nil(t, err)
	testx.AlmostEqual(t, 42.0, f, 0.0001)

	f, err = convx.ToFloat64(int8(-8))
	testx.Nil(t, err)
	testx.AlmostEqual(t, -8.0, f, 0.0001)

	f, err = convx.ToFloat64(uint(8))
	testx.Nil(t, err)
	testx.AlmostEqual(t, 8.0, f, 0.0001)

	f, err = convx.ToFloat64("3.14")
	testx.Nil(t, err)
	testx.AlmostEqual(t, 3.14, f, 0.0001)

	f, err = convx.ToFloat64("not a number")
	testx.NotNil(t, err)
	testx.AlmostEqual(t, 0.0, f, 0.0001)

	f, err = convx.ToFloat64(true)
	testx.Nil(t, err)
	testx.AlmostEqual(t, 1.0, f, 0.0001)

	f, err = convx.ToFloat64(false)
	testx.Nil(t, err)
	testx.AlmostEqual(t, 0.0, f, 0.0001)

	f, err = convx.ToFloat64(struct{}{})
	testx.NotNil(t, err)
	testx.AlmostEqual(t, 0.0, f, 0.0001)
}

func TestForceToFloat64(t *testing.T) {
	testx.AlmostEqual(t, 3.14, convx.ForceToFloat64("3.14"), 0.0001)
	testx.AlmostEqual(t, 0.0, convx.ForceToFloat64("not a number"), 0.0001)
}

func TestToFloat32(t *testing.T) {
	f, err := convx.ToFloat32(3.5)
	testx.Nil(t, err)
	testx.AlmostEqual(t, 3.5, float64(f), 0.0001)

	f, err = convx.ToFloat32("2.5")
	testx.Nil(t, err)
	testx.AlmostEqual(t, 2.5, float64(f), 0.0001)

	f, err = convx.ToFloat32("not a number")
	testx.NotNil(t, err)
	testx.AlmostEqual(t, 0.0, float64(f), 0.0001)
}

func TestForceToFloat32(t *testing.T) {
	testx.AlmostEqual(t, 2.5, float64(convx.ForceToFloat32("2.5")), 0.0001)
	testx.AlmostEqual(t, 0.0, float64(convx.ForceToFloat32("not a number")), 0.0001)
}
