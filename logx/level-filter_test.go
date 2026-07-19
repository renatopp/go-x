package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestLevelFilter(t *testing.T) {
	mock := &mockHandler{}
	filter := NewLevelFilter(LevelWarn)
	filter.Next(mock)

	tests := []struct {
		level    Level
		msg      string
		wantPass bool
	}{
		{LevelDebug, "debug msg", false},
		{LevelInfo, "info msg", false},
		{LevelWarn, "warn msg", true},
		{LevelError, "error msg", true},
	}

	for _, tt := range tests {
		mock.records = nil
		record := slog.NewRecord(time.Now(), tt.level, tt.msg, 0)
		filter.Handle(context.Background(), record)

		if tt.wantPass && len(mock.records) == 0 {
			t.Errorf("LevelFilter: level %v should pass but didn't", tt.level)
		}
		if !tt.wantPass && len(mock.records) > 0 {
			t.Errorf("LevelFilter: level %v should not pass but did", tt.level)
		}
	}
}
