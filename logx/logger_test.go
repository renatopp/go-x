package logx

import (
	"context"
	"sync"
	"testing"
)

// noopHandler is a Handler with no shared state, used for concurrency tests
// where we only care about Logger's own field safety, not about recording
// what was logged.
type noopHandler struct{}

func (noopHandler) Enabled(context.Context, Level) bool  { return true }
func (noopHandler) Handle(context.Context, Record) error { return nil }
func (noopHandler) WithAttrs(attrs []Attr) Handler       { return noopHandler{} }
func (noopHandler) WithGroup(name string) Handler        { return noopHandler{} }

func TestLoggerDefaultLevelFiltersDebug(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock)

	l.Debug("debug")
	if len(mock.records) != 0 {
		t.Errorf("expected Debug to be filtered out by default LevelInfo")
	}

	l.Info("info")
	if len(mock.records) != 1 {
		t.Errorf("expected Info to pass by default")
	}
}

func TestLoggerWithLevelFilters(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock).WithLevel(LevelWarn)

	l.Debug("debug")
	l.Info("info")
	if len(mock.records) != 0 {
		t.Errorf("expected Debug and Info to be filtered by WithLevel(LevelWarn), got %d records", len(mock.records))
	}

	l.Warn("warn")
	if len(mock.records) != 1 {
		t.Errorf("expected Warn to pass WithLevel(LevelWarn)")
	}
}

func TestLoggerSkipCallersDoesNotPanic(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock).WithCallerInfo(true)

	for _, n := range []int{0, 1, 2, 5, 50} {
		mock.records = nil
		l.WithSkipCallers(n)
		l.Info("test")
		if len(mock.records) != 1 {
			t.Errorf("skipCallers=%d: expected record to be logged, got %d records", n, len(mock.records))
		}
	}
}

func TestLoggerTimestamp(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock)

	l.Info("no ts")
	if !mock.records[0].Time.IsZero() {
		t.Errorf("expected zero time when timestamp is disabled")
	}

	mock.records = nil
	l.WithTimestamp(true)
	l.Info("with ts")
	if mock.records[0].Time.IsZero() {
		t.Errorf("expected non-zero time when timestamp is enabled")
	}
}

func TestLoggerClone(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock).WithLevel(LevelError).WithTimestamp(true).WithCallerInfo(true).WithSkipCallers(2)

	clone := l.Clone()
	if clone.Level() != LevelError {
		t.Errorf("Clone: expected level %v, got %v", LevelError, clone.Level())
	}
	if !clone.timestamp {
		t.Errorf("Clone: expected timestamp to be copied as true")
	}
	if !clone.callerInfo {
		t.Errorf("Clone: expected callerInfo to be copied as true")
	}
	if clone.skipCallers != 2 {
		t.Errorf("Clone: expected skipCallers 2, got %d", clone.skipCallers)
	}
	if clone.Handler() != mock {
		t.Errorf("Clone: expected handler to be copied")
	}
}

func TestLoggerLogf(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock)

	l.Infof("hello %s, you are %d", "world", 42)
	if len(mock.records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(mock.records))
	}
	if got := mock.records[0].Message; got != "hello world, you are 42" {
		t.Errorf("unexpected message: %q", got)
	}
}

func TestLoggerLogc(t *testing.T) {
	mock := &mockHandler{}
	l := NewLogger(mock)

	l.Infoc(context.Background(), "with ctx")
	if len(mock.records) != 1 || mock.records[0].Level != LevelInfo {
		t.Errorf("expected Infoc to log a single record at LevelInfo")
	}
}

func TestLoggerFatalVariantsExit(t *testing.T) {
	origExit := exitFunc
	defer func() { exitFunc = origExit }()

	tests := []struct {
		name string
		call func(l *Logger)
	}{
		{"Fatal", func(l *Logger) { l.Fatal("boom") }},
		{"Fatalf", func(l *Logger) { l.Fatalf("boom %d", 1) }},
		{"Fatalc", func(l *Logger) { l.Fatalc(context.Background(), "boom") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockHandler{}
			l := NewLogger(mock)

			var exitCode int
			var called bool
			exitFunc = func(code int) { called = true; exitCode = code }

			tt.call(l)

			if !called {
				t.Errorf("%s: expected exitFunc to be called", tt.name)
			}
			if exitCode != 1 {
				t.Errorf("%s: expected exit code 1, got %d", tt.name, exitCode)
			}
			if len(mock.records) != 1 || mock.records[0].Level != LevelFatal {
				t.Errorf("%s: expected a single LevelFatal record to be logged", tt.name)
			}
		})
	}
}

func TestLoggerWithAttributesNilHandler(t *testing.T) {
	l := NewLogger(nil)
	if got := l.WithAttributes("a", 1); got != l {
		t.Errorf("expected same logger to be returned when handler is nil")
	}
	if got := l.WithGroup("g"); got != l {
		t.Errorf("expected same logger to be returned when handler is nil")
	}
}

func TestLoggerConcurrentAccess(t *testing.T) {
	l := NewLogger(noopHandler{})

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(3)
		go func() { defer wg.Done(); l.Info("x") }()
		go func() { defer wg.Done(); l.WithTimestamp(true) }()
		go func() { defer wg.Done(); l.WithAttributes("k", 1) }()
	}
	wg.Wait()
}
