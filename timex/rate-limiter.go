package timex

import (
	"sync"
	"time"
)

type RateLimiter interface {
	Consume(n int) bool
	Release(n int) bool
	ReleaseAll()
}

// TokenRateLimiter is a rate limiter that uses a token bucket algorithm.
// It allows a certain number of tokens to be consumed within a given time
// window, and replenishes tokens at a fixed rate.
type TokenRateLimiter struct {
	mu             sync.Mutex
	tokens         int
	maxTokens      int
	refillInterval time.Duration
	lastRefill     time.Time
}

func NewTokenRateLimiter(maxTokens int, refillInterval time.Duration) *TokenRateLimiter {
	return &TokenRateLimiter{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillInterval: refillInterval,
		lastRefill:     time.Now(),
	}
}

func (r *TokenRateLimiter) Consume(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill)

	if elapsed >= r.refillInterval {
		refills := int(elapsed / r.refillInterval)
		r.tokens = min(r.maxTokens, r.tokens+refills)
		r.lastRefill = now
	}

	if r.tokens >= n {
		r.tokens -= n
		return true
	}

	return false
}

func (r *TokenRateLimiter) Release(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.tokens+n > r.maxTokens {
		return false
	}

	r.tokens += n
	return true
}

func (r *TokenRateLimiter) ReleaseAll() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tokens = r.maxTokens
}

// TimeWindowRateLimiter is a rate limiter that allows a certain number of
// requests within a specified time window, for example, 100 requests per
// minute. It resets the count after the time window expires.
type TimeWindowRateLimiter struct {
	mu          sync.Mutex
	requests    int
	maxRequests int
	window      time.Duration
	startTime   time.Time
}

func NewTimeWindowRateLimiter(maxRequests int, window time.Duration) *TimeWindowRateLimiter {
	return &TimeWindowRateLimiter{
		maxRequests: maxRequests,
		window:      window,
		startTime:   time.Now(),
	}
}

func (r *TimeWindowRateLimiter) Consume(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.startTime) >= r.window {
		r.requests = 0
		r.startTime = now
	}

	if r.requests+n <= r.maxRequests {
		r.requests += n
		return true
	}

	return false
}

func (r *TimeWindowRateLimiter) Release(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.requests-n < 0 {
		return false
	}

	r.requests -= n
	return true
}

func (r *TimeWindowRateLimiter) ReleaseAll() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.requests = 0
}

// MultiRateLimiter is a rate limiter that combines multiple rate limiters. It
// allows a request to proceed only if all underlying rate limiters allow it.
// This can be useful for enforcing multiple constraints, such as both a token
// bucket and a time window limit.
type MultiRateLimiter struct {
	mu       sync.Mutex
	limiters []RateLimiter
}

func NewMultiRateLimiter(limiters ...RateLimiter) *MultiRateLimiter {
	return &MultiRateLimiter{
		limiters: limiters,
	}
}

func (r *MultiRateLimiter) Consume(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, limiter := range r.limiters {
		if !limiter.Consume(n) {
			// Rollback previously consumed limiters
			for j := 0; j < i; j++ {
				r.limiters[j].Release(n)
			}
			return false
		}
	}

	return true
}

func (r *MultiRateLimiter) Release(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	success := true
	for _, limiter := range r.limiters {
		if !limiter.Release(n) {
			success = false
		}
	}
	return success
}

func (r *MultiRateLimiter) ReleaseAll() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, limiter := range r.limiters {
		limiter.ReleaseAll()
	}
}
