package logx

import (
	"bytes"
	"context"
	"log/slog"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestColoredHandlerHandleWritesMessage(t *testing.T) {
	var buf bytes.Buffer
	h := NewColoredHandler().WithWriter(&buf)

	record := slog.NewRecord(time.Now(), LevelInfo, "hello world", 0)
	h.Handle(context.Background(), record)

	if !strings.Contains(buf.String(), "hello world") {
		t.Errorf("expected message in output, got %q", buf.String())
	}
}

// TestColoredHandlerBoundAndRecordAttrsBothAppear guards the fix that made
// Handle render both attrs bound via WithAttrs and attrs attached directly
// to the record; previously only the latter were rendered.
func TestColoredHandlerBoundAndRecordAttrsBothAppear(t *testing.T) {
	var buf bytes.Buffer
	base := NewColoredHandler().WithWriter(&buf)
	h := base.WithAttrs([]slog.Attr{String("bound", "1")})

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	record.AddAttrs(String("live", "2"))
	h.Handle(context.Background(), record)

	out := buf.String()
	if !strings.Contains(out, "bound=1") {
		t.Errorf("expected bound attr in output, got %q", out)
	}
	if !strings.Contains(out, "live=2") {
		t.Errorf("expected record attr in output, got %q", out)
	}
}

func TestColoredHandlerWithGroupPrefixesAttrKeys(t *testing.T) {
	var buf bytes.Buffer
	h := NewColoredHandler().WithWriter(&buf).WithGroup("g")

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	record.AddAttrs(String("k", "v"))
	h.Handle(context.Background(), record)

	if !strings.Contains(buf.String(), "g.k=v") {
		t.Errorf("expected grouped attr key g.k=v in output, got %q", buf.String())
	}
}

func TestColoredHandlerWithAttrsIsImmutable(t *testing.T) {
	var buf bytes.Buffer
	base := NewColoredHandler().WithWriter(&buf)
	_ = base.WithAttrs([]slog.Attr{String("k", "v")})

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	base.Handle(context.Background(), record)

	if strings.Contains(buf.String(), "k=v") {
		t.Errorf("expected base handler to be unaffected by WithAttrs, got %q", buf.String())
	}
}

func TestColoredHandlerNilLevelFormatterOmitsLevel(t *testing.T) {
	var buf bytes.Buffer
	h := NewColoredHandler().WithWriter(&buf).WithLevelFormatter(nil)

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	h.Handle(context.Background(), record)

	if strings.Contains(buf.String(), "INFO") {
		t.Errorf("expected level to be omitted, got %q", buf.String())
	}
}

func TestFullColoredHandlerIncludesTimestampAndCaller(t *testing.T) {
	var buf bytes.Buffer
	h := NewFullColoredHandler().WithWriter(&buf)

	pcs := make([]uintptr, 1)
	runtime.Callers(1, pcs)
	record := slog.NewRecord(time.Now(), LevelInfo, "hello", pcs[0])
	h.Handle(context.Background(), record)

	out := buf.String()
	if !strings.Contains(out, "hello") {
		t.Errorf("expected message in output, got %q", out)
	}
	if !strings.Contains(out, "colored-handler_test.go") {
		t.Errorf("expected caller info in output, got %q", out)
	}
}
