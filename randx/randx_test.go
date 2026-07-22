package randx

import (
	"regexp"
	"strings"
	"testing"

	"github.com/renatopp/go-x/testx"
)

const iterations = 1000

//-----------------------------------------------------------------------------
// INTEGERS
//-----------------------------------------------------------------------------

func TestIntN(t *testing.T) {
	for range iterations {
		v := IntN(10)
		testx.True(t, v >= 0 && v < 10)
	}
}

func TestInt8N(t *testing.T) {
	for range iterations {
		v := Int8N(10)
		testx.True(t, v >= 0 && v < 10)
	}
}

func TestInt16N(t *testing.T) {
	for range iterations {
		v := Int16N(10)
		testx.True(t, v >= 0 && v < 10)
	}
}

func TestInt32N(t *testing.T) {
	for range iterations {
		v := Int32N(10)
		testx.True(t, v >= 0 && v < 10)
	}
}

func TestInt64N(t *testing.T) {
	for range iterations {
		v := Int64N(10)
		testx.True(t, v >= 0 && v < 10)
	}
}

func TestUintN(t *testing.T) {
	for range iterations {
		v := UintN(10)
		testx.True(t, v < 10)
	}
}

func TestUint8N(t *testing.T) {
	for range iterations {
		v := Uint8N(10)
		testx.True(t, v < 10)
	}
}

func TestUint16N(t *testing.T) {
	for range iterations {
		v := Uint16N(10)
		testx.True(t, v < 10)
	}
}

func TestUint32N(t *testing.T) {
	for range iterations {
		v := Uint32N(10)
		testx.True(t, v < 10)
	}
}

func TestUint64N(t *testing.T) {
	for range iterations {
		v := Uint64N(10)
		testx.True(t, v < 10)
	}
}

func TestIntRange(t *testing.T) {
	for range iterations {
		v := IntRange(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestIntRange8(t *testing.T) {
	for range iterations {
		v := IntRange8(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestIntRange16(t *testing.T) {
	for range iterations {
		v := IntRange16(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestIntRange32(t *testing.T) {
	for range iterations {
		v := IntRange32(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestIntRange64(t *testing.T) {
	for range iterations {
		v := IntRange64(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestUintRange(t *testing.T) {
	for range iterations {
		v := UintRange(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestUintRange8(t *testing.T) {
	for range iterations {
		v := UintRange8(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestUintRange16(t *testing.T) {
	for range iterations {
		v := UintRange16(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestUintRange32(t *testing.T) {
	for range iterations {
		v := UintRange32(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestUintRange64(t *testing.T) {
	for range iterations {
		v := UintRange64(5, 15)
		testx.True(t, v >= 5 && v < 15)
	}
}

func TestIntNPanicsOnNonPositive(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	IntN(0)
}

//-----------------------------------------------------------------------------
// FLOAT
//-----------------------------------------------------------------------------

func TestFloat(t *testing.T) {
	for range iterations {
		v := Float()
		testx.True(t, v >= 0.0 && v < 1.0)
	}
}

func TestFloat32(t *testing.T) {
	for range iterations {
		v := Float32()
		testx.True(t, v >= 0.0 && v < 1.0)
	}
}

func TestFloat64(t *testing.T) {
	for range iterations {
		v := Float64()
		testx.True(t, v >= 0.0 && v < 1.0)
	}
}

func TestFloatN(t *testing.T) {
	for range iterations {
		v := FloatN(5.0)
		testx.True(t, v >= 0.0 && v < 5.0)
	}
}

func TestFloat32N(t *testing.T) {
	for range iterations {
		v := Float32N(5.0)
		testx.True(t, v >= 0.0 && v < 5.0)
	}
}

func TestFloat64N(t *testing.T) {
	for range iterations {
		v := Float64N(5.0)
		testx.True(t, v >= 0.0 && v < 5.0)
	}
}

func TestFloatExp(t *testing.T) {
	for range iterations {
		v := FloatExp()
		testx.True(t, v > 0.0)
	}
}

func TestFloatRange(t *testing.T) {
	for range iterations {
		v := FloatRange(5.0, 10.0)
		testx.True(t, v >= 5.0 && v < 10.0)
	}
}

func TestFloat64Range(t *testing.T) {
	for range iterations {
		v := Float64Range(5.0, 10.0)
		testx.True(t, v >= 5.0 && v < 10.0)
	}
}

func TestFloat32Range(t *testing.T) {
	for range iterations {
		v := Float32Range(5.0, 10.0)
		testx.True(t, v >= 5.0 && v < 10.0)
	}
}

//-----------------------------------------------------------------------------
// BOOL
//-----------------------------------------------------------------------------

func TestBool(t *testing.T) {
	seenTrue, seenFalse := false, false
	for range iterations {
		if Bool() {
			seenTrue = true
		} else {
			seenFalse = true
		}
	}
	testx.True(t, seenTrue)
	testx.True(t, seenFalse)
}

func TestCoin(t *testing.T) {
	seenTrue, seenFalse := false, false
	for range iterations {
		if Coin() {
			seenTrue = true
		} else {
			seenFalse = true
		}
	}
	testx.True(t, seenTrue)
	testx.True(t, seenFalse)
}

func TestChance(t *testing.T) {
	for range iterations {
		testx.False(t, Chance(0.0))
	}
	for range iterations {
		testx.True(t, Chance(1.0))
	}
}

//-----------------------------------------------------------------------------
// STRINGS
//-----------------------------------------------------------------------------

func TestStringFrom(t *testing.T) {
	s := StringFrom("ab", 20)
	testx.Equal(t, 20, len(s))
	for _, c := range s {
		testx.True(t, c == 'a' || c == 'b')
	}
}

func TestString(t *testing.T) {
	s := String(16)
	testx.Equal(t, 16, len(s))
	testx.True(t, isSubsetOf(s, alphanumeric))
}

func TestStringHex(t *testing.T) {
	s := StringHex(16)
	testx.Equal(t, 16, len(s))
	testx.True(t, isSubsetOf(s, hexDigits))
}

func TestStringAlpha(t *testing.T) {
	s := StringAlpha(16)
	testx.Equal(t, 16, len(s))
	testx.True(t, isSubsetOf(s, alpha))
}

func TestStringDigits(t *testing.T) {
	s := StringDigits(16)
	testx.Equal(t, 16, len(s))
	testx.True(t, isSubsetOf(s, digits))
}

//-----------------------------------------------------------------------------
// RUNES
//-----------------------------------------------------------------------------

func TestRuneFrom(t *testing.T) {
	for range iterations {
		r := RuneFrom("ab")
		testx.True(t, r == 'a' || r == 'b')
	}
}

func TestRune(t *testing.T) {
	for range iterations {
		r := Rune()
		testx.True(t, strings.ContainsRune(alphanumeric, r))
	}
}

func TestRuneHex(t *testing.T) {
	for range iterations {
		r := RuneHex()
		testx.True(t, strings.ContainsRune(hexDigits, r))
	}
}

func TestRuneAlpha(t *testing.T) {
	for range iterations {
		r := RuneAlpha()
		testx.True(t, strings.ContainsRune(alpha, r))
	}
}

func TestRuneDigit(t *testing.T) {
	for range iterations {
		r := RuneDigit()
		testx.True(t, strings.ContainsRune(digits, r))
	}
}

func TestRunesFrom(t *testing.T) {
	rs := RunesFrom("ab", 20)
	testx.Equal(t, 20, len(rs))
	for _, r := range rs {
		testx.True(t, r == 'a' || r == 'b')
	}
}

func TestRunes(t *testing.T) {
	rs := Runes(16)
	testx.Equal(t, 16, len(rs))
	for _, r := range rs {
		testx.True(t, strings.ContainsRune(alphanumeric, r))
	}
}

func TestRunesHex(t *testing.T) {
	rs := RunesHex(16)
	testx.Equal(t, 16, len(rs))
	for _, r := range rs {
		testx.True(t, strings.ContainsRune(hexDigits, r))
	}
}

func TestRunesAlpha(t *testing.T) {
	rs := RunesAlpha(16)
	testx.Equal(t, 16, len(rs))
	for _, r := range rs {
		testx.True(t, strings.ContainsRune(alpha, r))
	}
}

func TestRunesDigits(t *testing.T) {
	rs := RunesDigits(16)
	testx.Equal(t, 16, len(rs))
	for _, r := range rs {
		testx.True(t, strings.ContainsRune(digits, r))
	}
}

//-----------------------------------------------------------------------------
// BYTES
//-----------------------------------------------------------------------------

func TestBytes(t *testing.T) {
	b := Bytes(16)
	testx.Equal(t, 16, len(b))
}

func TestByte(t *testing.T) {
	_ = Byte() // any byte value is valid; just ensure it doesn't panic
}

var uuidPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func TestUUID(t *testing.T) {
	for range iterations {
		id := UUID()
		testx.True(t, uuidPattern.MatchString(id))
	}
}

//-----------------------------------------------------------------------------
// SLICES
//-----------------------------------------------------------------------------

func TestPick(t *testing.T) {
	choices := []int{1, 2, 3}
	for range iterations {
		v := Pick(choices...)
		testx.True(t, v == 1 || v == 2 || v == 3)
	}
}

func TestPickSlice(t *testing.T) {
	choices := []int{1, 2, 3}
	for range iterations {
		v := PickSlice(choices)
		testx.True(t, v == 1 || v == 2 || v == 3)
	}
}

func TestShuffle(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := make([]int, len(original))
	copy(slice, original)

	Shuffle(slice)

	testx.Equal(t, len(original), len(slice))
	counts := map[int]int{}
	for _, v := range slice {
		counts[v]++
	}
	for _, v := range original {
		testx.Equal(t, 1, counts[v])
	}
}

//-----------------------------------------------------------------------------
// CONTEXTUAL
//-----------------------------------------------------------------------------

func TestNormalStandard(t *testing.T) {
	sum := 0.0
	for range iterations {
		sum += NormalStandard()
	}
	mean := sum / float64(iterations)
	testx.AlmostEqual(t, 0.0, mean, 0.3)
}

func TestNormal(t *testing.T) {
	sum := 0.0
	for range iterations {
		sum += Normal(10.0, 1.0)
	}
	mean := sum / float64(iterations)
	testx.AlmostEqual(t, 10.0, mean, 0.3)
}

func TestDice(t *testing.T) {
	for range iterations {
		v := Dice(3, 6)
		testx.True(t, v >= 3 && v <= 18)
	}
}

//-----------------------------------------------------------------------------
// HELPERS
//-----------------------------------------------------------------------------

func isSubsetOf(s, chars string) bool {
	for _, c := range s {
		if !strings.ContainsRune(chars, c) {
			return false
		}
	}
	return true
}
