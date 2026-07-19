// Package logx extends log/slog with a configurable Logger: levels,
// prefixes, attributes, groups, pluggable formatters, optional
// timestamp/caller/stacktrace capture, sampling, and rate limiting.
package logx

import (
	"context"
	"fmt"
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

// Debugf logs at LevelDebug using fmt.Sprintf with the global Logger.
func Debugf(format string, args ...any) {
	Default().log(context.Background(), LevelDebug, fmt.Sprintf(format, args...), nil)
}

// Infof logs at LevelInfo using fmt.Sprintf with the global Logger.
func Infof(format string, args ...any) {
	Default().log(context.Background(), LevelInfo, fmt.Sprintf(format, args...), nil)
}

// Warnf logs at LevelWarn using fmt.Sprintf with the global Logger.
func Warnf(format string, args ...any) {
	Default().log(context.Background(), LevelWarn, fmt.Sprintf(format, args...), nil)
}

// Errorf logs at LevelError using fmt.Sprintf with the global Logger.
func Errorf(format string, args ...any) {
	Default().log(context.Background(), LevelError, fmt.Sprintf(format, args...), nil)
}

// Print logs at LevelInfo using the global Logger.
func Print(msg string, args ...any) {
	Default().log(context.Background(), LevelInfo, msg, args)
}

// PrintContext logs at LevelInfo with ctx using the global Logger.
func PrintContext(ctx context.Context, msg string) {
	Default().log(ctx, LevelInfo, msg, nil)
}

// Printf logs at LevelInfo using fmt.Sprintf with the global Logger.
func Printf(format string, args ...any) {
	Default().log(context.Background(), LevelInfo, fmt.Sprintf(format, args...), nil)
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
