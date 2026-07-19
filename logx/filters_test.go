package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

type mockHandler struct {
	records []slog.Record
}

func (m *mockHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (m *mockHandler) Handle(ctx context.Context, record slog.Record) error {
	m.records = append(m.records, record)
	return nil
}

func (m *mockHandler) WithAttrs(attrs []slog.Attr) Handler {
	return m
}

func (m *mockHandler) WithGroup(name string) Handler {
	return m
}

func TestPipeline(t *testing.T) {
	mock := &mockHandler{}

	levelFilter := NewLevelFilter(LevelWarn)
	regexFilter, _ := NewRegexFilter("error")

	pipeline := Pipeline(levelFilter, regexFilter, mock)

	tests := []struct {
		level    Level
		msg      string
		wantPass bool
	}{
		{LevelDebug, "error occurred", false},
		{LevelWarn, "error occurred", true},
		{LevelWarn, "warning message", false},
		{LevelError, "error occurred", true},
	}

	for _, tt := range tests {
		mock.records = nil
		record := slog.NewRecord(time.Now(), tt.level, tt.msg, 0)
		pipeline.Handle(context.Background(), record)

		if tt.wantPass && len(mock.records) == 0 {
			t.Errorf("Pipeline: level %v, msg %q should pass", tt.level, tt.msg)
		}
		if !tt.wantPass && len(mock.records) > 0 {
			t.Errorf("Pipeline: level %v, msg %q should not pass", tt.level, tt.msg)
		}
	}
}
