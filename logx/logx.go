package logx

import (
	"log/slog"
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

// var defaultLogger = NewLogger(
// func (l *Logger) Log(level Level, msg string, kvargs ...any)
// func (l *Logger) Debug(msg string, kvargs ...any)
// func (l *Logger) Info(msg string, kvargs ...any)
// func (l *Logger) Warn(msg string, kvargs ...any)
// func (l *Logger) Error(msg string, kvargs ...any)
// func (l *Logger) Fatal(msg string, kvargs ...any)

// func (l *Logger) Logf(level Level, msg string, args ...any)
// func (l *Logger) Debugf(msg string, args ...any)
// func (l *Logger) Infof(msg string, args ...any)
// func (l *Logger) Warnf(msg string, args ...any)
// func (l *Logger) Errorf(msg string, args ...any)
// func (l *Logger) Fatalf(msg string, args ...any)

// func (l *Logger) Logc(ctx context.Context, level Level, msg string, kvargs ...any)
// func (l *Logger) Debugc(ctx context.Context, msg string, kvargs ...any)
// func (l *Logger) Infoc(ctx context.Context, msg string, kvargs ...any)
// func (l *Logger) Warnc(ctx context.Context, msg string, kvargs ...any)
// func (l *Logger) Errorc(ctx context.Context, msg string, kvargs ...any)
// func (l *Logger) Fatalc(ctx context.Context, msg string, kvargs ...any)
