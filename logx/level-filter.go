package logx

import (
	"context"
	"log/slog"
)

type LevelFilter struct {
	Level Level
	next  Handler
}

func NewLevelFilter(level Level) *LevelFilter {
	return &LevelFilter{
		Level: level,
	}
}

func (f *LevelFilter) Next(next Handler) {
	f.next = next
}

func (f *LevelFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	if level < f.Level {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *LevelFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	if record.Level < f.Level {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *LevelFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &LevelFilter{
		Level: f.Level,
		next:  f.next.WithAttrs(attrs),
	}
}

func (f *LevelFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &LevelFilter{
		Level: f.Level,
		next:  f.next.WithGroup(name),
	}
}
