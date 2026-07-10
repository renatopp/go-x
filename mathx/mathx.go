package mathx

import "math"

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type SignedNumber interface {
	SignedInteger | Float
}

type Number interface {
	Integer | Float
}

// Clamp clamps a value between a minimum and maximum value. If the value is
// less than the minimum, it returns the minimum. If the value is greater
// than the maximum, it returns the maximum. Otherwise, it returns the value
// itself.
func Clamp[T Number](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Clamp01 clamps a value between 0 and 1. If the value is less than 0, it
// returns 0. If the value is greater than 1, it returns 1. Otherwise, it
// returns the value itself.
func Clamp01[T Number](x T) T {
	return Clamp(x, 0, 1)
}

// Lerp performs linear interpolation between two values a and b using a
// parameter t. When t is 0, it returns a. When t is 1, it returns b. For values
// of t between 0 and 1, it returns a value that is proportionally between a and b.
func Lerp[T Number](a, b, t T) T {
	return a + t*(b-a)
}

// SmoothStep performs smooth interpolation between two values edge0 and edge1
// using a parameter x. When x is less than or equal to edge0, it returns 0. When
// x is greater than or equal to edge1, it returns 1. For values of x between
// edge0 and edge1, it returns a value that transitions smoothly from 0 to 1.
func SmoothStep[T Number](edge0, edge1, x T) T {
	if x <= edge0 {
		return 0
	}
	if x >= edge1 {
		return 1
	}
	t := (x - edge0) / (edge1 - edge0)
	return t * t * (3 - 2*t)
}

// Remap remaps a value from one range to another. It takes an input value x, the
// minimum and maximum of the input range (inMin and inMax), and the minimum and
// maximum of the output range (outMin and outMax). It returns the remapped value
// that corresponds to x in the output range.
func Remap[T Number](x, inMin, inMax, outMin, outMax T) T {
	return (x-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
}

// Divide performs division of two numbers x and y. If y is zero, it returns zero
// to avoid division by zero errors. Otherwise, it returns the result of x divided by y.
func Divide[T Number](x, y T) T {
	if y == 0 {
		return 0
	}
	return x / y
}

// AlmostEqual checks if two numbers a and b are approximately equal within a default
// epsilon value of 1e-9. It returns true if the absolute difference between a and b
// is less than or equal to epsilon, and false otherwise.
func AlmostEqual[T Number](a, b T) bool {
	return AlmostEqualBy(a, b, 1e-9)
}

// Abs returns the absolute value of a number x. If x is less than zero, it returns
// the negation of x. Otherwise, it returns x itself.
func AlmostEqualBy[T Number](a, b T, epsilon float64) bool {
	diff := Abs(float64(a) - float64(b))
	return diff <= epsilon
}

// Sign returns the sign of a number x. It returns 1 if x is greater than zero, -1 if
// x is less than zero, and 0 if x is equal to zero.
func Sign[T Number](x T) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// Gaussian returns the value of the Gaussian function for a given input x,
// mean, and standard deviation.
func Gaussian[T Number](x, mean, stddev T) float64 {
	expPart := Exp(-0.5 * Pow((float64(x)-float64(mean))/float64(stddev), 2))
	return (1 / (float64(stddev) * Sqrt(2*math.Pi))) * expPart
}

// Sigmoid returns the value of the sigmoid function for a given input x.
func Sigmoid[T Number](x T) float64 {
	return 1 / (1 + Exp(-x))
}
