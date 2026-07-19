package httpx

import (
	"sync"
	"time"
)

type breakerState int

const (
	breakerClosed breakerState = iota
	breakerOpen
	breakerHalfOpen
)

// CircuitBreaker stops a Fetcher from issuing requests to a failing
// endpoint once a failure threshold is reached, giving it time to recover.
type CircuitBreaker struct {
	mu           sync.Mutex
	threshold    int
	resetTimeout time.Duration
	state        breakerState
	failures     int
	openedAt     time.Time
}

// NewCircuitBreaker creates a CircuitBreaker that opens after threshold
// consecutive failures and attempts to recover after resetTimeout.
func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold, resetTimeout: resetTimeout}
}

// Allow reports whether a request should be attempted. Once the breaker has
// been open past its reset timeout, it moves to half-open and allows one
// probe request through.
func (b *CircuitBreaker) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.state != breakerOpen {
		return true
	}
	if time.Since(b.openedAt) < b.resetTimeout {
		return false
	}
	b.state = breakerHalfOpen
	return true
}

// Success records a successful request, closing the breaker.
func (b *CircuitBreaker) Success() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures = 0
	b.state = breakerClosed
}

// Failure records a failed request, opening the breaker once the failure
// threshold is reached.
func (b *CircuitBreaker) Failure() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures++
	if b.failures >= b.threshold {
		b.state = breakerOpen
		b.openedAt = time.Now()
	}
}
