package timex

import (
	"testing"
	"time"
)

func TestNewTokenRateLimiter(t *testing.T) {
	rl := NewTokenRateLimiter(100, time.Second)
	if rl.maxTokens != 100 {
		t.Errorf("maxTokens: got %d, want %d", rl.maxTokens, 100)
	}
	if rl.tokens != 100 {
		t.Errorf("tokens: got %d, want %d", rl.tokens, 100)
	}
	if rl.refillInterval != time.Second {
		t.Errorf("refillInterval: got %v, want %v", rl.refillInterval, time.Second)
	}
}

func TestTokenRateLimiterConsume(t *testing.T) {
	tests := []struct {
		name      string
		maxTokens int
		consume   int
		want      bool
	}{
		{"consume available", 10, 5, true},
		{"consume exactly max", 10, 10, true},
		{"consume exceeds max", 10, 11, false},
		{"consume zero", 10, 0, true},
		{"consume one from empty", 0, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := NewTokenRateLimiter(tt.maxTokens, time.Second)
			got := rl.Consume(tt.consume)
			if got != tt.want {
				t.Errorf("Consume(%d): got %v, want %v", tt.consume, got, tt.want)
			}
		})
	}
}

func TestTokenRateLimiterRefill(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Millisecond*50)

	if !rl.Consume(10) {
		t.Fatal("initial consume failed")
	}

	if rl.Consume(1) {
		t.Fatal("consume should fail when no tokens")
	}

	time.Sleep(time.Millisecond * 60)

	if !rl.Consume(1) {
		t.Fatal("consume should succeed after refill")
	}
}

func TestTokenRateLimiterMultipleRefills(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Millisecond*30)

	if !rl.Consume(10) {
		t.Fatal("initial consume failed")
	}

	time.Sleep(time.Millisecond * 100)

	if !rl.Consume(3) {
		t.Fatal("consume should succeed after 3+ refill intervals")
	}
}

func TestTokenRateLimiterRefillCappedAtMax(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Millisecond*30)

	if !rl.Consume(5) {
		t.Fatal("initial consume failed")
	}

	time.Sleep(time.Millisecond * 150)

	rl.Consume(0)

	if rl.tokens != 10 {
		t.Errorf("tokens should be capped at max: got %d, want %d", rl.tokens, 10)
	}
}

func TestTokenRateLimiterRelease(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Second)

	rl.Consume(8)

	if !rl.Release(3) {
		t.Fatal("Release should succeed")
	}

	if !rl.Consume(5) {
		t.Fatal("consume should succeed after release")
	}
}

func TestTokenRateLimiterReleaseFails(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Second)

	if rl.Release(11) {
		t.Error("Release should fail when exceeding max tokens")
	}
}

func TestTokenRateLimiterReleaseAll(t *testing.T) {
	rl := NewTokenRateLimiter(10, time.Second)

	rl.Consume(8)
	rl.ReleaseAll()

	if rl.tokens != 10 {
		t.Errorf("tokens after ReleaseAll: got %d, want %d", rl.tokens, 10)
	}
}

func TestNewTimeWindowRateLimiter(t *testing.T) {
	rl := NewTimeWindowRateLimiter(100, time.Minute)
	if rl.maxRequests != 100 {
		t.Errorf("maxRequests: got %d, want %d", rl.maxRequests, 100)
	}
	if rl.window != time.Minute {
		t.Errorf("window: got %v, want %v", rl.window, time.Minute)
	}
}

func TestTimeWindowRateLimiterConsume(t *testing.T) {
	tests := []struct {
		name        string
		maxRequests int
		consume     int
		want        bool
	}{
		{"consume available", 10, 5, true},
		{"consume exactly max", 10, 10, true},
		{"consume exceeds max", 10, 11, false},
		{"consume zero", 10, 0, true},
		{"consume one from zero max", 0, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := NewTimeWindowRateLimiter(tt.maxRequests, time.Minute)
			got := rl.Consume(tt.consume)
			if got != tt.want {
				t.Errorf("Consume(%d): got %v, want %v", tt.consume, got, tt.want)
			}
		})
	}
}

func TestTimeWindowRateLimiterMultipleConsumes(t *testing.T) {
	rl := NewTimeWindowRateLimiter(10, time.Minute)

	if !rl.Consume(3) {
		t.Fatal("first consume failed")
	}
	if !rl.Consume(4) {
		t.Fatal("second consume failed")
	}
	if rl.Consume(4) {
		t.Fatal("third consume should fail")
	}
}

func TestTimeWindowRateLimiterWindowReset(t *testing.T) {
	rl := NewTimeWindowRateLimiter(10, time.Millisecond*50)

	if !rl.Consume(10) {
		t.Fatal("initial consume failed")
	}

	if rl.Consume(1) {
		t.Fatal("consume should fail when at limit")
	}

	time.Sleep(time.Millisecond * 60)

	if !rl.Consume(5) {
		t.Fatal("consume should succeed after window resets")
	}
}

func TestTimeWindowRateLimiterRelease(t *testing.T) {
	rl := NewTimeWindowRateLimiter(10, time.Minute)

	rl.Consume(8)

	if !rl.Release(3) {
		t.Fatal("Release should succeed")
	}

	if !rl.Consume(5) {
		t.Fatal("consume should succeed after release")
	}
}

func TestTimeWindowRateLimiterReleaseFails(t *testing.T) {
	rl := NewTimeWindowRateLimiter(10, time.Minute)

	if rl.Release(1) {
		t.Error("Release should fail when count would go negative")
	}
}

func TestTimeWindowRateLimiterReleaseAll(t *testing.T) {
	rl := NewTimeWindowRateLimiter(10, time.Minute)

	rl.Consume(8)
	rl.ReleaseAll()

	if rl.requests != 0 {
		t.Errorf("requests after ReleaseAll: got %d, want %d", rl.requests, 0)
	}
}

func TestNewMultiRateLimiter(t *testing.T) {
	limiter1 := NewTokenRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	if len(rl.limiters) != 2 {
		t.Errorf("limiters count: got %d, want %d", len(rl.limiters), 2)
	}
}

func TestMultiRateLimiterConsumeBothAllow(t *testing.T) {
	limiter1 := NewTokenRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	if !rl.Consume(5) {
		t.Fatal("Consume should succeed when all limiters allow")
	}

	if limiter1.tokens != 5 {
		t.Errorf("limiter1 tokens: got %d, want %d", limiter1.tokens, 5)
	}
	if limiter2.requests != 5 {
		t.Errorf("limiter2 requests: got %d, want %d", limiter2.requests, 5)
	}
}

func TestMultiRateLimiterConsumeSingleFails(t *testing.T) {
	limiter1 := NewTokenRateLimiter(5, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	if rl.Consume(10) {
		t.Fatal("Consume should fail when any limiter rejects")
	}

	if limiter1.tokens != 5 {
		t.Errorf("limiter1 tokens should not change: got %d, want %d", limiter1.tokens, 5)
	}
	if limiter2.requests != 0 {
		t.Errorf("limiter2 requests should not change: got %d, want %d", limiter2.requests, 0)
	}
}

func TestMultiRateLimiterConsumerRollback(t *testing.T) {
	limiter1 := NewTokenRateLimiter(20, time.Second)
	limiter2 := NewTimeWindowRateLimiter(5, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	if rl.Consume(10) {
		t.Fatal("Consume should fail when second limiter rejects")
	}

	if limiter1.tokens != 20 {
		t.Errorf("limiter1 tokens should be rolled back: got %d, want %d", limiter1.tokens, 20)
	}
	if limiter2.requests != 0 {
		t.Errorf("limiter2 requests should not change: got %d, want %d", limiter2.requests, 0)
	}
}

func TestMultiRateLimiterRelease(t *testing.T) {
	limiter1 := NewTokenRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	rl.Consume(5)

	if !rl.Release(3) {
		t.Fatal("Release should succeed")
	}

	if limiter1.tokens != 8 {
		t.Errorf("limiter1 tokens: got %d, want %d", limiter1.tokens, 8)
	}
	if limiter2.requests != 2 {
		t.Errorf("limiter2 requests: got %d, want %d", limiter2.requests, 2)
	}
}

func TestMultiRateLimiterReleaseFails(t *testing.T) {
	limiter1 := NewTokenRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(5, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	got := rl.Release(1)
	if got {
		t.Fatal("Release should fail when any limiter rejects")
	}
}

func TestMultiRateLimiterReleaseAll(t *testing.T) {
	limiter1 := NewTokenRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	rl.Consume(5)
	rl.ReleaseAll()

	if limiter1.tokens != 10 {
		t.Errorf("limiter1 tokens: got %d, want %d", limiter1.tokens, 10)
	}
	if limiter2.requests != 0 {
		t.Errorf("limiter2 requests: got %d, want %d", limiter2.requests, 0)
	}
}

func TestMultiRateLimiterDualLimit(t *testing.T) {
	limiter1 := NewTimeWindowRateLimiter(10, time.Second)
	limiter2 := NewTimeWindowRateLimiter(100, time.Minute)

	rl := NewMultiRateLimiter(limiter1, limiter2)

	for i := 0; i < 10; i++ {
		if !rl.Consume(1) {
			t.Fatalf("consume %d failed", i+1)
		}
	}

	if rl.Consume(1) {
		t.Fatal("eleventh consume should fail (second limiter limit)")
	}

	if limiter1.requests != 10 {
		t.Errorf("limiter1 requests: got %d, want %d", limiter1.requests, 10)
	}
	if limiter2.requests != 10 {
		t.Errorf("limiter2 requests: got %d, want %d", limiter2.requests, 10)
	}
}

func TestRateLimiterInterface(t *testing.T) {
	var rl RateLimiter

	rl = NewTokenRateLimiter(10, time.Second)
	if !rl.Consume(5) {
		t.Error("TokenRateLimiter should implement RateLimiter interface")
	}

	rl = NewTimeWindowRateLimiter(10, time.Minute)
	if !rl.Consume(5) {
		t.Error("TimeWindowRateLimiter should implement RateLimiter interface")
	}

	rl = NewMultiRateLimiter(
		NewTokenRateLimiter(10, time.Second),
		NewTimeWindowRateLimiter(100, time.Minute),
	)
	if !rl.Consume(5) {
		t.Error("MultiRateLimiter should implement RateLimiter interface")
	}
}
