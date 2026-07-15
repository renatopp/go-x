package dsx_test

import (
	"testing"

	"github.com/renatopp/go-x/dsx"
	"github.com/renatopp/go-x/testx"
)

func TestMarkovChainNextDeterministic(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	testx.Equal(t, "b", c.Next("a"))
	testx.Equal(t, "", c.Next("b"))
	testx.Equal(t, "", c.Next("missing"))
}

func TestMarkovChainRandom(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	testx.Equal(t, "a", c.Random())
}

func TestMarkovChainRandomAmongStates(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "x", 1)
	c.Add("b", "y", 1)

	for range 10 {
		state := c.Random()
		testx.True(t, state == "a" || state == "b")
	}
}

func TestMarkovChainNextFromEmptyUsesRandom(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	testx.Equal(t, "b", c.Next(""))
}

func TestMarkovChainNextAccumulatesWeight(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)
	c.Add("a", "b", 2)

	testx.Equal(t, "b", c.Next("a"))
}

func TestMarkovChainGenerateFrom(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	testx.Equal(t, "ab", c.GenerateFrom(5, "a"))
}

func TestMarkovChainGenerate(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	testx.Equal(t, "b", c.Generate(5))
}

func TestMarkovChainGenerateStopsAtDeadEnd(t *testing.T) {
	c := dsx.NewMarkovChain()
	c.Add("a", "b", 1)

	result := c.GenerateFrom(100, "a")
	testx.Equal(t, "ab", result)
}
