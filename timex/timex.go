package timex

import (
	"context"
	"sync"
	"time"
)

// TimeProvider returns the current time. Clock uses it instead of calling
// time.Now() directly so tests can substitute a deterministic clock.
type TimeProvider func() time.Time

// SystemProvider returns a TimeProvider that returns the current system time.
func SystemProvider() TimeProvider {
	return time.Now
}

// ConstantProvider returns a TimeProvider that always returns the same time.
func ConstantProvider(t time.Time) TimeProvider {
	return func() time.Time { return t }
}

// MockProvider returns a TimeProvider that calls f with the number of times
// it has been called and returns the time returned by f. It is safe for
// concurrent use.
func MockProvider(f func(i int) time.Time) TimeProvider {
	var (
		mu    sync.Mutex
		count int
	)
	return func() time.Time {
		mu.Lock()
		defer mu.Unlock()
		count++
		return f(count)
	}
}

// Debounce returns a function that will call f after d duration has passed
// since the last call to the returned function. If the returned function is
// called again before d duration has passed, the timer is reset and f will
// not be called until d duration has passed since the last call.
func Debounce(d time.Duration, f func()) func() {
	return DebounceCtx(context.Background(), d, f)
}

// DebounceCtx is like Debounce but respects the context. If the context is
// canceled before the duration has passed, f will not be called.
func DebounceCtx(ctx context.Context, d time.Duration, f func()) func() {
	var (
		mu    sync.Mutex
		timer *time.Timer
	)
	return func() {
		mu.Lock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(d, func() {
			select {
			case <-ctx.Done():
				return
			default:
				f()
			}
		})
		mu.Unlock()
	}
}

// Throttle returns a function that will call f at most once every d duration.
// If the returned function is called again before d duration has passed, f will
// not be called until d duration has passed since the last call.
func Throttle(d time.Duration, f func()) func() {
	return ThrottleCtx(context.Background(), d, f)
}

// ThrottleCtx is like Throttle but respects the context. If the context is
// canceled before the duration has passed, f will not be called.
func ThrottleCtx(ctx context.Context, d time.Duration, f func()) func() {
	var (
		mu       sync.Mutex
		lastCall time.Time
	)
	return func() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		mu.Lock()
		now := time.Now()
		if now.Sub(lastCall) >= d {
			lastCall = now
			mu.Unlock()
			f()
		} else {
			mu.Unlock()
		}
	}
}
