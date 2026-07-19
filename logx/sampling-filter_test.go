package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestSamplingFilter(t *testing.T) {
	mock := &mockHandler{}
	filter := NewSamplingFilter(3)
	filter.Next(mock)

	record := slog.NewRecord(time.Now(), LevelInfo, "test", 0)

	for i := 0; i < 9; i++ {
		mock.records = nil
		filter.Handle(context.Background(), record)
		if (i+1)%3 == 1 && len(mock.records) != 1 {
			t.Errorf("SamplingFilter: record %d should pass (1 in 3)", i+1)
		}
		if (i+1)%3 != 1 && len(mock.records) != 0 {
			t.Errorf("SamplingFilter: record %d should be sampled out", i+1)
		}
	}
}
