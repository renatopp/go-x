package logx

import (
	"context"
	"log/slog"

	"github.com/renatopp/go-x/timex"
)

type RateLimitFilter struct {
	limiter timex.RateLimiter
	next    Handler
}

func NewRateLimitFilter(limiter timex.RateLimiter) *RateLimitFilter {
	return &RateLimitFilter{
		limiter: limiter,
	}
}

func (f *RateLimitFilter) Next(next Handler) {
	f.next = next
}

func (f *RateLimitFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *RateLimitFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	if !f.limiter.Consume(1) {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *RateLimitFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &RateLimitFilter{
		limiter: f.limiter,
		next:    f.next.WithAttrs(attrs),
	}
}

func (f *RateLimitFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &RateLimitFilter{
		limiter: f.limiter,
		next:    f.next.WithGroup(name),
	}
}
