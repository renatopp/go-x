package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/renatopp/go-x/timex"
)

func TestRateLimitFilter(t *testing.T) {
	mock := &mockHandler{}
	limiter := timex.NewTokenRateLimiter(2, 10*time.Millisecond)
	filter := NewRateLimitFilter(limiter)
	filter.Next(mock)

	record := slog.NewRecord(time.Now(), LevelInfo, "test", 0)

	filter.Handle(context.Background(), record)
	if len(mock.records) != 1 {
		t.Errorf("RateLimitFilter: first record should pass")
	}

	mock.records = nil
	filter.Handle(context.Background(), record)
	if len(mock.records) != 1 {
		t.Errorf("RateLimitFilter: second record should pass")
	}

	mock.records = nil
	filter.Handle(context.Background(), record)
	if len(mock.records) != 0 {
		t.Errorf("RateLimitFilter: third record should be rate limited")
	}

	time.Sleep(15 * time.Millisecond)
	filter.Handle(context.Background(), record)
	if len(mock.records) != 1 {
		t.Errorf("RateLimitFilter: record should pass after refill")
	}
}
