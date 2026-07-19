package logx

import "sync/atomic"

// Sampler decides whether a log record should be emitted, letting only one
// out of every N records through.
type Sampler struct {
	n       uint64
	counter atomic.Uint64
}

// NewSampler creates a Sampler that allows one out of every n records
// through. n <= 1 allows every record.
func NewSampler(n int) *Sampler {
	if n < 1 {
		n = 1
	}
	return &Sampler{n: uint64(n)}
}

// Allow reports whether the current record should be logged, advancing the
// sampler's internal counter.
func (s *Sampler) Allow() bool {
	return s.counter.Add(1)%s.n == 1
}
