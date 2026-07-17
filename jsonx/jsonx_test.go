package jsonx_test

import (
	"testing"

	"github.com/renatopp/go-x/jsonx"
	"github.com/renatopp/go-x/testx"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshal(t *testing.T) {
	b, err := jsonx.Marshal(person{Name: "Ana", Age: 30})
	testx.Nil(t, err)
	testx.Equal(t, `{"name":"Ana","age":30}`, string(b))
}

func TestMarshalString(t *testing.T) {
	s, err := jsonx.MarshalString(person{Name: "Ana", Age: 30})
	testx.Nil(t, err)
	testx.Equal(t, `{"name":"Ana","age":30}`, s)
}

func TestMarshalIndent(t *testing.T) {
	b, err := jsonx.MarshalIndent(person{Name: "Ana", Age: 30}, "", "  ")
	testx.Nil(t, err)
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", string(b))
}

func TestMarshalIndentString(t *testing.T) {
	s, err := jsonx.MarshalIndentString(person{Name: "Ana", Age: 30}, "", "  ")
	testx.Nil(t, err)
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", s)
}

func TestForceMarshal(t *testing.T) {
	b := jsonx.ForceMarshal(person{Name: "Ana", Age: 30})
	testx.Equal(t, `{"name":"Ana","age":30}`, string(b))
}

func TestForceMarshalString(t *testing.T) {
	s := jsonx.ForceMarshalString(person{Name: "Ana", Age: 30})
	testx.Equal(t, `{"name":"Ana","age":30}`, s)
}

func TestForceMarshalIndent(t *testing.T) {
	b := jsonx.ForceMarshalIndent(person{Name: "Ana", Age: 30}, "", "  ")
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", string(b))
}

func TestForceMarshalIndentString(t *testing.T) {
	s := jsonx.ForceMarshalIndentString(person{Name: "Ana", Age: 30}, "", "  ")
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", s)
}

func TestUnmarshal(t *testing.T) {
	var p person
	err := jsonx.Unmarshal([]byte(`{"name":"Ana","age":30}`), &p)
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalString(t *testing.T) {
	var p person
	err := jsonx.UnmarshalString(`{"name":"Ana","age":30}`, &p)
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalAs(t *testing.T) {
	p, err := jsonx.UnmarshalAs[person]([]byte(`{"name":"Ana","age":30}`))
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestUnmarshalStringAs(t *testing.T) {
	p, err := jsonx.UnmarshalStringAs[person](`{"name":"Ana","age":30}`)
	testx.Nil(t, err)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestForceUnmarshalAs(t *testing.T) {
	p := jsonx.ForceUnmarshalAs[person]([]byte(`{"name":"Ana","age":30}`))
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestForceUnmarshalStringAs(t *testing.T) {
	p := jsonx.ForceUnmarshalStringAs[person](`{"name":"Ana","age":30}`)
	testx.Equal(t, person{Name: "Ana", Age: 30}, p)
}

func TestPrettyPrint(t *testing.T) {
	b, err := jsonx.PrettyPrint([]byte(`{"name":"Ana","age":30}`))
	testx.Nil(t, err)
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", string(b))
}

func TestPrettyPrintString(t *testing.T) {
	s, err := jsonx.PrettyPrintString(`{"name":"Ana","age":30}`)
	testx.Nil(t, err)
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", s)
}

func TestForcePrettyPrint(t *testing.T) {
	b := jsonx.ForcePrettyPrint([]byte(`{"name":"Ana","age":30}`))
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", string(b))
}

func TestForcePrettyPrintString(t *testing.T) {
	s := jsonx.ForcePrettyPrintString(`{"name":"Ana","age":30}`)
	testx.Equal(t, "{\n  \"name\": \"Ana\",\n  \"age\": 30\n}", s)
}

func TestIsValid(t *testing.T) {
	testx.True(t, jsonx.IsValid([]byte(`{"name":"Ana"}`)))
	testx.False(t, jsonx.IsValid([]byte(`{invalid}`)))
}

func TestIsValidString(t *testing.T) {
	testx.True(t, jsonx.IsValidString(`{"name":"Ana"}`))
	testx.False(t, jsonx.IsValidString(`{invalid}`))
}

func TestCompact(t *testing.T) {
	b, err := jsonx.Compact([]byte("{\n  \"name\": \"Ana\",\n  \"age\": 30\n}"))
	testx.Nil(t, err)
	testx.Equal(t, `{"name":"Ana","age":30}`, string(b))
}

func TestCompactString(t *testing.T) {
	s, err := jsonx.CompactString("{\n  \"name\": \"Ana\",\n  \"age\": 30\n}")
	testx.Nil(t, err)
	testx.Equal(t, `{"name":"Ana","age":30}`, s)
}

func TestForceCompact(t *testing.T) {
	b := jsonx.ForceCompact([]byte("{\n  \"name\": \"Ana\",\n  \"age\": 30\n}"))
	testx.Equal(t, `{"name":"Ana","age":30}`, string(b))
}

func TestForceCompactString(t *testing.T) {
	s := jsonx.ForceCompactString("{\n  \"name\": \"Ana\",\n  \"age\": 30\n}")
	testx.Equal(t, `{"name":"Ana","age":30}`, s)
}
