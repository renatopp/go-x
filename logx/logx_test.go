package logx

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestPipelineEmpty(t *testing.T) {
	if Pipeline() != nil {
		t.Errorf("expected Pipeline() with no handlers to return nil")
	}
}

// TestPipelineNonFilterHandlerInMiddle guards against a regression where a
// plain (non-Filter) Handler placed anywhere but last in the chain was
// silently dropped: AsFilter wrapped it in a filterWrapper, but Pipeline then
// called Next on that same wrapper to attach the following link, overwriting
// the reference to the wrapped handler.
func TestPipelineNonFilterHandlerInMiddle(t *testing.T) {
	middle := &mockHandler{}
	last := &mockHandler{}

	pipeline := Pipeline(NewLevelFilter(LevelInfo), middle, last)

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	pipeline.Handle(context.Background(), record)

	if len(middle.records) != 1 {
		t.Errorf("expected the middle non-Filter handler to receive the record, got %d", len(middle.records))
	}
	if len(last.records) != 1 {
		t.Errorf("expected the last handler to also receive the record, got %d", len(last.records))
	}
}

func TestPipelineMultipleNonFilterHandlersInARow(t *testing.T) {
	first := &mockHandler{}
	second := &mockHandler{}
	third := &mockHandler{}

	pipeline := Pipeline(first, second, third)

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	pipeline.Handle(context.Background(), record)

	for i, h := range []*mockHandler{first, second, third} {
		if len(h.records) != 1 {
			t.Errorf("handler %d: expected to receive the record, got %d", i, len(h.records))
		}
	}
}

func TestPipelineWithAttrsPropagatesThroughWrappedHandlers(t *testing.T) {
	middle := &mockHandler{}
	last := &mockHandler{}

	pipeline := Pipeline(middle, last).WithAttrs([]slog.Attr{String("k", "v")})

	record := slog.NewRecord(time.Now(), LevelInfo, "hello", 0)
	pipeline.Handle(context.Background(), record)

	if len(middle.records) != 1 {
		t.Errorf("expected the middle handler to still receive the record after WithAttrs, got %d", len(middle.records))
	}
	if len(last.records) != 1 {
		t.Errorf("expected the last handler to still receive the record after WithAttrs, got %d", len(last.records))
	}
}

func TestArgsToAttrSlice(t *testing.T) {
	attrs := argsToAttrSlice([]any{"a", 1, "b", "two", String("c", "three")})
	if len(attrs) != 3 {
		t.Fatalf("expected 3 attrs, got %d", len(attrs))
	}
	if attrs[0].Key != "a" || attrs[0].Value.Int64() != 1 {
		t.Errorf("unexpected attr[0]: %+v", attrs[0])
	}
	if attrs[1].Key != "b" || attrs[1].Value.String() != "two" {
		t.Errorf("unexpected attr[1]: %+v", attrs[1])
	}
	if attrs[2].Key != "c" || attrs[2].Value.String() != "three" {
		t.Errorf("unexpected attr[2]: %+v", attrs[2])
	}
}

func TestArgsToAttrSliceDanglingKey(t *testing.T) {
	attrs := argsToAttrSlice([]any{"lonely"})
	if len(attrs) != 1 || attrs[0].Key != badKey {
		t.Errorf("expected a single badKey attr for a dangling key, got %+v", attrs)
	}
}

func TestAttrConstructors(t *testing.T) {
	if a := Int("i", 1); a.Value.Kind() != slog.KindInt64 {
		t.Errorf("Int: expected KindInt64, got %v", a.Value.Kind())
	}
	if a := Bool("b", true); a.Value.Kind() != slog.KindBool {
		t.Errorf("Bool: expected KindBool, got %v", a.Value.Kind())
	}
	if a := Float64("f", 1.5); a.Value.Kind() != slog.KindFloat64 {
		t.Errorf("Float64: expected KindFloat64, got %v", a.Value.Kind())
	}
	if a := Duration("d", time.Second); a.Value.Kind() != slog.KindDuration {
		t.Errorf("Duration: expected KindDuration, got %v", a.Value.Kind())
	}
	group := Group("g", String("k", "v"))
	if group.Value.Kind() != slog.KindGroup || len(group.Value.Group()) != 1 {
		t.Errorf("Group: expected a single-element group, got %+v", group)
	}
}
