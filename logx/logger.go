package logx

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"
)

// exitFunc is called after a Fatal log is emitted. It is a variable so tests
// can stub it out instead of terminating the test process.
var exitFunc = os.Exit

type Logger struct {
	mu          sync.Mutex
	level       Level
	timestamp   bool
	callerInfo  bool
	skipCallers int
	handler     Handler
}

func NewLogger(handler Handler) *Logger {
	return &Logger{
		level:       LevelInfo,
		timestamp:   false,
		callerInfo:  false,
		skipCallers: 0,
		handler:     handler,
	}
}

func (l *Logger) Clone() *Logger {
	return &Logger{
		level:       l.level,
		timestamp:   l.timestamp,
		callerInfo:  l.callerInfo,
		skipCallers: l.skipCallers,
		handler:     l.handler,
	}
}

func (l *Logger) Level() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

func (l *Logger) Handler() Handler {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.handler
}

func (l *Logger) WithLevel(level Level) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
	return l
}

func (l *Logger) WithTimestamp(v bool) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = v
	return l
}

func (l *Logger) WithCallerInfo(v bool) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.callerInfo = v
	return l
}

func (l *Logger) WithSkipCallers(n int) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.skipCallers = n
	return l
}

func (l *Logger) WithAttributes(attrs ...any) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handler == nil {
		return l
	}
	l.handler = l.handler.WithAttrs(argsToAttrSlice(attrs))
	return l
}

func (l *Logger) WithGroup(name string) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handler == nil {
		return l
	}
	l.handler = l.handler.WithGroup(name)
	return l
}

func (l *Logger) Log(level Level, msg string, kvargs ...any) {
	l.log(context.Background(), level, msg, kvargs...)
}

func (l *Logger) Debug(msg string, kvargs ...any) {
	l.Log(LevelDebug, msg, kvargs...)
}

func (l *Logger) Info(msg string, kvargs ...any) {
	l.Log(LevelInfo, msg, kvargs...)
}

func (l *Logger) Warn(msg string, kvargs ...any) {
	l.Log(LevelWarn, msg, kvargs...)
}

func (l *Logger) Error(msg string, kvargs ...any) {
	l.Log(LevelError, msg, kvargs...)
}

// Fatal logs at LevelFatal and then terminates the process via os.Exit(1).
func (l *Logger) Fatal(msg string, kvargs ...any) {
	l.Log(LevelFatal, msg, kvargs...)
	exitFunc(1)
}

func (l *Logger) Logf(level Level, msg string, args ...any) {
	l.Log(level, fmt.Sprintf(msg, args...))
}

func (l *Logger) Debugf(msg string, args ...any) {
	l.Logf(LevelDebug, msg, args...)
}

func (l *Logger) Infof(msg string, args ...any) {
	l.Logf(LevelInfo, msg, args...)
}

func (l *Logger) Warnf(msg string, args ...any) {
	l.Logf(LevelWarn, msg, args...)
}

func (l *Logger) Errorf(msg string, args ...any) {
	l.Logf(LevelError, msg, args...)
}

// Fatalf logs at LevelFatal and then terminates the process via os.Exit(1).
func (l *Logger) Fatalf(msg string, args ...any) {
	l.Logf(LevelFatal, msg, args...)
	exitFunc(1)
}

func (l *Logger) Logc(ctx context.Context, level Level, msg string, kvargs ...any) {
	l.log(ctx, level, msg, kvargs...)
}

func (l *Logger) Debugc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelDebug, msg, kvargs...)
}

func (l *Logger) Infoc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelInfo, msg, kvargs...)
}

func (l *Logger) Warnc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelWarn, msg, kvargs...)
}

func (l *Logger) Errorc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelError, msg, kvargs...)
}

// Fatalc logs at LevelFatal and then terminates the process via os.Exit(1).
func (l *Logger) Fatalc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelFatal, msg, kvargs...)
	exitFunc(1)
}

func (l *Logger) log(ctx context.Context, level Level, msg string, kvargs ...any) {
	l.mu.Lock()
	handler := l.handler
	minLevel := l.level
	timestamp := l.timestamp
	callerInfo := l.callerInfo
	skipCallers := l.skipCallers
	l.mu.Unlock()

	if handler == nil || level < minLevel {
		return
	}

	var pc uintptr
	if callerInfo {
		pcs := make([]uintptr, skipCallers+1)
		// skip [runtime.Callers, this function, this function's caller]
		runtime.Callers(3, pcs)
		pc = pcs[skipCallers]
	}

	var t time.Time
	if timestamp {
		t = time.Now()
	}
	r := slog.NewRecord(t, level, msg, pc)
	r.Add(kvargs...)
	handler.Handle(ctx, r)
}
