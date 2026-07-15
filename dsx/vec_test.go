package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestVec1(t *testing.T) {
	a := dsx.NewVec1(3)
	b := dsx.NewVec1(4)

	testx.AlmostEqual(t, 7, dsx.NewVec1(3).Add(b).X, 0.0001)
	testx.AlmostEqual(t, -1, a.Sub(b).X, 0.0001)
	testx.AlmostEqual(t, 6, a.Mul(2).X, 0.0001)
	testx.AlmostEqual(t, 1.5, a.Div(2).X, 0.0001)
	testx.AlmostEqual(t, -3, a.Neg().X, 0.0001)
	testx.AlmostEqual(t, 3, dsx.NewVec1(-3).Abs().X, 0.0001)
	testx.AlmostEqual(t, 12, a.Dot(b), 0.0001)
	testx.AlmostEqual(t, 3, a.Length(), 0.0001)
	testx.AlmostEqual(t, 1, a.Normalize().Length(), 0.0001)
	testx.AlmostEqual(t, 0, dsx.NewVec1(0).Normalize().X, 0.0001)
}

func TestVec2(t *testing.T) {
	a := dsx.NewVec2(3, 4)
	b := dsx.NewVec2(1, 2)

	sum := a.Add(b)
	testx.AlmostEqual(t, 4, sum.X, 0.0001)
	testx.AlmostEqual(t, 6, sum.Y, 0.0001)

	diff := a.Sub(b)
	testx.AlmostEqual(t, 2, diff.X, 0.0001)
	testx.AlmostEqual(t, 2, diff.Y, 0.0001)

	mul := a.Mul(2)
	testx.AlmostEqual(t, 6, mul.X, 0.0001)
	testx.AlmostEqual(t, 8, mul.Y, 0.0001)

	div := a.Div(2)
	testx.AlmostEqual(t, 1.5, div.X, 0.0001)
	testx.AlmostEqual(t, 2, div.Y, 0.0001)

	neg := a.Neg()
	testx.AlmostEqual(t, -3, neg.X, 0.0001)
	testx.AlmostEqual(t, -4, neg.Y, 0.0001)

	abs := dsx.NewVec2(-3, -4).Abs()
	testx.AlmostEqual(t, 3, abs.X, 0.0001)
	testx.AlmostEqual(t, 4, abs.Y, 0.0001)

	testx.AlmostEqual(t, 11, a.Dot(b), 0.0001)
	testx.AlmostEqual(t, 5, a.Length(), 0.0001)
	testx.AlmostEqual(t, 1, a.Normalize().Length(), 0.0001)

	zero := dsx.NewVec2(0, 0).Normalize()
	testx.AlmostEqual(t, 0, zero.X, 0.0001)
	testx.AlmostEqual(t, 0, zero.Y, 0.0001)
}

func TestVec3(t *testing.T) {
	a := dsx.NewVec3(1, 2, 2)
	b := dsx.NewVec3(3, 1, 1)

	sum := a.Add(b)
	testx.AlmostEqual(t, 4, sum.X, 0.0001)
	testx.AlmostEqual(t, 3, sum.Y, 0.0001)
	testx.AlmostEqual(t, 3, sum.Z, 0.0001)

	diff := a.Sub(b)
	testx.AlmostEqual(t, -2, diff.X, 0.0001)
	testx.AlmostEqual(t, 1, diff.Y, 0.0001)
	testx.AlmostEqual(t, 1, diff.Z, 0.0001)

	mul := a.Mul(2)
	testx.AlmostEqual(t, 2, mul.X, 0.0001)
	testx.AlmostEqual(t, 4, mul.Y, 0.0001)
	testx.AlmostEqual(t, 4, mul.Z, 0.0001)

	div := a.Div(2)
	testx.AlmostEqual(t, 0.5, div.X, 0.0001)
	testx.AlmostEqual(t, 1, div.Y, 0.0001)
	testx.AlmostEqual(t, 1, div.Z, 0.0001)

	neg := a.Neg()
	testx.AlmostEqual(t, -1, neg.X, 0.0001)
	testx.AlmostEqual(t, -2, neg.Y, 0.0001)
	testx.AlmostEqual(t, -2, neg.Z, 0.0001)

	abs := dsx.NewVec3(-1, -2, -2).Abs()
	testx.AlmostEqual(t, 1, abs.X, 0.0001)
	testx.AlmostEqual(t, 2, abs.Y, 0.0001)
	testx.AlmostEqual(t, 2, abs.Z, 0.0001)

	testx.AlmostEqual(t, 7, a.Dot(b), 0.0001)
	testx.AlmostEqual(t, 3, a.Length(), 0.0001)
	testx.AlmostEqual(t, 1, a.Normalize().Length(), 0.0001)

	zero := dsx.NewVec3(0, 0, 0).Normalize()
	testx.AlmostEqual(t, 0, zero.X, 0.0001)
	testx.AlmostEqual(t, 0, zero.Y, 0.0001)
	testx.AlmostEqual(t, 0, zero.Z, 0.0001)
}
