package yamlx_test

import (
	"testing"

	"github.com/renatopp/go-x/testx"
	"github.com/renatopp/go-x/yamlx"
)

type person struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func TestMarshal(t *testing.T) {
	b, err := yamlx.Marshal(person{Name: "Ana", Age: 30})
	testx.Nil(t, err)
	testx.Equal(t, "name: Ana\nage: 30\n", string(b))
}

func TestMarshalString(t *testing.T) {
	s, err := yamlx.MarshalString(person{Name: "Ana", Age: 30})
	testx.Nil(t, err)
	testx.Equal(t, "name: Ana\nage: 30\n", s)
}

func TestMarshalIndent(t *testing.T) {
	b, err := yamlx.MarshalIndent(map[string]any{"outer": map[string]any{"inner": 1}}, 4)
	testx.Nil(t, err)
	testx.Equal(t, "outer:\n    inner: 1\n", string(b))
}

func TestMarshalIndentString(t *testing.T) {
	s, err := yamlx.MarshalIndentString(map[string]any{"outer": map[string]any{"inner": 1}}, 4)
	testx.Nil(t, err)
	testx.Equal(t, "outer:\n    inner: 1\n", s)
}

func TestForceMarshal(t *testing.T) {
	b := yamlx.ForceMarshal(person{Name: "Ana", Age: 30})
	testx.Equal(t, "name: Ana\nage: 30\n", string(b))
}

func TestForceMarshalString(t *testing.T) {
	s := yamlx.ForceMarshalString(person{Name: "Ana", Age: 30})
	testx.Equal(t, "name: Ana\nage: 30\n", s)
}

func TestForceMarshalIndent(t *testing.T) {
	b := yamlx.ForceMarshalIndent(map[string]any{"outer": map[string]any{"inner": 1}}, 4)
	testx.Equal(t, "outer:\n    inner: 1\n", string(b))
}

func TestForceMarshalIndentString(t *testing.T) {
	s := yamlx.ForceMarshalIndentString(map[string]any{"outer": map[string]any{"inner": 1}}, 4)
	testx.Equal(t, "outer:\n    inner: 1\n", s)
}

func TestUnmarshal(t *testing.T) {
	var p person
	err := yamlx.Unmarshal([]byte("name: Ana\nage: 30\n"), &p)
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalString(t *testing.T) {
	var p person
	err := yamlx.UnmarshalString("name: Ana\nage: 30\n", &p)
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalAs(t *testing.T) {
	p, err := yamlx.UnmarshalAs[person]([]byte("name: Ana\nage: 30\n"))
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalStringAs(t *testing.T) {
	p, err := yamlx.UnmarshalStringAs[person]("name: Ana\nage: 30\n")
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestForceUnmarshalAs(t *testing.T) {
	p := yamlx.ForceUnmarshalAs[person]([]byte("name: Ana\nage: 30\n"))
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestForceUnmarshalStringAs(t *testing.T) {
	p := yamlx.ForceUnmarshalStringAs[person]("name: Ana\nage: 30\n")
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestPrettyPrint(t *testing.T) {
	b, err := yamlx.PrettyPrint([]byte("name:   Ana\nage:  30\n"))
	testx.Nil(t, err)
	testx.Equal(t, "name: Ana\nage: 30\n", string(b))
}

func TestPrettyPrintString(t *testing.T) {
	s, err := yamlx.PrettyPrintString("name:   Ana\nage:  30\n")
	testx.Nil(t, err)
	testx.Equal(t, "name: Ana\nage: 30\n", s)
}

func TestForcePrettyPrint(t *testing.T) {
	b := yamlx.ForcePrettyPrint([]byte("name:   Ana\nage:  30\n"))
	testx.Equal(t, "name: Ana\nage: 30\n", string(b))
}

func TestForcePrettyPrintString(t *testing.T) {
	s := yamlx.ForcePrettyPrintString("name:   Ana\nage:  30\n")
	testx.Equal(t, "name: Ana\nage: 30\n", s)
}

func TestIsValid(t *testing.T) {
	testx.True(t, yamlx.IsValid([]byte("name: Ana\nage: 30\n")))
	testx.False(t, yamlx.IsValid([]byte("name: [Ana\n")))
}

func TestIsValidString(t *testing.T) {
	testx.True(t, yamlx.IsValidString("name: Ana\nage: 30\n"))
	testx.False(t, yamlx.IsValidString("name: [Ana\n"))
}

func TestCompact(t *testing.T) {
	b, err := yamlx.Compact([]byte("name: Ana\nage: 30\n"))
	testx.Nil(t, err)
	testx.Equal(t, "{name: Ana, age: 30}\n", string(b))
}

func TestCompactString(t *testing.T) {
	s, err := yamlx.CompactString("name: Ana\nage: 30\n")
	testx.Nil(t, err)
	testx.Equal(t, "{name: Ana, age: 30}\n", s)
}

func TestForceCompact(t *testing.T) {
	b := yamlx.ForceCompact([]byte("name: Ana\nage: 30\n"))
	testx.Equal(t, "{name: Ana, age: 30}\n", string(b))
}

func TestForceCompactString(t *testing.T) {
	s := yamlx.ForceCompactString("name: Ana\nage: 30\n")
	testx.Equal(t, "{name: Ana, age: 30}\n", s)
}
