package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestRegexFilter(t *testing.T) {
	mock := &mockHandler{}
	filter, err := NewRegexFilter("error|fatal")
	if err != nil {
		t.Fatalf("RegexFilter: failed to create filter: %v", err)
	}
	filter.Next(mock)

	tests := []struct {
		msg      string
		wantPass bool
	}{
		{"this is an error", true},
		{"fatal exception occurred", true},
		{"info message", false},
		{"warning about error", true},
	}

	for _, tt := range tests {
		mock.records = nil
		record := slog.NewRecord(time.Now(), LevelInfo, tt.msg, 0)
		filter.Handle(context.Background(), record)

		if tt.wantPass && len(mock.records) == 0 {
			t.Errorf("RegexFilter: message %q should match", tt.msg)
		}
		if !tt.wantPass && len(mock.records) > 0 {
			t.Errorf("RegexFilter: message %q should not match", tt.msg)
		}
	}
}

func TestRegexFilterInvalidPattern(t *testing.T) {
	_, err := NewRegexFilter("[invalid(")
	if err == nil {
		t.Errorf("RegexFilter: should fail with invalid pattern")
	}
}
