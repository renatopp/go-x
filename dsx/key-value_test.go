package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestNewKeyValue(t *testing.T) {
	kv := dsx.NewKeyValue("a", 1)
	testx.Equal(t, "a", kv.Key)
	testx.Equal(t, 1, kv.Value)
}
