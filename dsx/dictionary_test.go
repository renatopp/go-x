package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestDictionaryNewDictionaryFrom(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.Equal(t, 1, d.Size())

	d2 := dsx.NewDictionaryFrom(nil)
	testx.Equal(t, 0, d2.Size())
	d2.Set("a", 1)
	testx.Equal(t, 1, d2.Get("a"))
}

func TestDictionaryGet(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.Equal(t, 1, d.Get("a"))
	testx.Nil(t, d.Get("missing"))
}

func TestDictionaryGetOr(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.Equal(t, 1, d.GetOr("a", -1))
	testx.Equal(t, -1, d.GetOr("missing", -1))
}

func TestDictionaryGetOk(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	v, ok := d.GetOk("a")
	testx.Equal(t, 1, v)
	testx.True(t, ok)

	v, ok = d.GetOk("missing")
	testx.Nil(t, v)
	testx.False(t, ok)
}

func TestDictionarySet(t *testing.T) {
	d := dsx.NewDictionary()
	d.Set("a", 1)
	testx.Equal(t, 1, d.Get("a"))
	d.Set("a", 2)
	testx.Equal(t, 2, d.Get("a"))
}

func TestDictionaryDelete(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2, "c": 3})
	d.Delete("a", "b", "missing")
	testx.Equal(t, 1, d.Size())
	testx.True(t, d.ContainsKey("c"))
}

func TestDictionaryContainsKey(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.True(t, d.ContainsKey("a"))
	testx.False(t, d.ContainsKey("missing"))
}

func TestDictionaryContainsValue(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.True(t, d.ContainsValue(1))
	testx.False(t, d.ContainsValue(42))
}

func TestDictionaryContainsFunc(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	testx.True(t, d.ContainsFunc(func(k string, v any) bool { return k == "a" }))
	testx.False(t, d.ContainsFunc(func(k string, v any) bool { return k == "missing" }))
}

func TestDictionarySize(t *testing.T) {
	testx.Equal(t, 0, dsx.NewDictionary().Size())
	testx.Equal(t, 2, dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2}).Size())
}

func TestDictionaryClear(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	d.Clear()
	testx.Equal(t, 0, d.Size())
}

func TestDictionaryClone(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	clone := d.Clone()
	clone.Set("b", 2)
	testx.Equal(t, 1, d.Size())
	testx.Equal(t, 2, clone.Size())
}

func TestDictionaryKeys(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	sameElements(t, []string{"a", "b"}, d.Keys())
}

func TestDictionaryValues(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	sameElements(t, []any{1, 2}, d.Values())
}

func TestDictionaryToMap(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1})
	m := d.ToMap()
	testx.Equal(t, 1, m["a"])
	m["b"] = 2
	testx.Equal(t, 2, d.Get("b"))
}

func TestDictionaryConcat(t *testing.T) {
	d1 := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	d2 := dsx.NewDictionaryFrom(map[string]any{"b": 20, "c": 3})

	merged := d1.Concat(d2)
	testx.Equal(t, 2, d1.Size())
	testx.Equal(t, 1, d1.Get("a"))
	testx.Equal(t, 3, merged.Size())
	testx.Equal(t, 1, merged.Get("a"))
	testx.Equal(t, 20, merged.Get("b"))
	testx.Equal(t, 3, merged.Get("c"))
}

func TestDictionaryAssign(t *testing.T) {
	d1 := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	d2 := dsx.NewDictionaryFrom(map[string]any{"b": 20, "c": 3})

	d1.Assign(d2)
	testx.Equal(t, 3, d1.Size())
	testx.Equal(t, 1, d1.Get("a"))
	testx.Equal(t, 20, d1.Get("b"))
	testx.Equal(t, 3, d1.Get("c"))
}

func TestDictionaryEqual(t *testing.T) {
	d1 := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	d2 := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	d3 := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 3})

	testx.True(t, d1.Equal(d2))
	testx.False(t, d1.Equal(d3))
}

func TestDictionaryIter(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	got := map[string]any{}
	for k, v := range d.Iter() {
		got[k] = v
	}
	testx.Equal(t, 2, len(got))
	testx.Equal(t, 1, got["a"])
	testx.Equal(t, 2, got["b"])
}

func TestDictionaryForEach(t *testing.T) {
	d := dsx.NewDictionaryFrom(map[string]any{"a": 1, "b": 2})
	got := map[string]any{}
	d.ForEach(func(k string, v any) {
		got[k] = v
	})
	testx.Equal(t, 2, len(got))
	testx.Equal(t, 1, got["a"])
	testx.Equal(t, 2, got["b"])
}
