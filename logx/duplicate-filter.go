package logx

import (
	"context"
	"log/slog"
	"sync"
)

type duplicateControl struct {
	mu      sync.Mutex
	lastMsg string
	lastLvl slog.Level
	hasLast bool
}

type DuplicateFilter struct {
	control *duplicateControl
	next    Handler
}

func NewDuplicateFilter() *DuplicateFilter {
	return &DuplicateFilter{
		control: &duplicateControl{},
	}
}

func (f *DuplicateFilter) Next(next Handler) {
	f.next = next
}

func (f *DuplicateFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *DuplicateFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}

	f.control.mu.Lock()
	isDuplicate := f.control.hasLast && f.control.lastMsg == record.Message && f.control.lastLvl == record.Level
	f.control.lastMsg = record.Message
	f.control.lastLvl = record.Level
	f.control.hasLast = true
	f.control.mu.Unlock()

	if isDuplicate {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *DuplicateFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &DuplicateFilter{
		control: f.control,
		next:    f.next.WithAttrs(attrs),
	}
}

func (f *DuplicateFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &DuplicateFilter{
		control: f.control,
		next:    f.next.WithGroup(name),
	}
}
