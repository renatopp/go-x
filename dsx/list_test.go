package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestListPush(t *testing.T) {
	l := dsx.NewList[int]()
	l.Push(1, 2, 3)
	testx.Equal(t, 3, l.Size())
	equalSlice(t, []int{1, 2, 3}, l.ToSlice())
}

func TestListPushSlice(t *testing.T) {
	l := dsx.NewList[int]()
	l.PushSlice([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, l.ToSlice())
}

func TestListGet(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 1, l.Get(0))
	testx.Equal(t, 3, l.Get(2))
	testx.Equal(t, 3, l.Get(-1))
	testx.Equal(t, 1, l.Get(-3))
}

func TestListGetPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewListFrom([]int{1, 2, 3}).Get(3)
}

func TestListGetOr(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 1, l.GetOr(0, -1))
	testx.Equal(t, 3, l.GetOr(-1, -1))
	testx.Equal(t, -1, l.GetOr(3, -1))
	testx.Equal(t, -1, l.GetOr(-4, -1))
}

func TestListGetOk(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	v, ok := l.GetOk(1)
	testx.Equal(t, 2, v)
	testx.True(t, ok)

	v, ok = l.GetOk(3)
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestListSet(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	l.Set(1, 20)
	testx.Equal(t, 20, l.Get(1))
	l.Set(-1, 30)
	testx.Equal(t, 30, l.Get(2))
}

func TestListSetPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewListFrom([]int{1, 2, 3}).Set(3, 99)
}

func TestListInsert(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	l.Insert(1, 10, 11)
	equalSlice(t, []int{1, 10, 11, 2, 3}, l.ToSlice())

	l2 := dsx.NewListFrom([]int{1, 2, 3})
	l2.Insert(0, 0)
	equalSlice(t, []int{0, 1, 2, 3}, l2.ToSlice())

	l3 := dsx.NewListFrom([]int{1, 2, 3})
	l3.Insert(3, 4)
	equalSlice(t, []int{1, 2, 3, 4}, l3.ToSlice())
}

func TestListInsertClampsOutOfRange(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	l.Insert(100, 4)
	equalSlice(t, []int{1, 2, 3, 4}, l.ToSlice())

	l2 := dsx.NewListFrom([]int{1, 2, 3})
	l2.Insert(-100, 0)
	equalSlice(t, []int{0, 1, 2, 3}, l2.ToSlice())
}

func TestListRemoveAt(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 2, l.RemoveAt(1))
	equalSlice(t, []int{1, 3}, l.ToSlice())

	l2 := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 3, l2.RemoveAt(-1))
	equalSlice(t, []int{1, 2}, l2.ToSlice())
}

func TestListRemoveAtPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewListFrom([]int{1, 2, 3}).RemoveAt(3)
}

func TestListRemoveAtOk(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	v, ok := l.RemoveAtOk(1)
	testx.Equal(t, 2, v)
	testx.True(t, ok)
	equalSlice(t, []int{1, 3}, l.ToSlice())

	v, ok = l.RemoveAtOk(5)
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestListRemove(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3, 2})
	testx.True(t, l.Remove(2))
	equalSlice(t, []int{1, 3, 2}, l.ToSlice())
	testx.False(t, l.Remove(42))
}

func TestListRemoveFunc(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3, 4, 5})
	n := l.RemoveFunc(func(v int) bool { return v%2 == 0 })
	testx.Equal(t, 2, n)
	equalSlice(t, []int{1, 3, 5}, l.ToSlice())
}

func TestListFirst(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 1, l.First())
	testx.Equal(t, 1, l.FirstOr(-1))
	v, ok := l.FirstOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	empty := dsx.NewList[int]()
	testx.Equal(t, -1, empty.FirstOr(-1))
	v, ok = empty.FirstOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestListFirstPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewList[int]().First()
}

func TestListLast(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 3, l.Last())
	testx.Equal(t, 3, l.LastOr(-1))
	v, ok := l.LastOk()
	testx.Equal(t, 3, v)
	testx.True(t, ok)

	empty := dsx.NewList[int]()
	testx.Equal(t, -1, empty.LastOr(-1))
	v, ok = empty.LastOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestListLastPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewList[int]().Last()
}

func TestListPop(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 3, l.Pop())
	testx.Equal(t, 2, l.Size())
	testx.Equal(t, 2, l.PopOr(-1))
	v, ok := l.PopOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)
	testx.Equal(t, 0, l.Size())

	testx.Equal(t, -1, l.PopOr(-1))
	v, ok = l.PopOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestListPopPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewList[int]().Pop()
}

func TestListIndexOf(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 1, l.IndexOf(2))
	testx.Equal(t, -1, l.IndexOf(42))
}

func TestListIndexOfFunc(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.Equal(t, 2, l.IndexOfFunc(func(v int) bool { return v == 3 }))
	testx.Equal(t, -1, l.IndexOfFunc(func(v int) bool { return v == 42 }))
}

func TestListContains(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.True(t, l.Contains(2))
	testx.False(t, l.Contains(42))
}

func TestListContainsFunc(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	testx.True(t, l.ContainsFunc(func(v int) bool { return v == 3 }))
	testx.False(t, l.ContainsFunc(func(v int) bool { return v == 42 }))
}

func TestListSize(t *testing.T) {
	testx.Equal(t, 0, dsx.NewList[int]().Size())
	testx.Equal(t, 3, dsx.NewListFrom([]int{1, 2, 3}).Size())
}

func TestListClear(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	l.Clear()
	testx.Equal(t, 0, l.Size())
}

func TestListClone(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	clone := l.Clone()
	clone.Push(4)
	testx.Equal(t, 3, l.Size())
	testx.Equal(t, 4, clone.Size())
}

func TestListConcat(t *testing.T) {
	l1 := dsx.NewListFrom([]int{1, 2})
	l2 := dsx.NewListFrom([]int{3, 4})
	l3 := dsx.NewListFrom([]int{5})

	merged := l1.Concat(l2, l3)
	equalSlice(t, []int{1, 2}, l1.ToSlice())
	equalSlice(t, []int{1, 2, 3, 4, 5}, merged.ToSlice())
}

func TestListToSlice(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, l.ToSlice())
}

func TestListIter(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	values := []int{}
	for _, v := range l.Iter() {
		values = append(values, v)
	}
	equalSlice(t, []int{1, 2, 3}, values)
}

func TestListForEach(t *testing.T) {
	l := dsx.NewListFrom([]int{1, 2, 3})
	values := []int{}
	l.ForEach(func(i, v int) {
		values = append(values, v)
	})
	equalSlice(t, []int{1, 2, 3}, values)
}
