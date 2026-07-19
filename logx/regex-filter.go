package logx

import (
	"context"
	"log/slog"
	"regexp"
)

type RegexFilter struct {
	pattern *regexp.Regexp
	next    Handler
}

func NewRegexFilter(pattern string) (*RegexFilter, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexFilter{
		pattern: re,
	}, nil
}

func (f *RegexFilter) Next(next Handler) {
	f.next = next
}

func (f *RegexFilter) Enabled(ctx context.Context, level slog.Level) bool {
	if f.next == nil {
		return false
	}
	return f.next.Enabled(ctx, level)
}

func (f *RegexFilter) Handle(ctx context.Context, record slog.Record) error {
	if f.next == nil {
		return nil
	}
	if !f.pattern.MatchString(record.Message) {
		return nil
	}
	return f.next.Handle(ctx, record)
}

func (f *RegexFilter) WithAttrs(attrs []slog.Attr) Handler {
	if f.next == nil {
		return f
	}
	return &RegexFilter{
		pattern: f.pattern,
		next:    f.next.WithAttrs(attrs),
	}
}

func (f *RegexFilter) WithGroup(name string) Handler {
	if f.next == nil {
		return f
	}
	return &RegexFilter{
		pattern: f.pattern,
		next:    f.next.WithGroup(name),
	}
}
