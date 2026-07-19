package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestTransformFilter(t *testing.T) {
	mock := &mockHandler{}
	transformFn := func(record slog.Record) slog.Record {
		record.Message = "transformed: " + record.Message
		return record
	}
	filter := NewTransformFilter(transformFn)
	filter.Next(mock)

	record := slog.NewRecord(time.Now(), LevelInfo, "original", 0)
	filter.Handle(context.Background(), record)

	if len(mock.records) != 1 {
		t.Errorf("TransformFilter: record should pass")
	}
	if mock.records[0].Message != "transformed: original" {
		t.Errorf("TransformFilter: expected message 'transformed: original', got %q", mock.records[0].Message)
	}
}
