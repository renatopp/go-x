package logx

import (
	"context"
	"log/slog"
)

type filterWrapper struct {
	next Handler
}

func AsFilter(h Handler) Filter {
	if f, ok := h.(Filter); ok {
		return f
	}
	return &filterWrapper{next: h}
}

func (f *filterWrapper) Next(next Handler) {
	f.next = next
}

func (f *filterWrapper) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *filterWrapper) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *filterWrapper) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &filterWrapper{
		next: f.next.WithAttrs(attrs),
	}
}

func (f *filterWrapper) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &filterWrapper{
		next: f.next.WithGroup(name),
	}
}
