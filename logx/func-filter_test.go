package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestFuncFilter(t *testing.T) {
	mock := &mockHandler{}

	filterFn := func(ctx context.Context, record slog.Record) bool {
		return record.Level >= LevelWarn
	}
	filter := NewFuncFilter(filterFn)
	filter.Next(mock)

	tests := []struct {
		level    Level
		msg      string
		wantPass bool
	}{
		{LevelDebug, "debug", false},
		{LevelInfo, "info", false},
		{LevelWarn, "warn", true},
		{LevelError, "error", true},
	}

	for _, tt := range tests {
		mock.records = nil
		record := slog.NewRecord(time.Now(), tt.level, tt.msg, 0)
		filter.Handle(context.Background(), record)

		if tt.wantPass && len(mock.records) == 0 {
			t.Errorf("FuncFilter: level %v should pass", tt.level)
		}
		if !tt.wantPass && len(mock.records) > 0 {
			t.Errorf("FuncFilter: level %v should not pass", tt.level)
		}
	}
}
