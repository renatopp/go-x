package logx

import (
	"context"
	"log/slog"
)

type FuncFilter struct {
	fn   func(ctx context.Context, record slog.Record) bool
	next Handler
}

func NewFuncFilter(fn func(ctx context.Context, record slog.Record) bool) *FuncFilter {
	return &FuncFilter{
		fn: fn,
	}
}

func (f *FuncFilter) Next(next Handler) {
	f.next = next
}

func (f *FuncFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *FuncFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	if !f.fn(ctx, record) {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *FuncFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &FuncFilter{
		fn:   f.fn,
		next: f.next.WithAttrs(attrs),
	}
}

func (f *FuncFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &FuncFilter{
		fn:   f.fn,
		next: f.next.WithGroup(name),
	}
}
