package logx

import "time"

// TimeProvider returns the current time. Logger uses it instead of calling
// time.Now() directly so tests can substitute a deterministic clock.
type TimeProvider func() time.Time

// defaultTimeProvider is the TimeProvider used when none is configured.
func defaultTimeProvider() time.Time {
	return time.Now()
}
