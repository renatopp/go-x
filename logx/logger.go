package logx

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	mu          sync.Mutex
	level       Level
	callerInfo  bool
	skipCallers int
	handler     Handler
}

func NewLogger(handler Handler) *Logger {
	return &Logger{
		level:       LevelInfo,
		callerInfo:  false,
		skipCallers: 0,
		handler:     handler,
	}
}

func (l *Logger) Clone() *Logger {
	return &Logger{
		level:       l.level,
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
	return l.handler
}

func (l *Logger) WithLevel(level Level) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
	return l
}

func (l *Logger) WithCallerInfo(v bool) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.callerInfo = v
	return l
}

func (l *Logger) WithAttributes(attrs ...any) *Logger {
	if l.handler == nil {
		return l
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handler = l.handler.WithAttrs(argsToAttrSlice(attrs))
	return l
}

func (l *Logger) WithGroup(name string) *Logger {
	if l.handler == nil {
		return l
	}
	l.mu.Lock()
	defer l.mu.Unlock()
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

func (l *Logger) Fatal(msg string, kvargs ...any) {
	l.Log(LevelFatal, msg, kvargs...)
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

func (l *Logger) Fatalf(msg string, args ...any) {
	l.Logf(LevelFatal, msg, args...)
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

func (l *Logger) Fatalc(ctx context.Context, msg string, kvargs ...any) {
	l.Logc(ctx, LevelFatal, msg, kvargs...)
}

func (l *Logger) log(ctx context.Context, level Level, msg string, kvargs ...any) {
	if l.handler == nil {
		return
	}

	var pc uintptr
	if l.callerInfo {
		var pcs [1]uintptr
		// skip [runtime.Callers, this function, this function's caller]
		runtime.Callers(3, pcs[l.skipCallers:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(kvargs...)
	l.handler.Handle(ctx, r)
}
