package testx

import (
	"math"
	"testing"
)

func True(t *testing.T, condition bool) {
	if !condition {
		t.Fatalf("Expected condition to be true, but it was false.")
	}
}

func False(t *testing.T, condition bool) {
	if condition {
		t.Fatalf("Expected condition to be false, but it was true.")
	}
}

func AlmostEqual(t *testing.T, expected, actual float64, tolerance float64) {
	if math.Abs(expected-actual) > tolerance {
		t.Fatalf("Expected %v, but got %v.", expected, actual)
	}
}

func Equal(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Fatalf("Expected %v, but got %v.", expected, actual)
	}
}

func NotEqual(t *testing.T, expected, actual any) {
	if expected == actual {
		t.Fatalf("Expected %v to not equal %v.", expected, actual)
	}
}

func Nil(t *testing.T, value any) {
	if value != nil {
		t.Fatalf("Expected value to be nil, but it was %v.", value)
	}
}

func NotNil(t *testing.T, value any) {
	if value == nil {
		t.Fatalf("Expected value to not be nil, but it was.")
	}
}
