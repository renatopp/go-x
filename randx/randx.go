package randx

import (
	crand "crypto/rand"
	"math/rand/v2"
)

//-----------------------------------------------------------------------------
// INTEGERS
//-----------------------------------------------------------------------------

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int { return rand.Int() }

// Int8 returns a non-negative pseudo-random int8 from the default Source.
func Int8() int8 { return int8(rand.IntN(1 << 8)) }

// Int16 returns a non-negative pseudo-random int16 from the default Source.
func Int16() int16 { return int16(rand.IntN(1 << 16)) }

// Int32 returns a non-negative pseudo-random int32 from the default Source.
func Int32() int32 { return rand.Int32() }

// Int64 returns a non-negative pseudo-random int64 from the default Source.
func Int64() int64 { return rand.Int64() }

// Uint returns a non-negative pseudo-random uint from the default Source.
func Uint() uint { return rand.Uint() }

// Uint8 returns a non-negative pseudo-random uint8 from the default Source.
func Uint8() uint8 { return uint8(rand.UintN(1 << 8)) }

// Uint16 returns a non-negative pseudo-random uint16 from the default Source.
func Uint16() uint16 { return uint16(rand.UintN(1 << 16)) }

// Uint32 returns a non-negative pseudo-random uint32 from the default Source.
func Uint32() uint32 { return rand.Uint32() }

// Uint64 returns a non-negative pseudo-random uint64 from the default Source.
func Uint64() uint64 { return rand.Uint64() }

// IntN returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func IntN(n int) int { return rand.IntN(n) }

// Int8N returns, as an int8, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Int8N(n int8) int8 { return int8(rand.IntN(int(n))) }

// Int16N returns, as an int16, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Int16N(n int16) int16 { return int16(rand.IntN(int(n))) }

// Int32N returns, as an int32, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Int32N(n int32) int32 { return rand.Int32N(n) }

// Int64N returns, as an int64, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Int64N(n int64) int64 { return rand.Int64N(n) }

// UintN returns, as a uint, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func UintN(n uint) uint { return rand.UintN(n) }

// Uint8N returns, as a uint8, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Uint8N(n uint8) uint8 { return uint8(rand.UintN(uint(n))) }

// Uint16N returns, as a uint16, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Uint16N(n uint16) uint16 { return uint16(rand.UintN(uint(n))) }

// Uint32N returns, as a uint32, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Uint32N(n uint32) uint32 { return rand.Uint32N(n) }

// Uint64N returns, as a uint64, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
func Uint64N(n uint64) uint64 { return rand.Uint64N(n) }

// IntRange returns, as an int, a non-negative pseudo-random number in [min, max) from the default Source.
func IntRange(min, max int) int { return min + IntN(max-min) }

// IntRange8 returns, as an int8, a non-negative pseudo-random number in [min, max) from the default Source.
func IntRange8(min, max int8) int8 { return min + Int8N(max-min) }

// IntRange16 returns, as an int16, a non-negative pseudo-random number in [min, max) from the default Source.
func IntRange16(min, max int16) int16 { return min + Int16N(max-min) }

// IntRange32 returns, as an int32, a non-negative pseudo-random number in [min, max) from the default Source.
func IntRange32(min, max int32) int32 { return min + Int32N(max-min) }

// IntRange64 returns, as an int64, a non-negative pseudo-random number in [min, max) from the default Source.
func IntRange64(min, max int64) int64 { return min + Int64N(max-min) }

// UintRange returns, as a uint, a non-negative pseudo-random number in [min, max) from the default Source.
func UintRange(min, max uint) uint { return min + UintN(max-min) }

// UintRange8 returns, as a uint8, a non-negative pseudo-random number in [min, max) from the default Source.
func UintRange8(min, max uint8) uint8 { return min + Uint8N(max-min) }

// UintRange16 returns, as a uint16, a non-negative pseudo-random number in [min, max) from the default Source.
func UintRange16(min, max uint16) uint16 { return min + Uint16N(max-min) }

// UintRange32 returns, as a uint32, a non-negative pseudo-random number in [min, max) from the default Source.
func UintRange32(min, max uint32) uint32 { return min + Uint32N(max-min) }

// UintRange64 returns, as a uint64, a non-negative pseudo-random number in [min, max) from the default Source.
func UintRange64(min, max uint64) uint64 { return min + Uint64N(max-min) }

//-----------------------------------------------------------------------------
// FLOAT
//-----------------------------------------------------------------------------

// Float returns a pseudo-random float64 in [0.0, 1.0) from the default Source.
func Float() float64 { return rand.Float64() }

// Float32 returns a pseudo-random float32 in [0.0, 1.0) from the default Source.
func Float32() float32 { return rand.Float32() }

// Float64 returns a pseudo-random float64 in [0.0, 1.0) from the default Source.
func Float64() float64 { return rand.Float64() }

// FloatN returns a pseudo-random float64 in [0.0, n) from the default Source.
func FloatN(n float64) float64 { return rand.Float64() * n }

// Float32N returns a pseudo-random float32 in [0.0, n) from the default Source.
func Float32N(n float32) float32 { return rand.Float32() * n }

// Float64N returns a pseudo-random float64 in [0.0, n) from the default Source.
func Float64N(n float64) float64 { return rand.Float64() * n }

// FloatExp returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = FloatExp() / desiredRateParameter
func FloatExp() float64 { return rand.ExpFloat64() }

// FloatRange returns a pseudo-random float64 in [min, max) from the default Source.
func FloatRange(min, max float64) float64 { return min + rand.Float64()*(max-min) }

// Float64Range returns a pseudo-random float64 in [min, max) from the default Source.
func Float64Range(min, max float64) float64 { return FloatRange(min, max) }

// Float32Range returns a pseudo-random float32 in [min, max) from the default Source.
func Float32Range(min, max float32) float32 { return min + rand.Float32()*(max-min) }

//-----------------------------------------------------------------------------
// BOOL
//-----------------------------------------------------------------------------

// Bool returns a pseudo-random bool from the default Source.
func Bool() bool { return IntN(2) == 0 }

// Coin returns a pseudo-random bool from the default Source, with a 50% chance of being true or false. Alias to Bool.
func Coin() bool { return Bool() }

// Chance returns true with the given percentage chance (between 0 and 1) from the default Source.
func Chance(percent float64) bool { return Float() < percent }

//-----------------------------------------------------------------------------
// STRINGS
//-----------------------------------------------------------------------------

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const hexDigits = "0123456789abcdef"

// StringFrom returns a random string of length n from the given characters.
func StringFrom(chars string, n int) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = chars[IntN(len(chars))]
	}
	return string(result)
}

// String returns a random string of length n from the alphanumeric characters.
func String(n int) string { return StringFrom(alphanumeric, n) }

// StringHex returns a random string of length n from the hexadecimal characters.
func StringHex(n int) string { return StringFrom(hexDigits, n) }

// StringAlpha returns a random string of length n from the alphabetic characters.
func StringAlpha(n int) string { return StringFrom(alpha, n) }

// StringDigits returns a random string of length n from the digit characters.
func StringDigits(n int) string { return StringFrom(digits, n) }

//-----------------------------------------------------------------------------
// RUNES
//-----------------------------------------------------------------------------

// RuneFrom returns a random rune from the given characters.
func RuneFrom(chars string) rune { return rune(chars[IntN(len(chars))]) }

// Rune returns a random rune from the alphanumeric characters.
func Rune() rune { return RuneFrom(alphanumeric) }

// RuneHex returns a random rune from the hexadecimal characters.
func RuneHex() rune { return RuneFrom(hexDigits) }

// RuneAlpha returns a random rune from the alphabetic characters.
func RuneAlpha() rune { return RuneFrom(alpha) }

// RuneDigit returns a random rune from the digit characters.
func RuneDigit() rune { return RuneFrom(digits) }

// RunesFrom returns a random slice of runes of length n from the given characters.
func RunesFrom(chars string, n int) []rune {
	result := make([]rune, n)
	for i := range result {
		result[i] = RuneFrom(chars)
	}
	return result
}

// Runes returns a random slice of runes of length n from the alphanumeric characters.
func Runes(n int) []rune { return RunesFrom(alphanumeric, n) }

// RunesHex returns a random slice of runes of length n from the hexadecimal characters.
func RunesHex(n int) []rune { return RunesFrom(hexDigits, n) }

// RunesAlpha returns a random slice of runes of length n from the alphabetic characters.
func RunesAlpha(n int) []rune { return RunesFrom(alpha, n) }

// RunesDigits returns a random slice of runes of length n from the digit characters.
func RunesDigits(n int) []rune { return RunesFrom(digits, n) }

//-----------------------------------------------------------------------------
// BYTES
//-----------------------------------------------------------------------------

// Bytes returns a slice of n random bytes using crypto/rand for secure randomness.
func Bytes(n int) []byte {
	result := make([]byte, n)
	crand.Read(result)
	return result
}

// Byte returns a random byte using crypto/rand for secure randomness.
func Byte() byte { return Bytes(1)[0] }

// UUID returns a random UUID string in the format xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx,
// where x is a random hexadecimal digit and y is a random hexadecimal digit from 8 to b.
func UUID() string {
	return StringHex(8) + "-" + StringHex(4) + "-4" + StringHex(3) + "-" + StringHex(4) + "-" + StringHex(12)
}

//-----------------------------------------------------------------------------
// SLICES
//-----------------------------------------------------------------------------

// Pick is an alias for Choice, returning a random element from the given choices.
func Pick[T any](choices ...T) T { return choices[IntN(len(choices))] }

// PickSlice returns a random element from the given slice of choices.
func PickSlice[T any](choices []T) T { return choices[IntN(len(choices))] }

// Shuffle randomly shuffles the elements of the given slice in place.
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

//-----------------------------------------------------------------------------
// CONTEXTUAL
//-----------------------------------------------------------------------------

// NormalStandard returns a pseudo-random float64 drawn from the standard
// normal distribution (mean = 0, stddev = 1) from the default Source.
func NormalStandard() float64 {
	return rand.NormFloat64()
}

// Normal returns a pseudo-random float64 drawn from a normal distribution with
// the given mean and standard deviation from the default Source.
func Normal(mu, sigma float64) float64 {
	return rand.NormFloat64()*sigma + mu
}

// Dice simulates rolling a specified number of dice with a specified number of
// sides and returns the total.
func Dice(rolls, sides int) int {
	total := 0
	for range rolls {
		total += IntRange(1, sides)
	}
	return total
}
