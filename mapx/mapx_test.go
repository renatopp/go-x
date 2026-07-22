package mapx_test

import (
	"testing"

	"github.com/renatopp/go-x/mapx"
	"github.com/renatopp/go-x/testx"
)

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1}
	testx.Equal(t, 1, mapx.GetOr(m, "a", 99))
	testx.Equal(t, 99, mapx.GetOr(m, "b", 99))
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	keys := mapx.Keys(m)
	testx.Equal(t, 2, len(keys))
	testx.True(t, mapx.ContainsValue(m, 1))
	testx.True(t, mapx.ContainsValue(m, 2))
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	values := mapx.Values(m)
	testx.Equal(t, 2, len(values))
	sum := 0
	for _, v := range values {
		sum += v
	}
	testx.Equal(t, 3, sum)
}

func TestKeyOf(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	k, ok := mapx.KeyOf(m, 2)
	testx.True(t, ok)
	testx.Equal(t, "b", k)

	_, ok = mapx.KeyOf(m, 3)
	testx.False(t, ok)
}

func TestKeyOfFunc(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	k, ok := mapx.KeyOfFunc(m, func(key string, v int) bool { return v > 1 })
	testx.True(t, ok)
	testx.Equal(t, "b", k)

	_, ok = mapx.KeyOfFunc(m, func(key string, v int) bool { return v > 10 })
	testx.False(t, ok)
}

func TestSize(t *testing.T) {
	testx.Equal(t, 0, mapx.Size(map[string]int{}))
	testx.Equal(t, 2, mapx.Size(map[string]int{"a": 1, "b": 2}))
}

func TestConcat(t *testing.T) {
	m := mapx.Concat(map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 3})
	testx.Equal(t, 2, mapx.Size(m))
	testx.Equal(t, 3, m["a"])
	testx.Equal(t, 2, m["b"])
}

func TestAssign(t *testing.T) {
	m := map[string]int{"a": 1}
	mapx.Assign(m, map[string]int{"b": 2}, map[string]int{"a": 3})
	testx.Equal(t, 2, mapx.Size(m))
	testx.Equal(t, 3, m["a"])
	testx.Equal(t, 2, m["b"])
}

func TestRemove(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	mapx.Remove(m, "a", "c")
	testx.Equal(t, 1, mapx.Size(m))
	testx.True(t, mapx.ContainsKey(m, "b"))
}

func TestClear(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	mapx.Clear(m)
	testx.Equal(t, 0, mapx.Size(m))
}

func TestContainsKey(t *testing.T) {
	m := map[string]int{"a": 1}
	testx.True(t, mapx.ContainsKey(m, "a"))
	testx.False(t, mapx.ContainsKey(m, "b"))
}

func TestContainsValue(t *testing.T) {
	m := map[string]int{"a": 1}
	testx.True(t, mapx.ContainsValue(m, 1))
	testx.False(t, mapx.ContainsValue(m, 2))
}

func TestContainsFunc(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	testx.True(t, mapx.ContainsFunc(m, func(k string, v int) bool { return v > 1 }))
	testx.False(t, mapx.ContainsFunc(m, func(k string, v int) bool { return v > 10 }))
}

func TestForEach(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	sum := 0
	mapx.ForEach(m, func(k string, v int) { sum += v })
	testx.Equal(t, 3, sum)
}

func TestForEachValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	sum := 0
	mapx.ForEachValue(m, func(v int) { sum += v })
	testx.Equal(t, 3, sum)
}

func TestForEachKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	count := 0
	mapx.ForEachKey(m, func(k string) { count++ })
	testx.Equal(t, 2, count)
}
