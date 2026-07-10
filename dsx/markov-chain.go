package dsx

import (
	"strings"

	"github.com/renatopp/go-x/randx"
)

// MarkovChain represents a Markov chain, which is a stochastic model that
// describes a sequence of possible events where the probability of each event
// depends only on the state attained in the previous event.
type MarkovChain struct {
	totals      map[string]float64
	transitions map[string]map[string]float64
}

// NewMarkovChain creates a new Markov chain instance with initialized maps for
// totals and transitions.
func NewMarkovChain() *MarkovChain {
	c := &MarkovChain{
		totals:      make(map[string]float64),
		transitions: make(map[string]map[string]float64),
	}
	return c
}

// Add adds a transition from one state to another with a specified weight. If
// the transition already exists, it increments the weight.
func (c *MarkovChain) Add(from, to string, weight float64) {
	if c.transitions[from] == nil {
		c.transitions[from] = make(map[string]float64)
	}
	if c.transitions[to] == nil {
		c.transitions[to] = make(map[string]float64)
	}
	c.transitions[from][to] += weight
	c.totals[from] += weight
}

// Random returns a random state equiprobably from the set of all states in the
// Markov chain. This is used as the starting point for generating sequences.
func (c *MarkovChain) Random() string {
	var total float64
	for _, t := range c.totals {
		total += t
	}
	r := randx.FloatN(total)
	var cumulative float64
	for state, weight := range c.totals {
		cumulative += weight
		if r < cumulative {
			return state
		}
	}
	return ""
}

// Next returns the next state given the current state. If the current state
// has no transitions, it stops. If the current state is empty, it also returns
// a random state from the Markov chain.
func (c *MarkovChain) Next(from string) string {
	if c.transitions[from] == nil && from == "" {
		from = c.Random()
	}

	transitions := c.transitions[from]
	if transitions == nil {
		return ""
	}

	total := c.totals[from]
	r := randx.FloatN(total)
	var cumulative float64
	for to, weight := range transitions {
		cumulative += weight
		if r < cumulative {
			return to
		}
	}
	return ""
}

// GenerateFrom generates a sequence of states starting from a given state. It
// continues to generate states until it reaches a state with no transitions or
// until it has generated the specified number of states. The generated sequence is
// returned as a string.
func (c *MarkovChain) GenerateFrom(n int, from string) string {
	var result strings.Builder
	result.WriteString(from)
	next := from
	for range n {
		next = c.Next(next)
		if next == "" {
			break
		}
		result.WriteString(next)
	}
	return result.String()
}

// Generate generates a sequence of states starting from a random state. It
// continues to generate states until it reaches a state with no transitions or
// until it has generated the specified number of states. The generated sequence is
// returned as a string.
func (c *MarkovChain) Generate(n int) string {
	return c.GenerateFrom(n, "")
}
