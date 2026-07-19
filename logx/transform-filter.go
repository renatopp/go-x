package logx

import (
	"context"
	"log/slog"
)

type TransformFilter struct {
	fn   func(slog.Record) slog.Record
	next Handler
}

func NewTransformFilter(fn func(slog.Record) slog.Record) *TransformFilter {
	return &TransformFilter{
		fn: fn,
	}
}

func (f *TransformFilter) Next(next Handler) {
	f.next = next
}

func (f *TransformFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *TransformFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	transformed := f.fn(record)
	return f.next.Handle(ctx, transformed)
}

func (f *TransformFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &TransformFilter{
		fn:   f.fn,
		next: f.next.WithAttrs(attrs),
	}
}

func (f *TransformFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &TransformFilter{
		fn:   f.fn,
		next: f.next.WithGroup(name),
	}
}
