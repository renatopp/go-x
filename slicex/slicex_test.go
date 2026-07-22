package slicex_test

import (
	"testing"

	"github.com/renatopp/go-x/slicex"
	"github.com/renatopp/go-x/testx"
)

func TestNew(t *testing.T) {
	s := slicex.New(1, 2, 3)
	testx.Equal(t, 3, len(s))
	testx.Equal(t, 1, s[0])
	testx.Equal(t, 2, s[1])
	testx.Equal(t, 3, s[2])
}

func TestAppend(t *testing.T) {
	s := slicex.Append([]int{1, 2}, 3, 4)
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestAppendSlice(t *testing.T) {
	s := slicex.AppendSlice([]int{1, 2}, []int{3, 4})
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestPrepend(t *testing.T) {
	s := slicex.Prepend([]int{3, 4}, 1, 2)
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestPrependSlice(t *testing.T) {
	s := slicex.PrependSlice([]int{3, 4}, []int{1, 2})
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestInsert(t *testing.T) {
	s := slicex.Insert([]int{1, 2, 4}, 2, 3)
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))

	s = slicex.Insert([]int{1, 2, 4}, -1, 3)
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestInsertSlice(t *testing.T) {
	s := slicex.InsertSlice([]int{1, 4}, 1, []int{2, 3})
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))

	s = slicex.InsertSlice([]int{1, 4}, -1, []int{2, 3})
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4}))
}

func TestRemove(t *testing.T) {
	s := slicex.Remove([]int{1, 2, 3, 4}, 1)
	testx.True(t, slicex.Equal(s, []int{1, 3, 4}))

	s = slicex.Remove([]int{1, 2, 3, 4}, -1)
	testx.True(t, slicex.Equal(s, []int{1, 2, 3}))
}

func TestRemoveRange(t *testing.T) {
	s := slicex.RemoveRange([]int{1, 2, 3, 4, 5}, 1, 3)
	testx.True(t, slicex.Equal(s, []int{1, 4, 5}))

	s = slicex.RemoveRange([]int{1, 2, 3, 4, 5}, 3, 1)
	testx.True(t, slicex.Equal(s, []int{1, 4, 5}))

	s = slicex.RemoveRange([]int{1, 2, 3, 4, 5}, 1, -1)
	testx.True(t, slicex.Equal(s, []int{1, 5}))
}

func TestRemoveFunc(t *testing.T) {
	s := slicex.RemoveFunc([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	testx.True(t, slicex.Equal(s, []int{1, 3, 5}))
}

func TestRemoveValue(t *testing.T) {
	s := slicex.RemoveValue([]int{1, 2, 3, 2, 1}, 2)
	testx.True(t, slicex.Equal(s, []int{1, 3, 1}))
}

func TestAssign(t *testing.T) {
	s := slicex.Assign([]int{1, 2}, []int{3, 4}, []int{5, 6})
	testx.True(t, slicex.Equal(s, []int{1, 2, 3, 4, 5, 6}))
}

func TestReversed(t *testing.T) {
	original := []int{1, 2, 3}
	s := slicex.Reversed(original)
	testx.True(t, slicex.Equal(s, []int{3, 2, 1}))
	testx.True(t, slicex.Equal(original, []int{1, 2, 3}))
}

func TestMap(t *testing.T) {
	s := slicex.Map([]int{1, 2, 3}, func(v int) int { return v * 2 })
	testx.True(t, slicex.Equal(s, []int{2, 4, 6}))
}

func TestFilter(t *testing.T) {
	s := slicex.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	testx.True(t, slicex.Equal(s, []int{2, 4}))
}

func TestReduce(t *testing.T) {
	sum := slicex.Reduce([]int{1, 2, 3, 4}, func(acc, v int) int { return acc + v }, 0)
	testx.Equal(t, 10, sum)
}

func TestForEach(t *testing.T) {
	sum := 0
	slicex.ForEach([]int{1, 2, 3}, func(v int) { sum += v })
	testx.Equal(t, 6, sum)
}

func TestForEachIndex(t *testing.T) {
	sum := 0
	slicex.ForEachIndex([]int{1, 2, 3}, func(i, v int) { sum += i * v })
	testx.Equal(t, 8, sum) // 0*1 + 1*2 + 2*3
}
