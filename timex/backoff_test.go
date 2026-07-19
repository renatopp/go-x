package timex

import (
	"math"
	"testing"
	"time"
)

func TestNewExponentialBackoff(t *testing.T) {
	tests := []struct {
		name       string
		initial    time.Duration
		max        time.Duration
		multiplier float64
		jitter     float64
		expected   ExponentialBackoff
	}{
		{
			name:       "all valid values",
			initial:    time.Millisecond * 100,
			max:        time.Second * 60,
			multiplier: 2.5,
			jitter:     0.3,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond * 100,
				Max:        time.Second * 60,
				Multiplier: 2.5,
				Jitter:     0.3,
			},
		},
		{
			name:       "zero initial defaults to 1ms",
			initial:    0,
			max:        time.Second,
			multiplier: 2,
			jitter:     0.5,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond,
				Max:        time.Second,
				Multiplier: 2,
				Jitter:     0.5,
			},
		},
		{
			name:       "zero max defaults to 30s",
			initial:    time.Millisecond * 100,
			max:        0,
			multiplier: 2,
			jitter:     0.5,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond * 100,
				Max:        time.Second * 30,
				Multiplier: 2,
				Jitter:     0.5,
			},
		},
		{
			name:       "multiplier <= 1 defaults to 2",
			initial:    time.Millisecond,
			max:        time.Second,
			multiplier: 1,
			jitter:     0.5,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond,
				Max:        time.Second,
				Multiplier: 2,
				Jitter:     0.5,
			},
		},
		{
			name:       "negative jitter defaults to 0.5",
			initial:    time.Millisecond,
			max:        time.Second,
			multiplier: 2,
			jitter:     -0.1,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond,
				Max:        time.Second,
				Multiplier: 2,
				Jitter:     0.5,
			},
		},
		{
			name:       "jitter > 1 defaults to 0.5",
			initial:    time.Millisecond,
			max:        time.Second,
			multiplier: 2,
			jitter:     1.5,
			expected: ExponentialBackoff{
				Initial:    time.Millisecond,
				Max:        time.Second,
				Multiplier: 2,
				Jitter:     0.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewExponentialBackoff(tt.initial, tt.max, tt.multiplier, tt.jitter)
			if b.Initial != tt.expected.Initial {
				t.Errorf("Initial: got %v, want %v", b.Initial, tt.expected.Initial)
			}
			if b.Max != tt.expected.Max {
				t.Errorf("Max: got %v, want %v", b.Max, tt.expected.Max)
			}
			if b.Multiplier != tt.expected.Multiplier {
				t.Errorf("Multiplier: got %v, want %v", b.Multiplier, tt.expected.Multiplier)
			}
			if b.Jitter != tt.expected.Jitter {
				t.Errorf("Jitter: got %v, want %v", b.Jitter, tt.expected.Jitter)
			}
		})
	}
}

func TestExponentialBackoffDuration(t *testing.T) {
	b := NewExponentialBackoff(time.Millisecond*100, time.Second*10, 2, 0)

	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, time.Millisecond * 100},
		{1, time.Millisecond * 100},
		{2, time.Millisecond * 200},
		{3, time.Millisecond * 400},
		{4, time.Millisecond * 800},
		{5, time.Millisecond * 1600},
		{6, time.Millisecond * 3200},
		{7, time.Millisecond * 6400},
		{8, time.Second * 10}, // capped at max
		{9, time.Second * 10}, // capped at max
	}

	for _, tt := range tests {
		t.Run("attempt_"+string(rune(tt.attempt)), func(t *testing.T) {
			d := b.Duration(tt.attempt)
			if d != tt.expected {
				t.Errorf("attempt %d: got %v, want %v", tt.attempt, d, tt.expected)
			}
		})
	}
}

func TestExponentialBackoffDurationWithJitter(t *testing.T) {
	b := NewExponentialBackoff(time.Millisecond*100, time.Second*10, 2, 0.5)

	for attempt := 0; attempt <= 5; attempt++ {
		d := b.Duration(attempt)

		var expected float64
		if attempt <= 0 {
			expected = float64(time.Millisecond * 100)
		} else {
			expected = float64(time.Millisecond*100) * math.Pow(2, float64(attempt-1))
		}

		if expected > float64(b.Max) {
			expected = float64(b.Max)
		}

		// With jitter 0.5, duration should be within expected ± (expected * 0.5)
		minExpected := expected * (1 - 0.5)
		maxExpected := expected * (1 + 0.5)

		if float64(d) < minExpected || float64(d) > maxExpected {
			t.Errorf("attempt %d: duration %v not in range [%v, %v]", attempt, d, time.Duration(minExpected), time.Duration(maxExpected))
		}
	}
}

func TestNewLinearBackoff(t *testing.T) {
	tests := []struct {
		name     string
		initial  time.Duration
		max      time.Duration
		step     time.Duration
		jitter   float64
		expected LinearBackoff
	}{
		{
			name:     "all valid values",
			initial:  time.Millisecond * 100,
			max:      time.Second * 60,
			step:     time.Millisecond * 200,
			jitter:   0.3,
			expected: LinearBackoff{
				Initial: time.Millisecond * 100,
				Max:     time.Second * 60,
				Step:    time.Millisecond * 200,
				Jitter:  0.3,
			},
		},
		{
			name:     "zero initial defaults to 1ms",
			initial:  0,
			max:      time.Second,
			step:     time.Millisecond * 100,
			jitter:   0.5,
			expected: LinearBackoff{
				Initial: time.Millisecond,
				Max:     time.Second,
				Step:    time.Millisecond * 100,
				Jitter:  0.5,
			},
		},
		{
			name:     "zero max defaults to 30s",
			initial:  time.Millisecond * 100,
			max:      0,
			step:     time.Millisecond * 100,
			jitter:   0.5,
			expected: LinearBackoff{
				Initial: time.Millisecond * 100,
				Max:     time.Second * 30,
				Step:    time.Millisecond * 100,
				Jitter:  0.5,
			},
		},
		{
			name:     "zero step defaults to 100ms",
			initial:  time.Millisecond * 100,
			max:      time.Second,
			step:     0,
			jitter:   0.5,
			expected: LinearBackoff{
				Initial: time.Millisecond * 100,
				Max:     time.Second,
				Step:    time.Millisecond * 100,
				Jitter:  0.5,
			},
		},
		{
			name:     "negative jitter defaults to 0.5",
			initial:  time.Millisecond,
			max:      time.Second,
			step:     time.Millisecond * 100,
			jitter:   -0.1,
			expected: LinearBackoff{
				Initial: time.Millisecond,
				Max:     time.Second,
				Step:    time.Millisecond * 100,
				Jitter:  0.5,
			},
		},
		{
			name:     "jitter > 1 defaults to 0.5",
			initial:  time.Millisecond,
			max:      time.Second,
			step:     time.Millisecond * 100,
			jitter:   1.5,
			expected: LinearBackoff{
				Initial: time.Millisecond,
				Max:     time.Second,
				Step:    time.Millisecond * 100,
				Jitter:  0.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLinearBackoff(tt.initial, tt.max, tt.step, tt.jitter)
			if b.Initial != tt.expected.Initial {
				t.Errorf("Initial: got %v, want %v", b.Initial, tt.expected.Initial)
			}
			if b.Max != tt.expected.Max {
				t.Errorf("Max: got %v, want %v", b.Max, tt.expected.Max)
			}
			if b.Step != tt.expected.Step {
				t.Errorf("Step: got %v, want %v", b.Step, tt.expected.Step)
			}
			if b.Jitter != tt.expected.Jitter {
				t.Errorf("Jitter: got %v, want %v", b.Jitter, tt.expected.Jitter)
			}
		})
	}
}

func TestLinearBackoffDuration(t *testing.T) {
	b := NewLinearBackoff(time.Millisecond*100, time.Second*2, time.Millisecond*300, 0)

	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, time.Millisecond * 100},
		{1, time.Millisecond * 100},
		{2, time.Millisecond * 400},
		{3, time.Millisecond * 700},
		{4, time.Millisecond * 1000},
		{5, time.Millisecond * 1300},
		{6, time.Millisecond * 1600},
		{7, time.Millisecond * 1900},
		{8, time.Second * 2}, // capped at max
		{9, time.Second * 2}, // capped at max
	}

	for _, tt := range tests {
		t.Run("attempt_"+string(rune(tt.attempt)), func(t *testing.T) {
			d := b.Duration(tt.attempt)
			if d != tt.expected {
				t.Errorf("attempt %d: got %v, want %v", tt.attempt, d, tt.expected)
			}
		})
	}
}

func TestLinearBackoffDurationWithJitter(t *testing.T) {
	b := NewLinearBackoff(time.Millisecond*100, time.Second*10, time.Millisecond*100, 0.5)

	for attempt := 0; attempt <= 5; attempt++ {
		d := b.Duration(attempt)

		var expected float64
		if attempt <= 0 {
			expected = float64(time.Millisecond * 100)
		} else {
			expected = float64(time.Millisecond*100 + time.Millisecond*100*time.Duration(attempt-1))
		}

		if expected > float64(b.Max) {
			expected = float64(b.Max)
		}

		// With jitter 0.5, duration should be within expected ± (expected * 0.5)
		minExpected := expected * (1 - 0.5)
		maxExpected := expected * (1 + 0.5)

		if float64(d) < minExpected || float64(d) > maxExpected {
			t.Errorf("attempt %d: duration %v not in range [%v, %v]", attempt, d, time.Duration(minExpected), time.Duration(maxExpected))
		}
	}
}

func TestNewConstantBackoff(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		jitter   float64
		expected ConstantBackoff
	}{
		{
			name:     "all valid values",
			duration: time.Millisecond * 500,
			jitter:   0.3,
			expected: ConstantBackoff{
				Delay:  time.Millisecond * 500,
				Jitter: 0.3,
			},
		},
		{
			name:     "zero duration defaults to 100ms",
			duration: 0,
			jitter:   0.5,
			expected: ConstantBackoff{
				Delay:  time.Millisecond * 100,
				Jitter: 0.5,
			},
		},
		{
			name:     "negative jitter defaults to 0.5",
			duration: time.Millisecond * 100,
			jitter:   -0.1,
			expected: ConstantBackoff{
				Delay:  time.Millisecond * 100,
				Jitter: 0.5,
			},
		},
		{
			name:     "jitter > 1 defaults to 0.5",
			duration: time.Millisecond * 100,
			jitter:   1.5,
			expected: ConstantBackoff{
				Delay:  time.Millisecond * 100,
				Jitter: 0.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewConstantBackoff(tt.duration, tt.jitter)
			if b.Delay != tt.expected.Delay {
				t.Errorf("Delay: got %v, want %v", b.Delay, tt.expected.Delay)
			}
			if b.Jitter != tt.expected.Jitter {
				t.Errorf("Jitter: got %v, want %v", b.Jitter, tt.expected.Jitter)
			}
		})
	}
}

func TestConstantBackoffDuration(t *testing.T) {
	b := NewConstantBackoff(time.Millisecond*500, 0)

	for attempt := 0; attempt <= 5; attempt++ {
		d := b.Duration(attempt)
		if d != time.Millisecond*500 {
			t.Errorf("attempt %d: got %v, want %v", attempt, d, time.Millisecond*500)
		}
	}
}

func TestConstantBackoffDurationWithJitter(t *testing.T) {
	b := NewConstantBackoff(time.Millisecond*500, 0.5)

	for attempt := 0; attempt <= 5; attempt++ {
		d := b.Duration(attempt)

		expected := float64(time.Millisecond * 500)
		minExpected := expected * (1 - 0.5)
		maxExpected := expected * (1 + 0.5)

		if float64(d) < minExpected || float64(d) > maxExpected {
			t.Errorf("attempt %d: duration %v not in range [%v, %v]", attempt, d, time.Duration(minExpected), time.Duration(maxExpected))
		}
	}
}

func TestBackoffInterface(t *testing.T) {
	var b Backoff

	b = NewExponentialBackoff(time.Millisecond, time.Second, 2, 0)
	if d := b.Duration(1); d == 0 {
		t.Error("ExponentialBackoff should implement Backoff interface")
	}

	b = NewLinearBackoff(time.Millisecond, time.Second, time.Millisecond*100, 0)
	if d := b.Duration(1); d == 0 {
		t.Error("LinearBackoff should implement Backoff interface")
	}

	b = NewConstantBackoff(time.Millisecond*100, 0)
	if d := b.Duration(1); d == 0 {
		t.Error("ConstantBackoff should implement Backoff interface")
	}
}
