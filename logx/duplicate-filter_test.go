package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestDuplicateFilter(t *testing.T) {
	mock := &mockHandler{}
	filter := NewDuplicateFilter()
	filter.Next(mock)

	record1 := slog.NewRecord(time.Now(), LevelInfo, "msg1", 0)
	record2 := slog.NewRecord(time.Now(), LevelInfo, "msg1", 0)
	record3 := slog.NewRecord(time.Now(), LevelInfo, "msg2", 0)
	record4 := slog.NewRecord(time.Now(), LevelWarn, "msg1", 0)

	filter.Handle(context.Background(), record1)
	if len(mock.records) != 1 {
		t.Errorf("DuplicateFilter: first record should pass")
	}

	mock.records = nil
	filter.Handle(context.Background(), record2)
	if len(mock.records) != 0 {
		t.Errorf("DuplicateFilter: duplicate message at same level should be dropped")
	}

	filter.Handle(context.Background(), record3)
	if len(mock.records) != 1 {
		t.Errorf("DuplicateFilter: different message should pass")
	}

	mock.records = nil
	filter.Handle(context.Background(), record4)
	if len(mock.records) != 1 {
		t.Errorf("DuplicateFilter: same message at different level should pass")
	}
}
