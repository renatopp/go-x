package timex

import (
	"math"
	"math/rand/v2"
	"time"
)

type Backoff interface {
	Duration(attempt int) time.Duration
}

type ExponentialBackoff struct {
	Initial    time.Duration
	Max        time.Duration
	Multiplier float64
	Jitter     float64
}

func NewExponentialBackoff(initial, max time.Duration, multiplier, jitter float64) *ExponentialBackoff {
	if initial <= 0 {
		initial = time.Millisecond
	}
	if max <= 0 {
		max = time.Second * 30
	}
	if multiplier <= 1 {
		multiplier = 2
	}
	if jitter < 0 || jitter > 1 {
		jitter = 0.5
	}
	return &ExponentialBackoff{
		Initial:    initial,
		Max:        max,
		Multiplier: multiplier,
		Jitter:     jitter,
	}
}

func (b *ExponentialBackoff) Duration(attempt int) time.Duration {
	if attempt <= 0 {
		return b.Initial
	}
	d := float64(b.Initial) * math.Pow(b.Multiplier, float64(attempt-1))
	if d > float64(b.Max) {
		d = float64(b.Max)
	}
	if b.Jitter > 0 {
		jitter := d * b.Jitter * (rand.Float64()*2 - 1)
		d += jitter
		if d < 0 {
			d = 0
		}
	}
	return time.Duration(d)
}

type LinearBackoff struct {
	Initial time.Duration
	Max     time.Duration
	Step    time.Duration
	Jitter  float64
}

func NewLinearBackoff(initial, max, step time.Duration, jitter float64) *LinearBackoff {
	if initial <= 0 {
		initial = time.Millisecond
	}
	if max <= 0 {
		max = time.Second * 30
	}
	if step <= 0 {
		step = time.Millisecond * 100
	}
	if jitter < 0 || jitter > 1 {
		jitter = 0.5
	}
	return &LinearBackoff{
		Initial: initial,
		Max:     max,
		Step:    step,
		Jitter:  jitter,
	}
}

func (b *LinearBackoff) Duration(attempt int) time.Duration {
	if attempt <= 0 {
		return b.Initial
	}
	d := min(b.Initial+b.Step*time.Duration(attempt-1), b.Max)
	if b.Jitter > 0 {
		jitter := float64(d) * b.Jitter * (rand.Float64()*2 - 1)
		d += time.Duration(jitter)
		if d < 0 {
			d = 0
		}
	}
	return d
}

type ConstantBackoff struct {
	Delay  time.Duration
	Jitter float64
}

func NewConstantBackoff(duration time.Duration, jitter float64) *ConstantBackoff {
	if duration <= 0 {
		duration = time.Millisecond * 100
	}
	if jitter < 0 || jitter > 1 {
		jitter = 0.5
	}
	return &ConstantBackoff{
		Delay:  duration,
		Jitter: jitter,
	}
}

func (b *ConstantBackoff) Duration(attempt int) time.Duration {
	d := b.Delay
	if b.Jitter > 0 {
		jitter := float64(d) * b.Jitter * (rand.Float64()*2 - 1)
		d += time.Duration(jitter)
		if d < 0 {
			d = 0
		}
	}
	return d
}
