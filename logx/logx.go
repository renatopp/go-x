package logx

import (
	"context"
	"io"
	"log/slog"
	"sync/atomic"
	"time"
)

// Attr is a key-value pair attached to a log record. It retypes slog.Attr
// so Logger can build on top of log/slog without exposing it directly.
type Attr = slog.Attr

// String returns an Attr for a string value.
func String(key, value string) Attr { return slog.String(key, value) }

// Int returns an Attr for an int value.
func Int(key string, value int) Attr { return slog.Int(key, value) }

// Int64 returns an Attr for an int64 value.
func Int64(key string, value int64) Attr { return slog.Int64(key, value) }

// Uint64 returns an Attr for a uint64 value.
func Uint64(key string, value uint64) Attr { return slog.Uint64(key, value) }

// Float64 returns an Attr for a float64 value.
func Float64(key string, value float64) Attr { return slog.Float64(key, value) }

// Bool returns an Attr for a bool value.
func Bool(key string, value bool) Attr { return slog.Bool(key, value) }

// Time returns an Attr for a time.Time value.
func Time(key string, value time.Time) Attr { return slog.Time(key, value) }

// Duration returns an Attr for a time.Duration value.
func Duration(key string, value time.Duration) Attr { return slog.Duration(key, value) }

// Any returns an Attr for a value of any type.
func Any(key string, value any) Attr { return slog.Any(key, value) }

// Group returns an Attr that groups the given Attrs under key.
func Group(key string, args ...any) Attr { return slog.Group(key, args...) }

// Level is a log level. It retypes slog.Level so Logger can build on top of
// log/slog without exposing it directly.
type Level = slog.Level

// Handler is a log handler. It retypes slog.Handler so Logger can build on top
// of log/slog without exposing it directly.
type Handler = slog.Handler

type Source = slog.Source
type Record = slog.Record

// Filter is a Handler that can be chained with other Filters. It has a Next
// method to set the next Handler in the chain. When a Filter handles a log
// record, it can choose to pass it to the next Handler or not.
type Filter interface {
	Handler
	Next(Handler)
}

const LevelDebug = slog.LevelDebug
const LevelInfo = slog.LevelInfo
const LevelWarn = slog.LevelWarn
const LevelError = slog.LevelError
const LevelFatal Level = slog.LevelError * 2

// Pipeline chain multiple handlers together using Filter interface to chain
// them.
//
// If the handler is not a Filter, it will be wrapped in a Filter that just
// passes the record to the next handler.
func Pipeline(handlers ...Handler) Handler {
	if len(handlers) == 0 {
		return nil
	}

	first := AsFilter(handlers[0])
	prev := first
	for _, h := range handlers[1:] {
		cur := AsFilter(h)
		prev.Next(cur)
		prev = cur
	}
	return first
}

// Copied from slog
const badKey = "!BADKEY"

// Copied from slog
func argsToAttrSlice(args []any) []Attr {
	var (
		attr  Attr
		attrs []Attr
	)
	for len(args) > 0 {
		attr, args = argsToAttr(args)
		attrs = append(attrs, attr)
	}
	return attrs
}

// Copied from slog
func argsToAttr(args []any) (Attr, []any) {
	switch x := args[0].(type) {
	case string:
		if len(args) == 1 {
			return String(badKey, x), nil
		}
		return Any(x, args[1]), args[2:]

	case Attr:
		return x, args[1:]

	default:
		return Any(badKey, x), args[1:]
	}
}

func NewPlainTextHandler(std io.Writer) *slog.TextHandler {
	return slog.NewTextHandler(std, nil)
}

func NewJsonHandler(std io.Writer) *slog.JSONHandler {
	return slog.NewJSONHandler(std, nil)
}

var defaultLogger = atomic.Pointer[Logger]{}

func init() {
	defaultLogger.Store(NewLogger(NewColoredHandler()))
}

func DefaultLogger() *Logger {
	return defaultLogger.Load()
}

func WithTimestamp(v bool) *Logger {
	return DefaultLogger().WithTimestamp(v)
}

func WithCallerInfo(v bool) *Logger {
	return DefaultLogger().WithCallerInfo(v)
}

func WithSkipCallers(n int) *Logger {
	return DefaultLogger().WithSkipCallers(n)
}

func WithLevel(level Level) *Logger {
	return DefaultLogger().WithLevel(level)
}

func WithAttributes(attrs ...any) *Logger {
	return DefaultLogger().WithAttributes(attrs...)
}

func WithGroup(name string) *Logger {
	return DefaultLogger().WithGroup(name)
}

func Log(level Level, msg string, kvargs ...any) { DefaultLogger().Log(level, msg, kvargs...) }
func Debug(msg string, kvargs ...any)            { DefaultLogger().Debug(msg, kvargs...) }
func Info(msg string, kvargs ...any)             { DefaultLogger().Info(msg, kvargs...) }
func Warn(msg string, kvargs ...any)             { DefaultLogger().Warn(msg, kvargs...) }
func Error(msg string, kvargs ...any)            { DefaultLogger().Error(msg, kvargs...) }
func Fatal(msg string, kvargs ...any)            { DefaultLogger().Fatal(msg, kvargs...) }

func Logf(level Level, msg string, args ...any) { DefaultLogger().Logf(level, msg, args...) }
func Debugf(msg string, args ...any)            { DefaultLogger().Debugf(msg, args...) }
func Infof(msg string, args ...any)             { DefaultLogger().Infof(msg, args...) }
func Warnf(msg string, args ...any)             { DefaultLogger().Warnf(msg, args...) }
func Errorf(msg string, args ...any)            { DefaultLogger().Errorf(msg, args...) }
func Fatalf(msg string, args ...any)            { DefaultLogger().Fatalf(msg, args...) }

func Logc(ctx context.Context, level Level, msg string, kvargs ...any) {
	DefaultLogger().Logc(ctx, level, msg, kvargs...)
}
func Debugc(ctx context.Context, msg string, kvargs ...any) {
	DefaultLogger().Debugc(ctx, msg, kvargs...)
}
func Infoc(ctx context.Context, msg string, kvargs ...any) {
	DefaultLogger().Infoc(ctx, msg, kvargs...)
}
func Warnc(ctx context.Context, msg string, kvargs ...any) {
	DefaultLogger().Warnc(ctx, msg, kvargs...)
}
func Errorc(ctx context.Context, msg string, kvargs ...any) {
	DefaultLogger().Errorc(ctx, msg, kvargs...)
}
func Fatalc(ctx context.Context, msg string, kvargs ...any) {
	DefaultLogger().Fatalc(ctx, msg, kvargs...)
}
