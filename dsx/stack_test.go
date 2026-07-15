package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestStackPush(t *testing.T) {
	s := dsx.NewStack[int]()
	s.Push(1, 2, 3)
	testx.Equal(t, 3, s.Size())
	equalSlice(t, []int{1, 2, 3}, s.ToSlice())
}

func TestStackPushSlice(t *testing.T) {
	s := dsx.NewStack[int]()
	s.PushSlice([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, s.ToSlice())
}

func TestStackGet(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 1, s.Get(0))
	testx.Equal(t, 3, s.Get(2))
	testx.Equal(t, 3, s.Get(-1))
	testx.Equal(t, 1, s.Get(-3))
}

func TestStackGetPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewStackFrom([]int{1, 2, 3}).Get(3)
}

func TestStackGetOr(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 1, s.GetOr(0, -1))
	testx.Equal(t, 3, s.GetOr(-1, -1))
	testx.Equal(t, -1, s.GetOr(3, -1))
	testx.Equal(t, -1, s.GetOr(-4, -1))
}

func TestStackGetOk(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	v, ok := s.GetOk(1)
	testx.Equal(t, 2, v)
	testx.True(t, ok)

	v, ok = s.GetOk(3)
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestStackFirst(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 3, s.First())
	testx.Equal(t, 3, s.FirstOr(-1))
	v, ok := s.FirstOk()
	testx.Equal(t, 3, v)
	testx.True(t, ok)

	empty := dsx.NewStack[int]()
	testx.Equal(t, -1, empty.FirstOr(-1))
	v, ok = empty.FirstOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestStackFirstPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewStack[int]().First()
}

func TestStackLast(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 1, s.Last())
	testx.Equal(t, 1, s.LastOr(-1))
	v, ok := s.LastOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	empty := dsx.NewStack[int]()
	testx.Equal(t, -1, empty.LastOr(-1))
	v, ok = empty.LastOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestStackLastPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewStack[int]().Last()
}

func TestStackPop(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 3, s.Pop())
	testx.Equal(t, 2, s.Size())
	testx.Equal(t, 2, s.PopOr(-1))
	v, ok := s.PopOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)
	testx.Equal(t, 0, s.Size())

	testx.Equal(t, -1, s.PopOr(-1))
	v, ok = s.PopOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestStackPopPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewStack[int]().Pop()
}

func TestStackIndexOf(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 1, s.IndexOf(2))
	testx.Equal(t, -1, s.IndexOf(42))
}

func TestStackIndexOfFunc(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.Equal(t, 2, s.IndexOfFunc(func(v int) bool { return v == 3 }))
	testx.Equal(t, -1, s.IndexOfFunc(func(v int) bool { return v == 42 }))
}

func TestStackContains(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.True(t, s.Contains(2))
	testx.False(t, s.Contains(42))
}

func TestStackContainsFunc(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	testx.True(t, s.ContainsFunc(func(v int) bool { return v == 3 }))
	testx.False(t, s.ContainsFunc(func(v int) bool { return v == 42 }))
}

func TestStackSize(t *testing.T) {
	testx.Equal(t, 0, dsx.NewStack[int]().Size())
	testx.Equal(t, 3, dsx.NewStackFrom([]int{1, 2, 3}).Size())
}

func TestStackClear(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	s.Clear()
	testx.Equal(t, 0, s.Size())
}

func TestStackClone(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	clone := s.Clone()
	clone.Push(4)
	testx.Equal(t, 3, s.Size())
	testx.Equal(t, 4, clone.Size())
}

func TestStackToSlice(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, s.ToSlice())
}

func TestStackIter(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	values := []int{}
	for _, v := range s.Iter() {
		values = append(values, v)
	}
	equalSlice(t, []int{3, 2, 1}, values)
}

func TestStackForEach(t *testing.T) {
	s := dsx.NewStackFrom([]int{1, 2, 3})
	values := []int{}
	s.ForEach(func(i, v int) {
		values = append(values, v)
	})
	equalSlice(t, []int{3, 2, 1}, values)
}
