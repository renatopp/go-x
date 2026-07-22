package envx

import (
	"sort"
	"testing"
)

func TestHas(t *testing.T) {
	t.Setenv("EX_HAS", "value")
	if !Has("EX_HAS") {
		t.Errorf("Has: got false, want true")
	}
	if Has("EX_HAS_MISSING") {
		t.Errorf("Has: got true, want false")
	}
}

func TestGet(t *testing.T) {
	t.Setenv("EX_GET", "value")
	if got := Get("EX_GET"); got != "value" {
		t.Errorf("Get: got %q, want %q", got, "value")
	}
	if got := Get("EX_GET_MISSING"); got != "" {
		t.Errorf("Get: got %q, want empty string", got)
	}
}

func TestGetOr(t *testing.T) {
	t.Setenv("EX_GETOR", "value")
	if got := GetOr("EX_GETOR", "default"); got != "value" {
		t.Errorf("GetOr: got %q, want %q", got, "value")
	}
	if got := GetOr("EX_GETOR_MISSING", "default"); got != "default" {
		t.Errorf("GetOr: got %q, want %q", got, "default")
	}
}

func TestGetOk(t *testing.T) {
	t.Setenv("EX_GETOK", "value")

	got, ok := GetOk("EX_GETOK")
	if !ok {
		t.Errorf("GetOk: got ok=false, want true")
	}
	if got != "value" {
		t.Errorf("GetOk: got %q, want %q", got, "value")
	}

	got, ok = GetOk("EX_GETOK_MISSING")
	if ok {
		t.Errorf("GetOk: got ok=true, want false")
	}
	if got != "" {
		t.Errorf("GetOk: got %q, want empty string", got)
	}
}

func TestSet(t *testing.T) {
	if err := Set("EX_SET", "value"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := Get("EX_SET"); got != "value" {
		t.Errorf("Set: got %q, want %q", got, "value")
	}
}

func TestUnset(t *testing.T) {
	t.Setenv("EX_UNSET", "value")
	if err := Unset("EX_UNSET"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if Has("EX_UNSET") {
		t.Errorf("Unset: expected variable to be unset")
	}
}

func TestClean(t *testing.T) {
	t.Setenv("EX_CLEAN", "value")
	Clean()
	if Has("EX_CLEAN") {
		t.Errorf("Clean: expected all variables to be unset")
	}
	if len(List()) != 0 {
		t.Errorf("Clean: expected no environment variables, got %v", List())
	}
}

func TestExpand(t *testing.T) {
	t.Setenv("EX_EXPAND", "world")
	if got := Expand("hello ${EX_EXPAND}"); got != "hello world" {
		t.Errorf("Expand: got %q, want %q", got, "hello world")
	}
	if got := Expand("hello $EX_EXPAND"); got != "hello world" {
		t.Errorf("Expand: got %q, want %q", got, "hello world")
	}
}

func TestList(t *testing.T) {
	Clean()
	t.Setenv("EX_LIST", "value")

	list := List()
	found := false
	for _, entry := range list {
		if entry == "EX_LIST=value" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("List: got %v, want to contain %q", list, "EX_LIST=value")
	}
}

func TestKeys(t *testing.T) {
	Clean()
	t.Setenv("EX_KEYS_A", "1")
	t.Setenv("EX_KEYS_B", "2")

	keys := Keys()
	sort.Strings(keys)

	want := []string{"EX_KEYS_A", "EX_KEYS_B"}
	if len(keys) != len(want) {
		t.Fatalf("Keys: got %v, want %v", keys, want)
	}
	for i := range want {
		if keys[i] != want[i] {
			t.Errorf("Keys: got %v, want %v", keys, want)
			break
		}
	}
}

func TestMap(t *testing.T) {
	Clean()
	t.Setenv("EX_MAP", "value")

	m := Map()
	if got, ok := m["EX_MAP"]; !ok || got != "value" {
		t.Errorf("Map: got %v, want to contain EX_MAP=value", m)
	}
}
