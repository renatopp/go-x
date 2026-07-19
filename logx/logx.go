// Package logx extends log/slog with a configurable Logger: levels,
// prefixes, attributes, groups, pluggable formatters, optional
// timestamp/caller/stacktrace capture, sampling, and rate limiting.
package logx

import (
	"context"
	"sync/atomic"
)

var defaultLogger atomic.Pointer[Logger]

func init() {
	defaultLogger.Store(NewLogger(Options{EnableTime: true}))
}

// Default returns the package's global Logger, used by the package-level
// logging functions.
func Default() *Logger {
	return defaultLogger.Load()
}

// SetDefault replaces the package's global Logger.
func SetDefault(l *Logger) {
	defaultLogger.Store(l)
}

// SetLevel changes the minimum level of the global Logger.
func SetLevel(level Level) {
	Default().SetLevel(level)
}

// With returns the global Logger with args added as attributes.
func With(args ...any) *Logger {
	return Default().With(args...)
}

// Debug logs at LevelDebug using the global Logger.
func Debug(msg string, args ...any) {
	Default().log(context.Background(), LevelDebug, msg, args)
}

// Info logs at LevelInfo using the global Logger.
func Info(msg string, args ...any) {
	Default().log(context.Background(), LevelInfo, msg, args)
}

// Warn logs at LevelWarn using the global Logger.
func Warn(msg string, args ...any) {
	Default().log(context.Background(), LevelWarn, msg, args)
}

// Error logs at LevelError using the global Logger.
func Error(msg string, args ...any) {
	Default().log(context.Background(), LevelError, msg, args)
}

// DebugContext logs at LevelDebug with ctx using the global Logger.
func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelDebug, msg, args)
}

// InfoContext logs at LevelInfo with ctx using the global Logger.
func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelInfo, msg, args)
}

// WarnContext logs at LevelWarn with ctx using the global Logger.
func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelWarn, msg, args)
}

// ErrorContext logs at LevelError with ctx using the global Logger.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().log(ctx, LevelError, msg, args)
}
