package logx

import (
	"context"
	"log/slog"
)

// filterWrapper adapts a plain Handler (one that doesn't implement Filter) so
// it can take part in a Pipeline. wrapped is the handler being adapted; next
// is the following handler in the chain, attached later via Next. Keeping
// these separate matters: Pipeline always calls Next on every link it
// creates, and if that reused the same field as the wrapped handler, the
// wrapped handler would be overwritten and silently dropped.
type filterWrapper struct {
	wrapped Handler
	next    Handler
}

func AsFilter(h Handler) Filter {
	if f, ok := h.(Filter); ok {
		return f
	}
	return &filterWrapper{wrapped: h}
}

func (f *filterWrapper) Next(next Handler) {
	f.next = next
}

func (f *filterWrapper) Enabled(ctx context.Context, level slog.Level) bool {
	return f.wrapped.Enabled(ctx, level)
}

func (f *filterWrapper) Handle(ctx context.Context, record slog.Record) error {
	if err := f.wrapped.Handle(ctx, record); err != nil {
		return err
	}
	if f.next == nil {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *filterWrapper) WithAttrs(attrs []slog.Attr) Handler {
	var next Handler
	if f.next != nil {
		next = f.next.WithAttrs(attrs)
	}
	return &filterWrapper{
		wrapped: f.wrapped.WithAttrs(attrs),
		next:    next,
	}
}

func (f *filterWrapper) WithGroup(name string) Handler {
	var next Handler
	if f.next != nil {
		next = f.next.WithGroup(name)
	}
	return &filterWrapper{
		wrapped: f.wrapped.WithGroup(name),
		next:    next,
	}
}
