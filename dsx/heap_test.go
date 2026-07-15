package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestHeapPush(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 10)
	h.PushSlice(2, []int{20, 21})
	testx.Equal(t, 3, h.Size())
}

func TestHeapGet(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	h.Push(3, 3)
	h.Push(9, 9)
	h.Push(2, 2)

	slice := h.ToSlice()
	testx.Equal(t, slice[0], h.Get(0))
	testx.Equal(t, slice[len(slice)-1], h.Get(-1))
	testx.Equal(t, slice[0], h.Get(-len(slice)))
}

func TestHeapGetPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	h.Get(5)
}

func TestHeapGetOr(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	testx.Equal(t, 1, h.GetOr(0, -1))
	testx.Equal(t, -1, h.GetOr(5, -1))
	testx.Equal(t, -1, h.GetOr(-5, -1))
}

func TestHeapGetOk(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	v, ok := h.GetOk(0)
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	v, ok = h.GetOk(5)
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestHeapFirst(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	h.Push(3, 3)
	testx.Equal(t, 1, h.First())
	testx.Equal(t, 1, h.FirstOr(-1))
	v, ok := h.FirstOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	empty := dsx.NewHeap[int]()
	testx.Equal(t, -1, empty.FirstOr(-1))
	v, ok = empty.FirstOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestHeapFirstPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewHeap[int]().First()
}

func TestHeapLast(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	h.Push(2, 2)
	slice := h.ToSlice()
	testx.Equal(t, slice[len(slice)-1], h.Last())
	testx.Equal(t, slice[len(slice)-1], h.LastOr(-1))
	v, ok := h.LastOk()
	testx.Equal(t, slice[len(slice)-1], v)
	testx.True(t, ok)

	empty := dsx.NewHeap[int]()
	testx.Equal(t, -1, empty.LastOr(-1))
	v, ok = empty.LastOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestHeapLastPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewHeap[int]().Last()
}

func TestHeapPop(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	h.Push(3, 3)
	h.Push(9, 9)
	h.Push(2, 2)

	popped := []int{}
	for h.Size() > 0 {
		popped = append(popped, h.Pop())
	}
	equalSlice(t, []int{1, 2, 3, 5, 9}, popped)
}

func TestHeapPopOr(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	testx.Equal(t, 1, h.PopOr(-1))
	testx.Equal(t, -1, h.PopOr(-1))
}

func TestHeapPopOk(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	v, ok := h.PopOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	v, ok = h.PopOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestHeapPopPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewHeap[int]().Pop()
}

func TestHeapIndexOf(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	h.Push(3, 3)
	testx.NotEqual(t, -1, h.IndexOf(3))
	testx.Equal(t, -1, h.IndexOf(42))
}

func TestHeapIndexOfFunc(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	testx.NotEqual(t, -1, h.IndexOfFunc(func(v int) bool { return v == 5 }))
	testx.Equal(t, -1, h.IndexOfFunc(func(v int) bool { return v == 42 }))
}

func TestHeapContains(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	testx.True(t, h.Contains(5))
	testx.False(t, h.Contains(42))
}

func TestHeapContainsFunc(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(5, 5)
	h.Push(1, 1)
	testx.True(t, h.ContainsFunc(func(v int) bool { return v == 1 }))
	testx.False(t, h.ContainsFunc(func(v int) bool { return v == 42 }))
}

func TestHeapSize(t *testing.T) {
	testx.Equal(t, 0, dsx.NewHeap[int]().Size())
	h := dsx.NewHeap[int]()
	h.Push(1, 1, 2, 3)
	testx.Equal(t, 3, h.Size())
}

func TestHeapClear(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1, 2, 3)
	h.Clear()
	testx.Equal(t, 0, h.Size())
}

func TestHeapClone(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1, 2)
	clone := h.Clone()
	clone.Push(1, 3)
	testx.Equal(t, 2, h.Size())
	testx.Equal(t, 3, clone.Size())
}

func TestHeapConcat(t *testing.T) {
	h1 := dsx.NewHeap[int]()
	h1.Push(1, 1)
	h1.Push(3, 3)
	h2 := dsx.NewHeap[int]()
	h2.Push(2, 2)
	h2.Push(4, 4)

	merged := h1.Concat(h2)
	testx.Equal(t, 2, h1.Size())
	testx.Equal(t, 2, h2.Size())
	testx.Equal(t, 4, merged.Size())

	popped := []int{}
	for merged.Size() > 0 {
		popped = append(popped, merged.Pop())
	}
	equalSlice(t, []int{1, 2, 3, 4}, popped)
}

func TestHeapToSlice(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1, 2, 3)
	testx.Equal(t, 3, len(h.ToSlice()))
}

func TestHeapIter(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	h.Push(2, 2)
	slice := h.ToSlice()
	values := []int{}
	for i, v := range h.Iter() {
		testx.Equal(t, slice[i], v)
		values = append(values, v)
	}
	testx.Equal(t, len(slice), len(values))
}

func TestHeapForEach(t *testing.T) {
	h := dsx.NewHeap[int]()
	h.Push(1, 1)
	h.Push(2, 2)
	slice := h.ToSlice()
	values := []int{}
	h.ForEach(func(i, v int) {
		testx.Equal(t, slice[i], v)
		values = append(values, v)
	})
	testx.Equal(t, len(slice), len(values))
}
