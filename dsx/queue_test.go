package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestQueuePush(t *testing.T) {
	q := dsx.NewQueue[int]()
	q.Push(1, 2, 3)
	testx.Equal(t, 3, q.Size())
	equalSlice(t, []int{1, 2, 3}, q.ToSlice())
}

func TestQueuePushSlice(t *testing.T) {
	q := dsx.NewQueue[int]()
	q.PushSlice([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, q.ToSlice())
}

func TestQueueGet(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 1, q.Get(0))
	testx.Equal(t, 3, q.Get(2))
	testx.Equal(t, 3, q.Get(-1))
	testx.Equal(t, 1, q.Get(-3))
}

func TestQueueGetPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewQueueFrom([]int{1, 2, 3}).Get(3)
}

func TestQueueGetOr(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 1, q.GetOr(0, -1))
	testx.Equal(t, 3, q.GetOr(-1, -1))
	testx.Equal(t, -1, q.GetOr(3, -1))
	testx.Equal(t, -1, q.GetOr(-4, -1))
}

func TestQueueGetOk(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	v, ok := q.GetOk(1)
	testx.Equal(t, 2, v)
	testx.True(t, ok)

	v, ok = q.GetOk(3)
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestQueueFirst(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 1, q.First())
	testx.Equal(t, 1, q.FirstOr(-1))
	v, ok := q.FirstOk()
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	empty := dsx.NewQueue[int]()
	testx.Equal(t, -1, empty.FirstOr(-1))
	v, ok = empty.FirstOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestQueueFirstPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewQueue[int]().First()
}

func TestQueueLast(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 3, q.Last())
	testx.Equal(t, 3, q.LastOr(-1))
	v, ok := q.LastOk()
	testx.Equal(t, 3, v)
	testx.True(t, ok)

	empty := dsx.NewQueue[int]()
	testx.Equal(t, -1, empty.LastOr(-1))
	v, ok = empty.LastOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestQueueLastPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewQueue[int]().Last()
}

func TestQueuePop(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 1, q.Pop())
	testx.Equal(t, 2, q.Size())
	testx.Equal(t, 2, q.PopOr(-1))
	v, ok := q.PopOk()
	testx.Equal(t, 3, v)
	testx.True(t, ok)
	testx.Equal(t, 0, q.Size())

	testx.Equal(t, -1, q.PopOr(-1))
	v, ok = q.PopOk()
	testx.Equal(t, 0, v)
	testx.False(t, ok)
}

func TestQueuePopPanics(t *testing.T) {
	defer func() {
		testx.NotNil(t, recover())
	}()
	dsx.NewQueue[int]().Pop()
}

func TestQueueIndexOf(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 1, q.IndexOf(2))
	testx.Equal(t, -1, q.IndexOf(42))
}

func TestQueueIndexOfFunc(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.Equal(t, 0, q.IndexOfFunc(func(v int) bool { return v == 1 }))
	testx.Equal(t, -1, q.IndexOfFunc(func(v int) bool { return v == 42 }))
}

func TestQueueContains(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.True(t, q.Contains(2))
	testx.False(t, q.Contains(42))
}

func TestQueueContainsFunc(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	testx.True(t, q.ContainsFunc(func(v int) bool { return v == 3 }))
	testx.False(t, q.ContainsFunc(func(v int) bool { return v == 42 }))
}

func TestQueueSize(t *testing.T) {
	testx.Equal(t, 0, dsx.NewQueue[int]().Size())
	testx.Equal(t, 3, dsx.NewQueueFrom([]int{1, 2, 3}).Size())
}

func TestQueueClear(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	q.Clear()
	testx.Equal(t, 0, q.Size())
}

func TestQueueClone(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	clone := q.Clone()
	clone.Push(4)
	testx.Equal(t, 3, q.Size())
	testx.Equal(t, 4, clone.Size())
}

func TestQueueToSlice(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	equalSlice(t, []int{1, 2, 3}, q.ToSlice())
}

func TestQueueIter(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	values := []int{}
	for _, v := range q.Iter() {
		values = append(values, v)
	}
	equalSlice(t, []int{1, 2, 3}, values)
}

func TestQueueForEach(t *testing.T) {
	q := dsx.NewQueueFrom([]int{1, 2, 3})
	values := []int{}
	q.ForEach(func(i, v int) {
		values = append(values, v)
	})
	equalSlice(t, []int{1, 2, 3}, values)
}
