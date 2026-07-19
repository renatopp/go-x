package logx

import (
	"context"
	"log/slog"
	"sync/atomic"
)

type samplingControl struct {
	n       uint64
	counter atomic.Uint64
}

type SamplingFilter struct {
	control *samplingControl
	next    Handler
}

// SamplingFilter is a Filter that only passes every n-th log record to the
// next Handler. It can be used to reduce the volume of logs in high-frequency
// scenarios.
func NewSamplingFilter(n int) *SamplingFilter {
	if n < 1 {
		n = 1
	}
	return &SamplingFilter{
		control: &samplingControl{
			n: uint64(n),
		},
	}
}

func (f *SamplingFilter) Next(next Handler) {
	f.next = next
}

func (f *SamplingFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *SamplingFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	if f.control.counter.Add(1)%f.control.n != 1 {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *SamplingFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &SamplingFilter{
		control: f.control,
		next:    f.next.WithAttrs(attrs),
	}
}

func (f *SamplingFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &SamplingFilter{
		control: f.control,
		next:    f.next.WithGroup(name),
	}
}
