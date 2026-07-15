package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestOptionalSome(t *testing.T) {
	o := dsx.Some(42)
	testx.True(t, o.Has())
	testx.Equal(t, 42, o.Get())
}

func TestOptionalNone(t *testing.T) {
	o := dsx.None[int]()
	testx.False(t, o.Has())
	testx.Equal(t, 0, o.Get())
}

func TestOptionalGetOr(t *testing.T) {
	testx.Equal(t, 42, dsx.Some(42).GetOr(-1))
	testx.Equal(t, -1, dsx.None[int]().GetOr(-1))
}

func TestOptionalGetOk(t *testing.T) {
	v, ok := dsx.Some(42).GetOk()
	testx.Equal(t, 42, v)
	testx.True(t, ok)

	v, ok = dsx.None[int]().GetOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestOptionalUnwrap(t *testing.T) {
	testx.Equal(t, 42, dsx.Some(42).Unwrap())
}

func TestOptionalUnwrapPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.None[int]().Unwrap()
}
